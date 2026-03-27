import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "cta",
  label: "CTA",
  type: "content",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("subheading", "Subheading"),
    f.text("button_text", "Button Text"),
    f.url("button_url", "Button URL"),
  ],
};

export default schema;
