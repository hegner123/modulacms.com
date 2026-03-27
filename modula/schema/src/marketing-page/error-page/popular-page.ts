import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "popular_page",
  label: "Popular Page Link",
  type: "content",
  fields: [
    f.text("title", "Title"),
    f.textarea("description", "Description"),
    f.url("url", "URL"),
    f.textarea("icon", "Icon"),
  ],
};

export default schema;
