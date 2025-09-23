> [!IMPORTANT]
> **Note**: This repo has recently been donated to the OPA GitHub organization
> and we are in the process of integrating it into the OPA website. There may be
> some references to the repo's previous location while we do that.

# Rego Cheat Sheet

This project contains the code to generate the
[Rego Cheat Sheet](https://docs.styra.com/opa/rego-cheat-sheet),
in both Markdown and [PDF](https://docs.styra.com/rego-cheat-sheet.pdf)
formats.

The cheat sheet is designed to be a quick reference for those,
learning and using Rego.

Rego is the declarative language used by
[Open Policy Agent](https://www.openpolicyagent.org/), the general-purpose
policy engine.

## Development

Content is defined in YAML, Rego, and JSON in [cheats](./cheats),
and rendered to Markdown and PDF in [build](./build) using a simple
Go program.

The cheats dir content is fairly self-explanatory.
Each top level dir is a section, each sub-dir is for a single 'cheat'.

To render the outputs, after having made changes to the cheats,
run `make build`.
This requires `tectonic` to be installed (`brew install tectonic`).
