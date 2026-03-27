import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "bento_grid_section",
  label: "Bento Grid Section",
  type: "section",
  fields: [
    f.text("eyebrow", "Eyebrow"),
    f.text("heading", "Heading"),
  ],
};

export default schema;
