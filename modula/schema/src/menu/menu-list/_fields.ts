import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "menu_list",
  label: "Menu List",
  type: "menu_component",
  fields: [
    f.text("label", "Label"),
  ],
};

export default schema;
