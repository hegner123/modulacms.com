import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "area",
  label: "Area",
  type: "layout",
  fields: [
    f.number("column_start", "Column Start"),
    f.number("column_end", "Column End"),
    f.number("row_start", "Row Start"),
    f.number("row_end", "Row End"),
  ],
};

export default schema;
