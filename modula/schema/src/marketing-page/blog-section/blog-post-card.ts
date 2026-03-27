import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "blog_post_card",
  label: "Blog Post Card",
  type: "content",
  fields: [
    f.text("title", "Title"),
    f.textarea("excerpt", "Excerpt"),
    f.date("date", "Date"),
    f.text("category", "Category"),
    f.url("category_url", "Category URL"),
    f.url("post_url", "Post URL"),
    f.media("featured_image", "Featured Image"),
    f.text("author_name", "Author Name"),
    f.text("author_role", "Author Role"),
    f.media("author_avatar", "Author Avatar"),
  ],
};

export default schema;
