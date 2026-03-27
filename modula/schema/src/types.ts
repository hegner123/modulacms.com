// Field types matching Modula's field_types table
export type FieldType =
  | "_id"
  | "_title"
  | "boolean"
  | "date"
  | "datetime"
  | "email"
  | "json"
  | "media"
  | "number"
  | "richtext"
  | "select"
  | "slug"
  | "text"
  | "textarea"
  | "url";

// Datatype categories matching Modula's type column
export type DatatypeType =
  | "_root"
  | "_global"
  | "_reference"
  | "section"
  | "layout"
  | "content"
  | "navigation"
  | "menu_component"
  | "doc_component"
  | "footer_component"
  | "settings";

export interface Field {
  name: string;
  label: string;
  type: FieldType;
  data?: Record<string, unknown>;
}

// SchemaNode is what each schema file exports.
// Parent-child relationships are defined by directory nesting.
export interface SchemaNode {
  name: string;
  label: string;
  type: DatatypeType;
  fields: Field[];
}

// Field helpers — one per field type
export const f = {
  title: (name = "title"): Field => ({ name, label: "Title", type: "_title" }),
  slug: (name = "slug"): Field => ({ name, label: "Slug", type: "slug" }),
  id: (name: string, label: string): Field => ({ name, label, type: "_id" }),
  text: (name: string, label: string): Field => ({ name, label, type: "text" }),
  textarea: (name: string, label: string): Field => ({
    name,
    label,
    type: "textarea",
  }),
  richtext: (name: string, label: string): Field => ({
    name,
    label,
    type: "richtext",
  }),
  url: (name: string, label: string): Field => ({ name, label, type: "url" }),
  media: (name: string, label: string): Field => ({
    name,
    label,
    type: "media",
  }),
  number: (name: string, label: string): Field => ({
    name,
    label,
    type: "number",
  }),
  boolean: (name: string, label: string): Field => ({
    name,
    label,
    type: "boolean",
  }),
  date: (name: string, label: string): Field => ({ name, label, type: "date" }),
  datetime: (name: string, label: string): Field => ({
    name,
    label,
    type: "datetime",
  }),
  email: (name: string, label: string): Field => ({
    name,
    label,
    type: "email",
  }),
  json: (name: string, label: string): Field => ({ name, label, type: "json" }),
  select: (name: string, label: string, options: string[]): Field => ({
    name,
    label,
    type: "select",
    data: { options },
  }),
};

