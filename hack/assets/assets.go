//go:generate go run ./assets.go

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shurcooL/vfsgen"
)

type fileSystem struct {
	http.FileSystem
}

func (fs fileSystem) Open(name string) (http.File, error) {
	f, err := fs.FileSystem.Open(name)
	return file{f}, err
}

type file struct {
	http.File
}

func (f file) Stat() (os.FileInfo, error) {
	fi, err := f.File.Stat()
	return fileInfo{fi}, err
}

type fileInfo struct {
	os.FileInfo
}

func (fi fileInfo) ModTime() time.Time { return time.Time{} }

func main() {
	err := vfsgen.Generate(fileSystem{http.Dir("../../data/data")}, vfsgen.Options{
		Filename:     "../../data/assets_vfsdata.go",
		PackageName:  "data",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
