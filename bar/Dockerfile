FROM golang:1.20 AS build

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=./bar/go.mod,target=./bar/go.mod \
    --mount=type=bind,source=./bar/go.sum,target=./bar/go.sum \
    --mount=type=bind,source=./lib/go.mod,target=./lib/go.mod \
    --mount=type=bind,source=./lib/go.sum,target=./lib/go.sum \
    go work init ./bar ./lib && go work sync

COPY bar bar

COPY lib lib

WORKDIR /src/bar

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -x -o /bin/bar .

FROM gcr.io/distroless/static-debian11

COPY --from=build /bin/bar /bin/bar

CMD [ "/bin/bar" ]
