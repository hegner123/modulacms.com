import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "grid",
  label: "Grid",
  type: "layout",
  fields: [
    f.text("columns", "Columns"),
    f.text("rows", "Rows"),
    f.text("gap", "Gap"),
  ],
};

export default schema;
