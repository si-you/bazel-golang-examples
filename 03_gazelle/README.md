# Gazelle.
Thid dir contains a typical go proejct. You can run this:
```
export GOPATH=$PWD
go run src/hello/cmd/main.go
```

Now let's use Gazelle to turn this into a bazel project.
```
bazel run -c opt //:gazelle
```

Then `BUILD.bazel` will be automatically generated.
