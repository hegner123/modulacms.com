import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "newsletter_section",
  label: "Newsletter Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.text("placeholder", "Placeholder"),
    f.text("submit_text", "Submit Text"),
  ],
};

export default schema;
