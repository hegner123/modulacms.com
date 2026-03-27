import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "hero",
  label: "Hero",
  type: "content",
  fields: [
    f.text("header", "Header"),
    f.textarea("subtitle", "Subtitle"),
    f.media("image", "Image"),
    f.url("link_url", "Link URL"),
  ],
};

export default schema;
