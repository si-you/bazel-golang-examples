# Hello go image.

Run with bazel.
```bash
bazel run -c opt :hello_image
```

Or build and load into the docker

```bash
bazel build -c opt :hello_image.tar
sudo docker load -i bazel-bin/hello_image.tar
sudo docker run bazel:hello_image
```

If you are running this in MacOSX, you need to cross-compile the docker image
targeting linux & amd64 architecture.

```bash
bazel run -c opt \
  --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 \
  :hello_image
```

```bash
bazel build \
  -c opt \
  --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 \
  :hello_image.tar
docker load -i bazel-bin/hello_image.tar
docker run bazel:hello_image
```

