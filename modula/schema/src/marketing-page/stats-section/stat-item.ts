import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "stat_item",
  label: "Stat Item",
  type: "content",
  fields: [
    f.text("value", "Value"),
    f.text("label", "Label"),
    f.textarea("description", "Description"),
  ],
};

export default schema;
