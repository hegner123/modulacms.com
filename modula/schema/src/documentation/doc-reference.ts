import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "doc_reference",
  label: "Reference",
  type: "doc_component",
  fields: [
    f.text("label", "Label"),
    f.url("url", "URL"),
    f.textarea("description", "Description"),
  ],
};

export default schema;
