import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "menu_link",
  label: "Menu Link",
  type: "menu_component",
  fields: [
    f.text("label", "Label"),
    f.url("url", "URL"),
    f.select("target", "Target", ["_self","_blank"]),
  ],
};

export default schema;
