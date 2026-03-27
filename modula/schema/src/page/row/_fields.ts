import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "row",
  label: "Row",
  type: "layout",
  fields: [
    f.boolean("full_width", "Full Width"),
  ],
};

export default schema;
