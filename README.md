# my-genny-templates

## How to Develop

1. install statik
    ```bash
    go get -u github.com/rakyll/statik
    ```
1. create directory and tmpl.go at tmpl directory
1. add directory name to main.go
    ```go
    var types = []string{
        "pointer_slice",
        "slice",
        /* add new template direcotry name here */
    }
    ```

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