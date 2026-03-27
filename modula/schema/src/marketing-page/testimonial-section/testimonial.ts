import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "testimonial",
  label: "Testimonial",
  type: "content",
  fields: [
    f.textarea("quote", "Quote"),
    f.text("author_name", "Author Name"),
    f.text("author_handle", "Author Handle"),
    f.media("author_avatar", "Author Avatar"),
    f.media("company_logo", "Company Logo"),
    f.media("company_logo_dark", "Company Logo Dark"),
    f.boolean("featured", "Featured"),
  ],
};

export default schema;
