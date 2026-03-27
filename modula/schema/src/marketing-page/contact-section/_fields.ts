import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "contact_section",
  label: "Contact Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.text("submit_text", "Submit Text"),
  ],
};

export default schema;
