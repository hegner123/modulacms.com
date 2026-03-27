import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "doc_section",
  label: "Section",
  type: "doc_component",
  fields: [
    f.text("heading", "Heading"),
    f.richtext("content", "Content"),
  ],
};

export default schema;
