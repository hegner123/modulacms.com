import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "bento_cell",
  label: "Bento Cell",
  type: "content",
  fields: [
    f.text("title", "Title"),
    f.textarea("description", "Description"),
    f.media("image", "Image"),
    f.media("image_dark", "Image Dark"),
    f.select("span", "Span", ["1","2","3"]),
  ],
};

export default schema;
