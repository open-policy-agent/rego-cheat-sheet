package main

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/go-yaml/yaml"
)

type Section struct {
	Title     string `yaml:"title"`
	Subtitle  string `yaml:"subtitle"`
	Rank      int    `yaml:"rank"`
	PageBreak bool   `yaml:"page_break"`
	Cheats    []Cheat
}

type Cheat struct {
	Title         string `yaml:"title"`
	Link          string `yaml:"link"`
	Rank          int    `yaml:"rank"`
	Text          string
	Output        string
	TexHideOutput bool `yaml:"tex_hide_output"`
	Code          string
	Input         string
}

func (c *Cheat) CodeDisplay() string {
	code := c.Code
	code = strings.Replace(code, "package cheat", "", 1)
	code = strings.Replace(code, "import rego.v1", "", 1)
	code = strings.TrimSpace(code)

	return code
}

func (c *Cheat) PlaygroundLink() string {
	var u url.URL
	u.Scheme = "https"
	u.Host = "play.openpolicyagent.org"
	u.Path = "/"

	state := map[string]interface{}{
		"p": c.Code,
	}

	if c.Input != "" {
		state["i"] = c.Input
	} else {
		state["i"] = "{}"
	}

	jsonConfig, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}

	base64Config := base64.StdEncoding.EncodeToString(jsonConfig)

	params := make(url.Values)
	params.Add("state", base64Config)

	u.RawQuery = params.Encode()

	return u.String()
}

func main() {
	var err error

	sectionDirs, err := os.ReadDir("cheats")
	if err != nil {
		panic(err)
	}

	var sections []Section
	for _, sectionDir := range sectionDirs {
		if !sectionDir.IsDir() {
			continue
		}

		var cheatDirs []os.DirEntry
		cheatDirs, err = os.ReadDir(filepath.Join("cheats", sectionDir.Name()))
		if err != nil {
			panic(err)
		}

		var cheats []Cheat
		for _, cheatDir := range cheatDirs {
			if !cheatDir.IsDir() {
				continue
			}

			var cheatYAMLBs []byte
			cheatYAMLBs, err = os.ReadFile(filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "cheat.yaml"))
			if err != nil {
				panic(err)
			}

			var cheat Cheat
			err = yaml.Unmarshal(cheatYAMLBs, &cheat)
			if err != nil {
				panic(err)
			}

			var cheatCodeBs []byte
			cheatCodeBs, err = os.ReadFile(filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "cheat.rego"))
			if err != nil {
				panic(err)
			}
			cheat.Code = string(cheatCodeBs)

			inputPath := filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "input.json")
			if _, err = os.Stat(inputPath); !os.IsNotExist(err) {
				var cheatInputBs []byte
				cheatInputBs, err = os.ReadFile(inputPath)
				if err != nil {
					panic(err)
				}
				cheat.Input = string(cheatInputBs)
			}

			cheatOutputPath := filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "output.json")
			if _, err = os.Stat(cheatOutputPath); !os.IsNotExist(err) {
				cheatOutputBs, err := os.ReadFile(cheatOutputPath)
				if err != nil {
					panic(err)
				}

				cheat.Output = strings.TrimSpace(string(cheatOutputBs))
			}

			cheatTextPath := filepath.Join("cheats", sectionDir.Name(), cheatDir.Name(), "cheat.txt")
			if _, err = os.Stat(cheatTextPath); !os.IsNotExist(err) {
				cheatTextBs, err := os.ReadFile(cheatTextPath)
				if err != nil {
					panic(err)
				}

				cheat.Text = string(cheatTextBs)
			}

			cheats = append(cheats, cheat)
		}

		sort.Slice(cheats, func(i, j int) bool {
			return cheats[i].Rank < cheats[j].Rank
		})

		var section Section

		var sectionYAMLBs []byte
		sectionYAMLBs, err = os.ReadFile(filepath.Join("cheats", sectionDir.Name(), "section.yaml"))
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(sectionYAMLBs, &section)
		if err != nil {
			panic(err)
		}
		section.Cheats = cheats

		sections = append(sections, section)
	}

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].Rank < sections[j].Rank
	})

	// generate md
	mdTemplateBs, err := os.ReadFile("templates/doc.md")
	if err != nil {
		panic(err)
	}

	mdTemplate, err := template.New("md").
		Funcs(template.FuncMap{
			"trim": strings.TrimSpace,
		}).
		Parse(string(mdTemplateBs))
	if err != nil {
		panic(err)
	}

	mdOutputPath := "build/cheatsheet.md"
	if _, err = os.Stat(mdOutputPath); !os.IsNotExist(err) {
		err = os.Remove(mdOutputPath)
		if err != nil {
			panic(err)
		}
	}
	_, err = os.Create(mdOutputPath)
	if err != nil {
		panic(err)
	}

	mdOutputFile, err := os.OpenFile(mdOutputPath, os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}

	var footerBs []byte
	footerBs, err = os.ReadFile("assets/footer.md")
	if err != nil {
		panic(err)
	}

	err = mdTemplate.Execute(mdOutputFile, struct {
		Sections []Section
		Footer   string
	}{Sections: sections, Footer: string(footerBs)})
	if err != nil {
		panic(err)
	}

	// generate tex
	texTemplateBs, err := os.ReadFile("templates/doc.tex")
	if err != nil {
		panic(err)
	}

	texTemplate, err := template.New("tex").Parse(string(texTemplateBs))
	if err != nil {
		panic(err)
	}

	texOutputPath := "build/cheatsheet.tex"
	if _, err = os.Stat(texOutputPath); err == nil {
		err = os.Remove(texOutputPath)
		if err != nil {
			panic(err)
		}
	}
	_, err = os.Create(texOutputPath)
	if err != nil {
		panic(err)
	}

	texOutputFile, err := os.OpenFile(texOutputPath, os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}

	footerBs, err = os.ReadFile("assets/footer.tex")
	if err != nil {
		panic(err)
	}

	err = texTemplate.Execute(texOutputFile, struct {
		Sections []Section
		Footer   string
	}{Sections: sections, Footer: string(footerBs)})
	if err != nil {
		panic(err)
	}

	assets, err := os.ReadDir("assets")
	if err != nil {
		panic(err)
	}

	for _, asset := range assets {
		if asset.IsDir() {
			panic("dir in assets, can't process")
		}

		targetPath := filepath.Join("build", asset.Name())
		if _, err = os.Stat(targetPath); !os.IsNotExist(err) {
			err = os.Remove(targetPath)
			if err != nil {
				panic(err)
			}
		}

		err = os.Link(filepath.Join("assets", asset.Name()), targetPath)
		if err != nil {
			panic(err)
		}
	}
}
