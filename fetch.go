package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"
	"strings"

	"modulacms.com/content"

	modulacms "github.com/hegner123/modulacms/sdks/go"
)

// fetchPage calls the CMS API for the given slug in "clean" format
// and parses the response into a PageData with resolved children.
func fetchPage(ctx context.Context, client *modulacms.Client, slug string) (content.PageData, error) {
	raw, err := client.Content.GetPage(ctx, slug, "clean")
	if err != nil {
		return content.PageData{}, fmt.Errorf("fetch page %q: %w", slug, err)
	}

	var buf bytes.Buffer
	if json.Indent(&buf, raw, "", "  ") == nil {
		slog.Info("cms response\n" + buf.String())
	}

	var page content.Page
	if err := json.Unmarshal(raw, &page); err != nil {
		return content.PageData{}, fmt.Errorf("unmarshal page %q: %w", slug, err)
	}

	children, err := content.ParseChildren(page.RawChildren)
	if err != nil {
		return content.PageData{}, fmt.Errorf("parse page %q children: %w", slug, err)
	}

	buildLog := &content.BuildLog{}
	resolveMediaURLs(ctx, client, children, buildLog)

	data := content.PageData{
		Page:     page,
		Children: children,
		Log:      buildLog,
	}

	if page.Type == "documentation" {
		nav, err := fetchDocsNav(ctx, client)
		if err != nil {
			slog.Warn("failed to fetch docs nav", "error", err)
		} else {
			data.DocsNav = content.GroupDocsNav(nav)
		}
	}

	return data, nil
}

// resolveMediaURLs fetches the URL for every media reference in the tree.
func resolveMediaURLs(ctx context.Context, client *modulacms.Client, children []content.Child, log *content.BuildLog) {
	refs := content.CollectMediaRefs(children)
	for _, ref := range refs {
		media, err := client.Media.Get(ctx, modulacms.MediaID(ref.MediaID))
		if err != nil {
			msg := fmt.Sprintf("media %s: failed to resolve: %s", ref.MediaID, err)
			slog.Warn(msg)
			log.Add(msg)
			continue
		}
		if media.URL == "" {
			msg := fmt.Sprintf("media %s: empty URL", ref.MediaID)
			slog.Warn(msg)
			log.Add(msg)
			continue
		}
		*ref.URL = string(media.URL)
	}
}

// fetchDocsNav lists all CMS routes with the /docs path segment
// and returns them as navigation items sorted by slug.
func fetchDocsNav(ctx context.Context, client *modulacms.Client) ([]content.DocsNavItem, error) {
	routes, err := client.Routes.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list routes: %w", err)
	}
	var nav []content.DocsNavItem
	for _, r := range routes {
		slug := string(r.Slug)
		if !strings.HasPrefix(slug, "/docs") {
			continue
		}
		nav = append(nav, content.DocsNavItem{
			Title: r.Title,
			Slug:  slug,
		})
	}
	sort.Slice(nav, func(i, j int) bool {
		return nav[i].Slug < nav[j].Slug
	})
	return nav, nil
}
