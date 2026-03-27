import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "flyout_link",
  label: "Flyout Link",
  type: "navigation",
  fields: [
    f.text("label", "Label"),
    f.textarea("description", "Description"),
    f.url("url", "URL"),
    f.textarea("icon", "Icon"),
  ],
};

export default schema;
