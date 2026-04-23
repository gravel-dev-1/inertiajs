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

var FS fs.FS = &filesystems{[]fs.FS{os.DirFS("internal/vite/dev"), os.DirFS("public")}}

type filesystems struct{ filesystems []fs.FS }

func (f filesystems) Open(name string) (file fs.File, err error) {
	for _, filesystem := range f.filesystems {
		if file, err = filesystem.Open(name); file != nil {
			return file, err
		}
	}
	return nil, fs.ErrNotExist
}
