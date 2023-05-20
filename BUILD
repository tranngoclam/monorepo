load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.work",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)

# make sure that others can import tendermint either using the
# import or go_default_library naming conventions
# gazelle:go_naming_convention import_alias

# disable proto things, so that it uses .pb.go files instead
# https://github.com/bazelbuild/rules_go/blob/master/proto/core.rst#option-2-use-pre-generated-pb-go-files
# gazelle:proto disable_global
