import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "team_member",
  label: "Team Member",
  type: "content",
  fields: [
    f.text("name", "Name"),
    f.text("role", "Role"),
    f.textarea("bio", "Bio"),
    f.media("photo", "Photo"),
  ],
};

export default schema;
