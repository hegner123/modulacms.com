import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "testimonial_section",
  label: "Testimonial Section",
  type: "section",
  fields: [
    f.text("eyebrow", "Eyebrow"),
    f.text("heading", "Heading"),
  ],
};

export default schema;
