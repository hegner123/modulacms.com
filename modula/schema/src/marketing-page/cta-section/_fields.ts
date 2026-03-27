import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "cta_section",
  label: "CTA Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.media("image", "Image"),
    f.text("cta_text", "CTA Text"),
    f.url("cta_url", "CTA URL"),
  ],
};

export default schema;
