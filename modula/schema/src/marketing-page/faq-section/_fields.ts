import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "faq_section",
  label: "FAQ Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("support_text", "Support Text"),
    f.url("support_url", "Support URL"),
  ],
};

export default schema;
