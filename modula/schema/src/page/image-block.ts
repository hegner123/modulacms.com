import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "image_block",
  label: "Image",
  type: "content",
  fields: [
    f.media("image", "Image"),
    f.text("alt_text", "Alt Text"),
    f.textarea("caption", "Caption"),
  ],
};

export default schema;
