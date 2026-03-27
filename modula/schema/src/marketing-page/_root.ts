import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "marketing_page",
  label: "Marketing Page",
  type: "_root",
  fields: [
    f.title(),
    f.textarea("description", "Description"),
    f.text("meta_title", "Meta Title"),
    f.textarea("meta_description", "Meta Description"),
    f.media("og_image", "OG Image"),
  ],
};

export default schema;
