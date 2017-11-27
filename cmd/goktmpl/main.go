package main

import (
	"fmt"

	"github.com/aleccunningham/goktmpl/config"
	"github.com/aleccunningham/goktmpl/generate"
)

func main() {
	conf, err := config.loadConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	tmpl, err := generate.parseTemplate(conf)

	fmt.Printf("%+v\n", c)
}
