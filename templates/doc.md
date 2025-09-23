# Rego Cheat Sheet

<!-- The source of truth for this file's contents is https://github.com/StyraInc/rego-cheat-sheet -->

All code examples on this page share this preamble:

```rego
package cheat
import rego.v1
```
<RunSnippet id="preamble.rego"/>

{{ range .Sections }}
## {{ .Title }} - <sub><sup>{{ .Subtitle }}</sup></sub>

{{ range .Cheats }}
### {{ .Title }} 

{{ if ne .Text "" }}
{{ .Text }} ([Try It]({{ .PlaygroundLink }}))
{{ else }}
[Try It]({{ .PlaygroundLink }})
{{ end }}

{{ if ne .Input "" }}
```json title="input.json"
{{ .Input | trim }}
```
<RunSnippet id="input.{{.Title | urlquery}}.json"/>
{{ end }}

```rego
{{ .CodeDisplay }}
```

{{ if ne .Input "" }}
<RunSnippet command="data.cheat" files="#input.{{.Title | urlquery}}.json" depends="preamble.rego"/>
{{ else }}
<RunSnippet command="data.cheat" depends="preamble.rego"/>
{{ end }}
{{ end }}
{{ end }}

{{ .Footer }}
