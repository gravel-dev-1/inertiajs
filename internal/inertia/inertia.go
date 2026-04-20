//go:build !dev

package inertia

import (
	"encoding/json"
	"html/template"
	"path/filepath"
	"strings"
	"sync"

	"gravel/internal/vite"
)

type manifestEntry struct {
	File    string   `json:"file"`
	Imports []string `json:"imports"`
	CSS     []string `json:"css"`
}

var (
	manifestOnce sync.Once
	manifestData map[string]manifestEntry
	manifestErr  error
)

func (t inertia) Vite(entrypoints ...string) template.HTML {
	var builder strings.Builder

	manifestOnce.Do(func() {
		f, err := vite.FS.Open(filepath.Join(".vite", "manifest.json"))
		if err != nil {
			manifestErr = err
			return
		}
		defer f.Close()

		manifestData = make(map[string]manifestEntry)
		manifestErr = json.NewDecoder(f).Decode(&manifestData)
	})

	if manifestErr != nil {
		panic(manifestErr)
	}

	visited := make(map[string]struct{}, len(entrypoints))

	var walk func(string)
	walk = func(entry string) {
		if _, ok := visited[entry]; ok {
			return
		}
		visited[entry] = struct{}{}

		e, ok := manifestData[entry]
		if !ok {
			return
		}

		for _, imp := range e.Imports {
			walk(imp)
		}

		for _, css := range e.CSS {
			builder.WriteString(`<link rel="stylesheet" href="/`)
			builder.WriteString(css)
			builder.WriteString(`">`)
		}

		if e.File != "" {
			builder.WriteString(`<script type="module" src="/`)
			builder.WriteString(e.File)
			builder.WriteString(`"></script>`)
		}
	}

	for _, entry := range entrypoints {
		walk(entry)
	}

	return template.HTML(builder.String())
}
