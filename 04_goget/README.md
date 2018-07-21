# gazelle go_repository example.
Preivously you may want to do
```bash
go get github.com/mediocregopher/radix.v2
go build ...
```
but now you can just do
```bash
bazel build -c opt :redis
```


