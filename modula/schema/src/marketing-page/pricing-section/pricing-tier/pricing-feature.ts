import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "pricing_feature",
  label: "Pricing Feature",
  type: "content",
  fields: [
    f.text("text", "Text"),
  ],
};

export default schema;
