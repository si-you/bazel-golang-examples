load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "helloworld_proto",
    srcs = ["helloworld.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "helloworld_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    importpath = "github.com/si-you/examples/proto/helloworld",
    proto = ":helloworld_proto",
    visibility = ["//visibility:public"],
)
