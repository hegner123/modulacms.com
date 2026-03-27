import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "content_section",
  label: "Content Section",
  type: "section",
  fields: [
    f.text("eyebrow", "Eyebrow"),
    f.text("heading", "Heading"),
    f.richtext("body", "Body"),
    f.media("image", "Image"),
    f.text("cta_text", "CTA Text"),
    f.url("cta_url", "CTA URL"),
  ],
};

export default schema;
