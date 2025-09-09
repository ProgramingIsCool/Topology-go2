package topology

import (
	"encoding/xml"
	"os"
)

// xmlCI is an internal helper struct for XML encoding/decoding.
type xmlCI struct {
	XMLName xml.Name `xml:"CI"`
	Name    string   `xml:"name,attr"`
	Child   []*xmlCI `xml:"CI,omitempty"`
}

// toXMLCI recursively converts CI → xmlCI.
func (ci *CI) toXMLCI() *xmlCI {
	if ci == nil {
		return nil
	}
	children := make([]*xmlCI, 0, len(ci.Child))
	for _, c := range ci.Child {
		children = append(children, c.toXMLCI())
	}
	return &xmlCI{
		Name:  ci.Name,
		Child: children,
	}
}

// MarshalToXML returns the XML representation as []byte.
func (t *Topology) MarshalToXML() ([]byte, error) {
	xmlRoots := make([]*xmlCI, 0, len(t.Roots))
	for _, r := range t.Roots {
		xmlRoots = append(xmlRoots, r.toXMLCI())
	}

	output, err := xml.MarshalIndent(struct {
		XMLName xml.Name `xml:"Topology"`
		Roots   []*xmlCI `xml:"CI"`
	}{
		Roots: xmlRoots,
	}, "", "  ")
	if err != nil {
		return nil, err
	}

	// Add XML header
	return append([]byte(xml.Header), output...), nil
}

// MarshalToXMLFile writes the XML representation directly into a file.
func (t *Topology) MarshalToXMLFile(filename string) error {
	xmlBytes, err := t.MarshalToXML()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, xmlBytes, 0644)
}
