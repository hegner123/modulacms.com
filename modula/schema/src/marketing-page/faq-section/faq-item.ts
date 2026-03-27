import { type SchemaNode, f } from "../../types.js";

const schema: SchemaNode = {
  name: "faq_item",
  label: "FAQ Item",
  type: "content",
  fields: [
    f.text("question", "Question"),
    f.richtext("answer", "Answer"),
  ],
};

export default schema;
