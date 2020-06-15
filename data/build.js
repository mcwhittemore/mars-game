const { readFileSync, writeFileSync } = require('fs');
const { join } = require('path');

const root = join(__dirname, '..');

const files = process.argv.slice(2).map(f => join(process.cwd(), f))

const defs = []
for(let i=0; i<files.length; i++) {
  const fn = files[i];
  const name = fn.replace(root, '')
  const file = readFileSync(fn);
  const bytes = [];
  for(let j=0; j < file.length; j++) {
    bytes.push(file[j]);
  }
  defs.push(`files["${name}"] = []byte{${bytes.join(", ")}}`);
}

const out = `package data

func init() {
\tfiles = make(map[string][]byte, ${files.length})

\t${defs.join("\n\t")}

}
`;

writeFileSync(join(root, 'data', 'init.go'), out);
