package main

import (
	"context"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"

	"modulacms.com/content"

	modulacms "github.com/hegner123/modulacms/sdks/go"
)

// fetchHomePage calls the CMS API for the "/" slug in "clean" format
// and parses the response into a PageData with resolved children.
func fetchHomePage(ctx context.Context, client *modulacms.Client) (content.PageData, error) {
	raw, err := client.Content.GetPage(ctx, "/", "clean")
	if err != nil {
		return content.PageData{}, fmt.Errorf("fetch home page: %w", err)
	}

	var buf bytes.Buffer
	if json.Indent(&buf, raw, "", "  ") == nil {
		slog.Info("cms response\n" + buf.String())
	}

	var page content.Page
	if err := json.Unmarshal(raw, &page); err != nil {
		return content.PageData{}, fmt.Errorf("unmarshal home page: %w", err)
	}

	children, err := content.ParseChildren(page.RawChildren)
	if err != nil {
		return content.PageData{}, fmt.Errorf("parse home page children: %w", err)
	}

	buildLog := &content.BuildLog{}
	resolveMediaURLs(ctx, client, children, buildLog)

	return content.PageData{
		Page:     page,
		Children: children,
		Log:      buildLog,
	}, nil
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
