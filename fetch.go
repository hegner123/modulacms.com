package main

import (
	"context"
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

	var page content.Page
	if err := json.Unmarshal(raw, &page); err != nil {
		return content.PageData{}, fmt.Errorf("unmarshal home page: %w", err)
	}

	children, err := content.ParseChildren(page.RawChildren)
	if err != nil {
		return content.PageData{}, fmt.Errorf("parse home page children: %w", err)
	}

	resolveMediaURLs(ctx, client, children)

	return content.PageData{
		Page:     page,
		Children: children,
	}, nil
}

// resolveMediaURLs fetches the URL for every Image block in the tree.
func resolveMediaURLs(ctx context.Context, client *modulacms.Client, children []content.Child) {
	images := content.CollectImages(children)
	for _, img := range images {
		if img.ImageID == "" {
			continue
		}
		media, err := client.Media.Get(ctx, modulacms.MediaID(img.ImageID))
		if err != nil {
			slog.Warn("failed to resolve media URL", "mediaID", img.ImageID, "error", err)
			continue
		}
		img.MediaURL = string(media.URL)
	}
}
