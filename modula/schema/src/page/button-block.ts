import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "button_block",
  label: "Button",
  type: "content",
  fields: [
    f.text("label", "Label"),
    f.url("url", "URL"),
    f.select("variant", "Variant", ["primary","secondary","outline","ghost"]),
  ],
};

export default schema;
