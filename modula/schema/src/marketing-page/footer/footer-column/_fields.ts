import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "footer_column",
  label: "Footer Column",
  type: "navigation",
  fields: [
    f.text("heading", "Heading"),
  ],
};

export default schema;
