import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "header_section",
  label: "Header Section",
  type: "section",
  fields: [
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.media("background_image", "Background Image"),
    f.media("background_image_dark", "Background Image Dark"),
  ],
};

export default schema;
