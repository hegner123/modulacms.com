import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "columns",
  label: "Columns",
  type: "layout",
  fields: [
    f.number("count", "Count"),
  ],
};

export default schema;
