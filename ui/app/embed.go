package app

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var embedFS embed.FS

var Assets = http.FS(MustFSSub(embedFS, "dist"))

func MustFSSub(fsys fs.FS, dir string) fs.FS {
	sfs, err := fs.Sub(embedFS, "dist")
	if err != nil {
		panic(err)
	}
	return sfs
}
