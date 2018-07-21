load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_binary(
  name = "hello",
  srcs = ["hello.go"],
  deps = [
  ],
)

go_image(
  name = "hello_image",
  binary = ":hello",
)
