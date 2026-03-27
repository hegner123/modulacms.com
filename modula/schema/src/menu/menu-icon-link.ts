import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "menu_icon_link",
  label: "Menu Icon Link",
  type: "menu_component",
  fields: [
    f.text("label", "Label"),
    f.url("url", "URL"),
    f.select("target", "Target", ["_self","_blank"]),
    f.text("icon", "Icon"),
  ],
};

export default schema;
