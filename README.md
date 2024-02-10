# Monorepo

[![Continuous Integration](https://github.com/tranngoclam/monorepo/actions/workflows/ci.yaml/badge.svg)](https://github.com/tranngoclam/monorepo/actions/workflows/ci.yaml)

## Prerequisites

- Go 1.22+
- Bazelisk

## Using Go Workspace

- Initialize `go.work`

```
go work init ./bar ./foo ./lib ./api
```

- Sync modules in workspace

```
go work sync
```

## Using Bazel

- Update BUILD files

```
bazel run //:gazelle
```

- Sync go dependencies

```
bazel run //:gazellel-update-repos
```

- Run tests

```
bazel test //...
```

- Build binaries

```
bazel build //...
```

- Run binary

```
bazel run //foo
bazel run //bar
```

- Run all

```
tilt up
```
