import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "header",
  label: "Header",
  type: "navigation",
  fields: [
    f.media("logo", "Logo"),
    f.media("logo_dark", "Logo Dark"),
    f.text("login_text", "Login Text"),
    f.url("login_url", "Login URL"),
    f.text("cta_text", "CTA Text"),
    f.url("cta_url", "CTA URL"),
  ],
};

export default schema;
