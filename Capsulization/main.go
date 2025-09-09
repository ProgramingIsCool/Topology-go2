package main

import (
	"fmt"
	"kiliev/topology_demo_package/demo_packages/topology"
	"log"
)

func main() {
	// Read file from (input.txt)
	data, err := topology.LoadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create tree from data
	tree, err := topology.BuildTopologyFromData(data)
	if err != nil {
		log.Fatal(err)
	}

	// Write in XML file
	if err := tree.MarshalToXMLFile("output.xml"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ XML tree was saved sucksesfuly in output.xml")
}
