package main

import (
	"errors"
	"flag"
)

type Config struct {
	inputFileName *string
	xmlFileName   *string
}

func (c *Config) Load() {
	c.inputFileName = flag.String("topology-file", "", "Path to the file cantaining topology data.")
	c.xmlFileName = flag.String("xml-file", "", "Path to the xml file export.")
	flag.Parse()
}

func (c Config) Validate() error {
	errNoInputFile := errors.New("please, povide the path to the topology file")
	errNoOutputFile := errors.New("please, povide the path to the xml file")
	if *c.inputFileName == "" {
		return errNoInputFile
	}
	if *c.xmlFileName == "" {
		return errNoOutputFile
	}
	return nil
}
