// Package topology provides functionality to build a hierarchical.
// topology structure from a text file and export it to XML.
//
// # Requirements
//
// - Input file must be in the format: <bai>;<parent>;<child>
// - Each line represents one connection in the topology
//
// # Design
//
// - Topology is represented as a tree of CI nodes.
//
// - Each CI node has a Name and optional Child nodes.
//
// - The root nodes are stored in the Topology struct.

package topology

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"strings"
)

// Topology represents the root element of the topology tree.
// It contains all root CI nodes.
type Topology struct {
	XMLName xml.Name `xml:"Topology"`
	Root    []*CI    `xml:"CI"`
}

// CI represents a configuration item (node) in the topology tree.
// Each CI has a Name and optional Child nodes.
type CI struct {
	XMLName xml.Name `xml:"CI"`
	Name    string   `xml:"name,attr"`
	Child   []*CI    `xml:"CI"`
}

// LoadFromTxtFile reads a text file with ';'-separated lines
// and loads the triplets into the local Topology structure.
//
// Each line must have the format:
// <bai>;<parent>;<child>
//
// Example line:
// SUP_PR1;Cronjobs;Archive
//
// The function will log an error and exit if:
// - The file cannot be opened
// - The file cannot be read
// - Any line does not have exactly 3 parts
// - Any of the parts (bai, parent, child) is empty
func (t *Topology) LoadFromTxtFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("ERROR: Could not open file %s: %v", filename, err)
		os.Exit(-1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("ERROR: Could not read file %s: %v", filename, err)
		os.Exit(-2)
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			log.Printf("ERROR: Bad line format in file %s at line %d/%d", filename, i+1, len(lines))
			os.Exit(-3)
		}

		bai, parent, child := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2])

		if bai == "" || parent == "" || child == "" {
			log.Printf("ERROR: Empty field detected in file %s at line %d/%d", filename, i+1, len(lines))
			os.Exit(-4)
		}

		// --- root node ---
		var root *CI
		for _, r := range t.Root {
			if r.Name == bai {
				root = r
				break
			}
		}
		if root == nil {
			root = &CI{Name: bai}
			t.Root = append(t.Root, root)
		}

		// --- parent node ---
		var parentNode *CI
		if parent == bai {
			parentNode = root
		} else {
			for _, c := range root.Child {
				if c.Name == parent {
					parentNode = c
					break
				}
			}
			if parentNode == nil {
				parentNode = &CI{Name: parent}
				root.Child = append(root.Child, parentNode)
			}
		}

		// --- child node ---
		parentNode.Child = append(parentNode.Child, &CI{Name: child})
	}
}

// ExportToXmlFile serializes the Topology structure into an XML file.
//
// Example output XML looks like this:
/*
<Topology>
  <CI name="SUP_PR1">
    <CI name="Cronjobs">
      <CI name="Archive"></CI>
    </CI>
  </CI>
  <CI name="SAS_PR1">
    <CI name="Cronjobs">
      <CI name="Archive"></CI>
    </CI>
    <CI name="Process">
      <CI name="node.js"></CI>
      <CI name="oo_central"></CI>
    </CI>
  </CI>
</Topology>
*/
//
// It will log an error and exit if:
// - The XML cannot be created (marshal error)
// - The file cannot be written

func (t *Topology) ExportToXmlFile(filename string) {
	out, err := xml.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Printf("ERROR: Cannot create XML: %v", err)
		os.Exit(-5)
	}

	finalXML := xml.Header + string(out)
	err = os.WriteFile(filename, []byte(finalXML), 0644)
	if err != nil {
		log.Printf("ERROR: Cannot write XML file %s: %v", filename, err)
		os.Exit(-6)
	}
}
