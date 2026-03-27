import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "reference",
  label: "Reference",
  type: "_reference",
  fields: [
    f.id("target", "Target"),
  ],
};

export default schema;
