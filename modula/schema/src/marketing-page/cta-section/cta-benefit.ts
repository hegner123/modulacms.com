import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "cta_benefit",
  label: "CTA Benefit",
  type: "content",
  fields: [
    f.text("text", "Text"),
  ],
};

export default schema;
