package main

import (
	"fmt"
	"kiliev/topology_demo_package/internal/topology"
	"kiliev/topology_demo_package/lib/utils"
	"log"
)

func main() {
	var (
		cfg            Config
		err            error
		topology_input string
	)

	/***** Load and validate CLI tool configuration *****/
	cfg.Load()
	err = cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	/***** Read the input *****/
	log.Printf("Topology data will be read from file %s\n", *cfg.inputFileName)
	topology_input, err = utils.FileContent(*cfg.inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	/***** Build Topology based on the input *****/
	tree, err := topology.BuildTopologyFromData(topology_input)
	if err != nil {
		log.Fatal(err)
	}

	/***** Generate the output *****/
	if err := tree.MarshalToXMLFile(*cfg.xmlFileName); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ XML tree was saved sucksesfuly in output.xml")
}
