import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "documentation",
  label: "Documentation",
  type: "_root",
  fields: [
    f.title(),
    f.slug(),
    f.text("chapter", "Chapter"),
    f.number("sort_order", "Sort Order"),
    f.textarea("description", "Description"),
    f.boolean("published", "Published"),
  ],
};

export default schema;
