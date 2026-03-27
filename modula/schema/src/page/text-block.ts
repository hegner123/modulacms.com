import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "text_block",
  label: "Text",
  type: "content",
  fields: [
    f.textarea("content", "Content"),
  ],
};

export default schema;
