//go:build dev

package inertia

import (
	"fmt"
	"html/template"
	"strings"
)

func (t inertia) Vite(entrypoints ...string) template.HTML {
	builder := new(strings.Builder)
	for _, entrypoint := range append([]string{"@vite/client"}, entrypoints...) {
		fmt.Fprintf(builder, `<script type="module" src="http://localhost:5173/%s"></script>`, entrypoint)
	}
	return template.HTML(builder.String())
}
