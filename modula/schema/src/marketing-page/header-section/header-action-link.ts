import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "header_action_link",
  label: "Header Action Link",
  type: "content",
  fields: [
    f.text("label", "Label"),
    f.url("url", "URL"),
  ],
};

export default schema;
