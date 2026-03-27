import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "footer_text",
  label: "Footer Text",
  type: "footer_component",
  fields: [
    f.richtext("content", "Content"),
  ],
};

export default schema;
