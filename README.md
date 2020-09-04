# my-genny-templates

## Usage

1. `go get`
    ```bash
    go get -u github.com/doilux/my-genny-templates
    ```
1. run command to create genny directory and template files
    ```bash
    $ my-genny-templates
    $ tree
    .
    ├── genny
    │   └── template
    │       ├── pointer_slice
    │       │   └── tmpl.go
    │       └── slice
    │           └── tmpl.go
    ```
1. create source code from template(this is exsample)
    ```bash
    genny -in=genny/template/slice/tmpl.go -out=gen-int-slice.go -pkg=main gen "ElementType=int64"
    ```

### go generate

```generate.go
//go:generate my-genny-templates
//go:generate genny -in=genny/template/slice/tmpl.go -out=gen-int-slice.go -pkg=main gen "ElementType=int64"
```