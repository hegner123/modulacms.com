import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "inline_testimonial",
  label: "Inline Testimonial",
  type: "content",
  fields: [
    f.textarea("quote", "Quote"),
    f.text("author_name", "Author Name"),
    f.text("author_role", "Author Role"),
    f.media("company_logo", "Company Logo"),
  ],
};

export default schema;
