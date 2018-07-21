load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
  name = "server",
  srcs = ["server.go"],
  deps = [
    "//proto/helloworld:helloworld_go_proto",
    "@org_golang_x_net//context:go_default_library",
    "@org_golang_google_grpc//:go_default_library",
    "@org_golang_google_grpc//reflection:go_default_library",
  ],
)

go_binary(
  name = "client",
  srcs = ["client.go"],
  deps = [
    "//proto/helloworld:helloworld_go_proto",
    "@org_golang_x_net//context:go_default_library",
    "@org_golang_google_grpc//:go_default_library",
  ],
)

