import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "nav_link",
  label: "Nav Link",
  type: "navigation",
  fields: [
    f.text("label", "Label"),
    f.url("url", "URL"),
  ],
};

export default schema;
