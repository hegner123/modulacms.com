import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "stats_section",
  label: "Stats Section",
  type: "section",
  fields: [
    f.text("eyebrow", "Eyebrow"),
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
  ],
};

export default schema;
