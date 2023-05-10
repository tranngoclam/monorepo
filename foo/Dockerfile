FROM golang:1.20 AS build

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=./foo/go.mod,target=./foo/go.mod \
    --mount=type=bind,source=./foo/go.sum,target=./foo/go.sum \
    --mount=type=bind,source=./lib/go.mod,target=./lib/go.mod \
    --mount=type=bind,source=./lib/go.sum,target=./lib/go.sum \
    go work init ./foo ./lib && go work sync

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=./foo,target=./foo \
    --mount=type=bind,source=./lib,target=./lib \
    CGO_ENABLED=0 go build -x -o /bin/foo foo

FROM gcr.io/distroless/static-debian11

COPY --from=build /bin/foo /bin/foo

CMD [ "/bin/foo" ]
