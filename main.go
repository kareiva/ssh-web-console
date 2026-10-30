package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/genshen/ssh-web-console/src/routers"
	"github.com/genshen/ssh-web-console/src/utils"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:web/build
var webBuild embed.FS

var confFilePath string
var version bool
var buildVersion = "dev"

func init() {
	flag.StringVar(&confFilePath, "config", "conf/config.yaml", "filepath of config file.")
	flag.BoolVar(&version, "version", false, "show current version.")
}

func main() {
	flag.Parse()
	if version {
		fmt.Println(buildVersion)
		return
	}
	if err := utils.InitConfig(confFilePath); err != nil {
		log.Fatal("config error,", err)
		return
	}
	static, err := fs.Sub(webBuild, "web/build")
	if err != nil {
		log.Fatal(err)
	}
	routers.Register(static)
	log.Println("listening on port ", utils.Config.Site.ListenAddr)
	// listen http
	if err := http.ListenAndServe(utils.Config.Site.ListenAddr, nil); err != nil {
		log.Fatal(err)
		return
	}
}
