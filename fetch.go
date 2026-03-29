package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"sync"
	"time"

	"modulacms.com/content"

	modulacms "github.com/hegner123/modulacms/sdks/go"
)

// docsNavCache caches the sidebar navigation to avoid re-fetching on every page load.
var docsNavCache struct {
	sync.RWMutex
	sections []content.DocsNavSection
	expiry   time.Time
}

const docsNavTTL = 5 * time.Minute

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
		data.DocsNav = cachedDocsNav(ctx, client)
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

// cachedDocsNav returns the docs sidebar navigation, refreshing from
// the CMS API when the cache has expired.
func cachedDocsNav(ctx context.Context, client *modulacms.Client) []content.DocsNavSection {
	docsNavCache.RLock()
	if time.Now().Before(docsNavCache.expiry) {
		sections := docsNavCache.sections
		docsNavCache.RUnlock()
		return sections
	}
	docsNavCache.RUnlock()

	items, err := fetchDocsNav(ctx, client)
	if err != nil {
		slog.Warn("failed to fetch docs nav", "error", err)
		return nil
	}
	sections := content.GroupDocsNav(items)

	docsNavCache.Lock()
	docsNavCache.sections = sections
	docsNavCache.expiry = time.Now().Add(docsNavTTL)
	docsNavCache.Unlock()
	return sections
}

// fetchDocsNav lists all CMS routes with the /docs segment, fetches
// sort_order for each page concurrently, and returns items sorted by
// sort_order.
func fetchDocsNav(ctx context.Context, client *modulacms.Client) ([]content.DocsNavItem, error) {
	routes, err := client.Routes.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list routes: %w", err)
	}

	// Collect docs routes.
	type docRoute struct {
		title string
		slug  string
	}
	var docs []docRoute
	for _, r := range routes {
		slug := string(r.Slug)
		if slug == "/docs" {
			continue // skip root docs landing page
		}
		if strings.HasPrefix(slug, "/docs/") {
			docs = append(docs, docRoute{title: r.Title, slug: slug})
		}
	}

	// Fetch sort_order concurrently for each route.
	nav := make([]content.DocsNavItem, len(docs))
	var wg sync.WaitGroup
	for i, d := range docs {
		nav[i] = content.DocsNavItem{Title: d.title, Slug: d.slug}
		wg.Add(1)
		go func(idx int, slug string) {
			defer wg.Done()
			raw, err := client.Content.GetPage(ctx, slug, "clean")
			if err != nil {
				return
			}
			var page struct {
				SortOrder *int `json:"sort_order"`
			}
			if json.Unmarshal(raw, &page) == nil && page.SortOrder != nil {
				nav[idx].SortOrder = *page.SortOrder
			}
		}(i, d.slug)
	}
	wg.Wait()

	sort.Slice(nav, func(i, j int) bool {
		if nav[i].SortOrder != nav[j].SortOrder {
			return nav[i].SortOrder < nav[j].SortOrder
		}
		return nav[i].Slug < nav[j].Slug
	})
	return nav, nil
}
