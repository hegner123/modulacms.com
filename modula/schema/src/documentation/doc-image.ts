import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "doc_image",
  label: "Image",
  type: "doc_component",
  fields: [
    f.media("image", "Image"),
    f.text("alt_text", "Alt Text"),
    f.text("caption", "Caption"),
  ],
};

export default schema;
