import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "rich_text_block",
  label: "Rich Text",
  type: "content",
  fields: [
    f.richtext("content", "Content"),
  ],
};

export default schema;
