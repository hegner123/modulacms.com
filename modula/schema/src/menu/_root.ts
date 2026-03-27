import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "menu",
  label: "Menu",
  type: "_global",
  fields: [
    f.title(),
    f.slug(),
    f.select("position", "Position", ["header","sidebar"]),
  ],
};

export default schema;
