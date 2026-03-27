import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "blog_section",
  label: "Blog Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
  ],
};

export default schema;
