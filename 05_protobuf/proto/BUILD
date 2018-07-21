package(default_visibility=["//visibility:public"])

load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "greet_proto",
    srcs = ["greet.proto"],
)

go_proto_library(
    name = "greet_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    importpath = "github.com/si-you/examples/proto/greet",
    proto = ":greet_proto",
)
