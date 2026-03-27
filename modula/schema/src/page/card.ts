import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "card",
  label: "Card",
  type: "content",
  fields: [
    f.text("title", "Title"),
    f.textarea("description", "Description"),
    f.media("image", "Image"),
    f.url("link_url", "Link URL"),
  ],
};

export default schema;
