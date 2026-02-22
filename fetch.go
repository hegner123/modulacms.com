package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"

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

	buildLog := &content.BuildLog{}
	resolveMediaURLs(ctx, client, children, buildLog)

	return content.PageData{
		Page:     page,
		Children: children,
		Log:      buildLog,
	}, nil
}

// mediaURLResponse is a minimal struct to extract just the URL from the
// media API response, bypassing SDK deserialization issues.
type mediaURLResponse struct {
	URL string `json:"url"`
}

// resolveMediaURLs fetches the URL for every Image block in the tree.
func resolveMediaURLs(ctx context.Context, client *modulacms.Client, children []content.Child, log *content.BuildLog) {
	images := content.CollectImages(children)
	for _, img := range images {
		if img.ImageID == "" {
			log.Add(fmt.Sprintf("Image %s: empty media ID", img.ID))
			continue
		}
		params := url.Values{}
		params.Set("q", img.ImageID)
		raw, err := client.Media.RawList(ctx, params)
		if err != nil {
			msg := fmt.Sprintf("Image %s: failed to fetch media %s: %s", img.ID, img.ImageID, err)
			slog.Warn(msg)
			log.Add(msg)
			continue
		}
		var media mediaURLResponse
		if err := json.Unmarshal(raw, &media); err != nil {
			msg := fmt.Sprintf("Image %s: failed to decode media %s: %s", img.ID, img.ImageID, err)
			slog.Warn(msg)
			log.Add(msg)
			continue
		}
		if media.URL == "" {
			msg := fmt.Sprintf("Image %s: media %s has empty URL", img.ID, img.ImageID)
			slog.Warn(msg)
			log.Add(msg)
			continue
		}
		img.MediaURL = media.URL
	}
}
