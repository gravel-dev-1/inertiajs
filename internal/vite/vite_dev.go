//go:build dev

package vite

import (
	"io/fs"
	"os"
	"os/exec"
)

func init() {
	cmd := exec.Command("node", "node_modules/vite/bin/vite")
	cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
	cmd.Start()
}

var FS fs.FS = os.DirFS("public")
