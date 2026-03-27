import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "post_content",
  label: "Content",
  type: "content",
  fields: [
    f.richtext("content", "Content"),
  ],
};

export default schema;
