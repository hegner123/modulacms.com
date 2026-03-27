import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "step_header",
  label: "Step Header",
  type: "doc_component",
  fields: [
    f.number("step_number", "Step Number"),
    f.text("title", "Title"),
    f.textarea("description", "Description"),
  ],
};

export default schema;
