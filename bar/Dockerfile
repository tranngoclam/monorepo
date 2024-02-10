FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS build

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=./bar/go.mod,target=./bar/go.mod \
    --mount=type=bind,source=./bar/go.sum,target=./bar/go.sum \
    --mount=type=bind,source=./lib/go.mod,target=./lib/go.mod \
    --mount=type=bind,source=./lib/go.sum,target=./lib/go.sum \
    --mount=type=bind,source=./api/go.mod,target=./api/go.mod \
    --mount=type=bind,source=./api/go.sum,target=./api/go.sum \
    go work init ./bar ./lib ./api && go mod download

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=./bar,target=./bar \
    --mount=type=bind,source=./lib,target=./lib \
    --mount=type=bind,source=./api,target=./api \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/bar bar

FROM gcr.io/distroless/static-debian11

COPY --from=build /bin/bar /bin/bar

CMD [ "/bin/bar" ]
