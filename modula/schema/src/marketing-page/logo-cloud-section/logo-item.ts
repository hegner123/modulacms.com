import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "logo_item",
  label: "Logo Item",
  type: "content",
  fields: [
    f.text("company_name", "Company Name"),
    f.media("logo", "Logo"),
    f.media("logo_dark", "Logo Dark"),
  ],
};

export default schema;
