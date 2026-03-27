import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "footer_social",
  label: "Footer Social",
  type: "footer_component",
  fields: [
    f.select("platform", "Platform", ["github","twitter","linkedin","youtube","discord","instagram","facebook","mastodon"]),
    f.url("url", "URL"),
  ],
};

export default schema;
