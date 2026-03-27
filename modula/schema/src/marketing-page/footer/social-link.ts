import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "social_link",
  label: "Social Link",
  type: "navigation",
  fields: [
    f.text("platform", "Platform"),
    f.url("url", "URL"),
    f.textarea("icon", "Icon"),
  ],
};

export default schema;
