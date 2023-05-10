# Go Workspace with Bazel

[![Continuous Integration](https://github.com/tranngoclam/go-workspace-bazel/actions/workflows/ci.yaml/badge.svg)](https://github.com/tranngoclam/go-workspace-bazel/actions/workflows/ci.yaml)

## Prerequisites

- Go 1.20+
- Bazelisk

## Using Go Workspace

- Initialize `go.work`

```
go work init ./bar ./foo ./lib
```

- Sync modules in workspace

```
go work sync
```

## Using Bazel

- Update BUILD files

```
bazelisk run //:gazelle
```

- Sync go dependencies

```
bazelisk run //:gazellel-update-repos
```

- Run tests

```
bazelisk test //...
```

- Build binaries

```
bazelisk build //...
```

- Run binary

```
bazelisk run //foo:foo
bazelisk run //bar:bar
```