import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "footer",
  label: "Footer",
  type: "navigation",
  fields: [
    f.title(),
    f.slug(),
    f.text("copyright", "Copyright"),
    f.text("newsletter_heading", "Newsletter Heading"),
    f.textarea("newsletter_description", "Newsletter Description"),
  ],
};

export default schema;
