import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "feature_item",
  label: "Feature Item",
  type: "content",
  fields: [
    f.text("title", "Title"),
    f.textarea("description", "Description"),
    f.textarea("icon", "Icon"),
    f.url("link_url", "Link URL"),
  ],
};

export default schema;
