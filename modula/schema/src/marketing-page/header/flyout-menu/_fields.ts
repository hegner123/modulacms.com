import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "flyout_menu",
  label: "Flyout Menu",
  type: "navigation",
  fields: [
    f.text("label", "Label"),
  ],
};

export default schema;
