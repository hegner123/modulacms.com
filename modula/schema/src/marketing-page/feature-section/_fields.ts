import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "feature_section",
  label: "Feature Section",
  type: "section",
  fields: [
    f.text("eyebrow", "Eyebrow"),
    f.text("heading", "Heading"),
    f.textarea("description", "Description"),
    f.media("image", "Image"),
    f.media("image_dark", "Image Dark"),
  ],
};

export default schema;
