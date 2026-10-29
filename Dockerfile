# build method: just run `docker build --rm -t ssh-web-console .`

# build frontend code
FROM --platform=$BUILDPLATFORM node:22-alpine AS frontend-builder

COPY web web-console/

RUN cd web-console \
    && yarn install --frozen-lockfile \
    && yarn build


FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download \
    && go install github.com/rakyll/statik@v0.1.8

COPY ./ ./
COPY --from=frontend-builder web-console/build ./web/build/

RUN statik -src=web/build \
    && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/ssh-web-console .

## copy binary
FROM alpine:3.22

ARG HOME="/home/web"

RUN adduser -D web -h ${HOME}

COPY --from=builder --chown=web /out/ssh-web-console ${HOME}/ssh-web-console

WORKDIR ${HOME}
USER web

VOLUME ["${HOME}/conf"]

CMD ["./ssh-web-console"]
