import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "team_section",
  label: "Team Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
  ],
};

export default schema;
