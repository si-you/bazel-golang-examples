load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
  name = "redis",
  srcs = ["redis.go"],
  deps = [
    "@com_github_mediocregopher_radix_v2//redis:go_default_library",
  ],
)
