import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "logo_cloud_section",
  label: "Logo Cloud Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.text("cta_text", "CTA Text"),
    f.url("cta_url", "CTA URL"),
  ],
};

export default schema;
