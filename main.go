package main

import (
	_ "github.com/doilux/my-genny-templates/statik"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"os"
)

const rootDir = "genny/template"

var types = []string{
	"pointer_slice",
	"slice",
}


func main() {
	fileSystem, err := fs.New()
	if err != nil {
		panic(err)
	}

	for _, v := range types {
		f, err := fileSystem.Open("/" + v + "/tmpl.go")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		contents, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(rootDir + "/" + v , 0777);
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(rootDir + "/" + v + "/tmpl.go", []byte(contents), 0664)
		if err != nil {
			panic(err)
		}
	}
}

