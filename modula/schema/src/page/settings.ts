import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "settings",
  label: "Settings",
  type: "settings",
  fields: [
    f.text("margin", "Margin"),
    f.text("padding", "Padding"),
  ],
};

export default schema;
