import { type SchemaNode, f } from "../../../types.js";

const schema: SchemaNode = {
  name: "pricing_tier",
  label: "Pricing Tier",
  type: "content",
  fields: [
    f.text("name", "Name"),
    f.textarea("description", "Description"),
    f.text("price_monthly", "Price Monthly"),
    f.text("price_annual", "Price Annual"),
    f.text("price_period", "Price Period"),
    f.boolean("featured", "Featured"),
    f.text("badge_text", "Badge Text"),
    f.text("cta_text", "CTA Text"),
    f.url("cta_url", "CTA URL"),
  ],
};

export default schema;
