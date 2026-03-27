import { type SchemaNode, f } from "../types.js";

const schema: SchemaNode = {
  name: "animation",
  label: "Animation",
  type: "content",
  fields: [
    f.select("type", "Type", ["fade","slide","scale","rotate"]),
    f.text("duration", "Duration"),
    f.text("delay", "Delay"),
    f.select("easing", "Easing", ["ease","ease-in","ease-out","ease-in-out","linear"]),
    f.select("direction", "Direction", ["normal","reverse","alternate"]),
    f.text("iterations", "Iterations"),
  ],
};

export default schema;
