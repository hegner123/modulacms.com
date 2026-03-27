import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "footer_link",
  label: "Footer Link",
  type: "navigation",
  fields: [
    f.text("label", "Label"),
    f.url("url", "URL"),
    f.select("target", "Target", ["_self","_blank"]),
  ],
};

export default schema;
