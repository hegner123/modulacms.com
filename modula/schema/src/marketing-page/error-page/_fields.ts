import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "error_page",
  label: "Error Page",
  type: "section",
  fields: [
    f.text("error_code", "Error Code"),
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.text("home_text", "Home Text"),
    f.url("home_url", "Home URL"),
  ],
};

export default schema;
