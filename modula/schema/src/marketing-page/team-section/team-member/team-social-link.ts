import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "team_social_link",
  label: "Team Social Link",
  type: "content",
  fields: [
    f.text("platform", "Platform"),
    f.url("url", "URL"),
    f.textarea("icon", "Icon"),
  ],
};

export default schema;
