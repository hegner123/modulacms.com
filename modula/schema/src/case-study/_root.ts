import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "case_study",
  label: "Case Study",
  type: "_root",
  fields: [
    f.title(),
    f.slug(),
    f.text("company_name", "Company Name"),
    f.media("company_logo", "Company Logo"),
    f.url("company_url", "Company URL"),
    f.text("industry", "Industry"),
    f.textarea("description", "Description"),
    f.richtext("challenge", "Challenge"),
    f.richtext("solution", "Solution"),
    f.richtext("results", "Results"),
    f.textarea("testimonial", "Testimonial"),
    f.media("featured_image", "Featured Image"),
    f.boolean("published", "Published"),
  ],
};

export default schema;
