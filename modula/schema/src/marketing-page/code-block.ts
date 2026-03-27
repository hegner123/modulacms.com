import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "code_block",
  label: "Code Block",
  type: "content",
  fields: [
    f.textarea("code", "Code"),
    f.select("language", "Language", ["go","javascript","typescript","html","css","bash","sql","json","yaml","lua"]),
    f.text("caption", "Caption"),
  ],
};

export default schema;
