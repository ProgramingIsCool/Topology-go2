package topology

import (
	"fmt"
	"strings"
)

// BuildTopologyFromData accepts loaded file data (as []byte)
// and constructs a Topology tree.
// It parses each line, splits by ";", and ensures nodes are created/linked.
func BuildTopologyFromData(data string) (*Topology, error) {
	lines := strings.Split(data, "\n")
	t := &Topology{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // skip empty lines
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return nil, fmt.Errorf("bad line format: %s", line)
		}

		rootName, parentName, childName := parts[0], parts[1], parts[2]

		// Find or create root node if it doesn't exist
		rootNode := t.findCI(rootName, nil)
		if rootNode == nil {
			rootNode = t.createCI(rootName, nil)
		}

		// Check if parent is the same as root
		var parentNode *CI
		if parentName == rootName {
			parentNode = rootNode // Use root node as parent
		} else {
			// Find or create parent node under root
			parentNode = t.findCI(parentName, rootNode)
			if parentNode == nil {
				parentNode = t.createCI(parentName, rootNode)
			}
		}

		// Find or create child node if it doesn't exist
		childNode := t.findCI(childName, parentNode)
		if childNode == nil {
			_ = t.createCI(childName, parentNode)
		}
	}

	return t, nil
}
