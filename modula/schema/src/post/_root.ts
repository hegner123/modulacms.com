import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "post",
  label: "Post",
  type: "_root",
  fields: [
    f.title(),
    f.slug(),
    f.textarea("excerpt", "Excerpt"),
    f.media("featured_image", "Featured Image"),
    f.select("category", "Category", ["uncategorized","news","tutorial","opinion","review","announcement"]),
    f.text("tags", "Tags"),
    f.datetime("publish_date", "Publish Date"),
    f.text("meta_title", "Meta Title"),
    f.textarea("meta_description", "Meta Description"),
    f.boolean("published", "Published"),
  ],
};

export default schema;
