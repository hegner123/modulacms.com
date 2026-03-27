import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "page",
  label: "Page",
  type: "_root",
  fields: [
    f.title(),
    f.textarea("description", "Description"),
    f.media("featured_image", "Featured Image"),
    f.text("meta_title", "Meta Title"),
    f.textarea("meta_description", "Meta Description"),
    f.media("og_image", "OG Image"),
  ],
};

export default schema;
