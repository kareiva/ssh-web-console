# ssh-web-console

you can connect to your unix machine by ssh in your browser.

![Build](https://github.com/kareiva/ssh-web-console/actions/workflows/build-and-release.yml/badge.svg)
![Latest release](https://img.shields.io/github/v/release/kareiva/ssh-web-console?sort=semver)
![Go version](https://img.shields.io/github/go-mod/go-version/kareiva/ssh-web-console)

## Quick start

```bash
$ docker pull ghcr.io/kareiva/ssh-web-console:latest  # FreeBSD image: ghcr.io/kareiva/ssh-web-console:freebsd
# docker build -t ghcr.io/kareiva/ssh-web-console . # or build docker image on your own machine (from a clone with submodules)
$ docker run -v ${PWD}/conf:/home/web/conf -p 2222:2222 --rm ghcr.io/kareiva/ssh-web-console
```

Prebuilt binaries for Linux, FreeBSD (amd64/arm64), macOS and Windows (amd64) are on the [Releases](https://github.com/kareiva/ssh-web-console/releases) page.

or using Docker Compose:

```bash
$ git clone https://github.com/kareiva/ssh-web-console.git
$ cd ssh-web-console
$ docker compose up -d
```

Open your browser, visit `http://localhost:2222`. Enjoy it!

**NOTE**: To run docker container, make sure config.yaml file is in directory ${PWD}/conf

## Build & Run

Make sure your Go version is 1.27 or later, and that Node.js 22 and Yarn 1 are installed for the frontend

### Clone

```bash
git clone --recurse-submodules https://github.com/kareiva/ssh-web-console.git
cd ssh-web-console
```

### Build frontend

```bash
cd web
yarn install
yarn build
cd ../
```

### Build go

```bash
export GOTOOLCHAIN=auto # lets an older local Go download Go 1.27 automatically
go build                # embeds 'web/build' into the binary; or run 'make' to cross-compile all release binaries
```

## Run

run: `./ssh-web-console` (reads `conf/config.yaml`; use `-config <path>` to change it, `-version` to print the version), and then you can enjoy it in your browser by visiting `http://localhost:2222`.

## Screenshots

![](./Screenshots/shot2.png)
![](./Screenshots/shot3.png)
![](./Screenshots/shot4.png)

## Related Works

- [shibingli/webconsole](https://github.com/shibingli/webconsole)
