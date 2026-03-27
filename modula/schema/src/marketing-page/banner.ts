import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "banner",
  label: "Banner",
  type: "section",
  fields: [
    f.text("text", "Text"),
    f.text("highlight", "Highlight"),
    f.text("cta_text", "CTA Text"),
    f.url("cta_url", "CTA URL"),
    f.boolean("dismissible", "Dismissible"),
  ],
};

export default schema;
