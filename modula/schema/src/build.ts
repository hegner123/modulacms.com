/**
 * Walks the src/ directory tree and compiles schema files to flat JSON.
 *
 * Convention:
 *   _root.ts   — root datatype (directory = root container)
 *   _fields.ts — intermediate datatype that has children
 *   *.ts       — leaf datatype (no children)
 *
 * Parent-child relationships are inferred from directory nesting.
 */
import { readdirSync, statSync, writeFileSync, mkdirSync } from "node:fs";
import { join, resolve, dirname } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";
import type { SchemaNode, Field } from "./types.js";

const __dirname = dirname(fileURLToPath(import.meta.url));
const outDir = resolve(__dirname, "..", "dist");

interface FlatDatatype {
  name: string;
  label: string;
  type: string;
  parent: string | null;
  fields: Field[];
}

const SKIP_FILES = new Set(["types.ts", "build.ts"]);
const SKIP_DIRS = new Set(["dist", "node_modules"]);

async function loadNode(filePath: string): Promise<SchemaNode> {
  const mod = await import(pathToFileURL(filePath).href);
  return mod.default;
}

async function walk(
  dir: string,
  parentName: string | null,
  results: FlatDatatype[],
): Promise<void> {
  const entries = readdirSync(dir).sort();

  // First pass: find this directory's own definition (_root.ts or _fields.ts)
  let dirName: string | null = null;

  for (const entry of entries) {
    if (entry !== "_root.ts" && entry !== "_fields.ts") continue;
    const node = await loadNode(join(dir, entry));
    dirName = node.name;
    results.push({
      name: node.name,
      label: node.label,
      type: node.type,
      parent: parentName,
      fields: node.fields,
    });
  }

  // Second pass: leaf files and subdirectories
  for (const entry of entries) {
    const fullPath = join(dir, entry);

    if (statSync(fullPath).isDirectory()) {
      if (SKIP_DIRS.has(entry)) continue;
      await walk(fullPath, dirName, results);
      continue;
    }

    if (!entry.endsWith(".ts")) continue;
    if (entry === "_root.ts" || entry === "_fields.ts") continue;
    if (SKIP_FILES.has(entry)) continue;

    const node = await loadNode(fullPath);
    results.push({
      name: node.name,
      label: node.label,
      type: node.type,
      parent: dirName,
      fields: node.fields,
    });
  }
}

function descendants(all: FlatDatatype[], rootIdx: number): FlatDatatype[] {
  // Track by index to handle duplicate names (e.g. code_block under
  // both marketing_page and documentation).
  const included = new Set<number>([rootIdx]);
  const rootName = all[rootIdx].name;

  let added = true;
  while (added) {
    added = false;
    for (let i = 0; i < all.length; i++) {
      if (included.has(i)) continue;
      const dt = all[i];
      if (dt.parent === null) continue;
      // Check if parent is any already-included type by matching
      // both name AND ensuring it's in the included set.
      for (const j of included) {
        if (all[j].name === dt.parent) {
          included.add(i);
          added = true;
          break;
        }
      }
    }
  }
  return Array.from(included)
    .sort((a, b) => a - b)
    .map((i) => all[i]);
}

async function main(): Promise<void> {
  const results: FlatDatatype[] = [];
  await walk(__dirname, null, results);

  mkdirSync(outDir, { recursive: true });

  // Full flat schema
  writeFileSync(
    resolve(outDir, "schema.json"),
    JSON.stringify(results, null, 2) + "\n",
  );

  // Per-root files
  const rootIndices = results
    .map((d, i) => (d.parent === null ? i : -1))
    .filter((i) => i >= 0);

  for (const idx of rootIndices) {
    const group = descendants(results, idx);
    writeFileSync(
      resolve(outDir, `${results[idx].name}.json`),
      JSON.stringify(group, null, 2) + "\n",
    );
  }

  console.log(`wrote ${results.length} datatypes to ${outDir}/`);
  for (const idx of rootIndices) {
    const count = descendants(results, idx).length;
    console.log(`  ${results[idx].name}: ${count} types`);
  }
}

main().catch(console.error);
