# Hello go image.

Run with bazel.
```bash
bazel run -c opt :hello_image
```

Or build and load into the docker

```bash
bazel build -c opt :hello.tar
sudo docker load -i bazel-bin/hello.tar
sudo docker run bazel:hello_image
```
