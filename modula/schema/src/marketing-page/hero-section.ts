import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "hero_section",
  label: "Hero Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.media("image", "Image"),
    f.media("image_dark", "Image Dark"),
    f.text("primary_cta_text", "Primary CTA Text"),
    f.url("primary_cta_url", "Primary CTA URL"),
    f.text("secondary_cta_text", "Secondary CTA Text"),
    f.url("secondary_cta_url", "Secondary CTA URL"),
    f.text("announcement_text", "Announcement Text"),
    f.url("announcement_url", "Announcement URL"),
  ],
};

export default schema;
