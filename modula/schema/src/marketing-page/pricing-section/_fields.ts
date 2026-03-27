import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "pricing_section",
  label: "Pricing Section",
  type: "section",
  fields: [
    f.text("eyebrow", "Eyebrow"),
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
  ],
};

export default schema;
