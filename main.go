package main

import (
	"cloud-torrent/server"
	"github.com/jpillora/opts"
	"log"
)

var VERSION = "0.0.0-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "Cloud Torrent",
		Port:       3000,
		ConfigPath: "cloud-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.Parse()
	o.SetLineWidth(96)

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
