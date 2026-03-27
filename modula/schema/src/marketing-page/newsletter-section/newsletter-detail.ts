import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "newsletter_detail",
  label: "Newsletter Detail",
  type: "content",
  fields: [
    f.text("title", "Title"),
    f.textarea("description", "Description"),
    f.textarea("icon", "Icon"),
  ],
};

export default schema;
