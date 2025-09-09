package main

import (
	"demo_packages/topology"
	"fmt"
)

func main() {
	var t topology.Topology

	t.LoadFromTxtFile("TopologyData.txt")
	fmt.Printf("%#v\n", t)

	//	t.ExportToXmlFile("TopologyData.xml")

}
