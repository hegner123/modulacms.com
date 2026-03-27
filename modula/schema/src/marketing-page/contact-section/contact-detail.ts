import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "contact_detail",
  label: "Contact Detail",
  type: "content",
  fields: [
    f.select("type", "Type", ["address","phone","email"]),
    f.text("value", "Value"),
    f.url("url", "URL"),
    f.textarea("icon", "Icon"),
  ],
};

export default schema;
