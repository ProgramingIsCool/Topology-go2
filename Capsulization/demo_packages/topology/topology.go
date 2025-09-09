package topology

// Topology represents the whole tree structure.
// It contains one or more root nodes.
type Topology struct {
	Roots []*CI
}

// CI (Configuration Item) represents one node in the tree.
// Each node has a Name and a list of children (sub-nodes).
type CI struct {
	Name  string
	Child []*CI
}

// findCI searches for a CI node by name under a given parent.
// If parent is nil, search is done among root nodes.
func (t *Topology) findCI(name string, parent *CI) *CI {
	if parent == nil {
		for _, r := range t.Roots {
			if r.Name == name {
				return r
			}
		}
		return nil
	}

	for _, c := range parent.Child {
		if c.Name == name {
			return c
		}
	}
	return nil
}

// createCI creates a new CI node under the given parent (or as a root if parent is nil).
func (t *Topology) createCI(name string, parent *CI) *CI {
	newCI := &CI{Name: name}
	if parent == nil {
		t.Roots = append(t.Roots, newCI)
	} else {
		parent.Child = append(parent.Child, newCI)
	}
	return newCI
}
