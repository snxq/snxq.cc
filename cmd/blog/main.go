package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/snxq/snxq.xyz/pkg/blog/dao"
	"github.com/snxq/snxq.xyz/pkg/blog/server"
	"github.com/snxq/snxq.xyz/pkg/configs"
)

var (
	port = flag.String("addr", ":8080", "listen server port")
	conf = flag.String("conf", "config.yaml", "path to config.yaml")
)

func main() {
	flag.Parse()

	cfg, err := configs.Load(*conf)
	if err != nil {
		log.Fatal(err)
	}

	dao, err := dao.New(cfg.Mysql)
	if err != nil {
		log.Fatalf("init dao failed: %+v", err)
	}
	err = http.ListenAndServe(*port, server.New(dao))
	if err != nil {
		panic(err)
	}
}
