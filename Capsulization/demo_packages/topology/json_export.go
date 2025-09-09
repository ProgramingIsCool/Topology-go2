package topology

import (
	"encoding/json"
)

// MarshalToJSON converts the Topology into JSON format ([]byte).
func (t *Topology) MarshalToJSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

// UnmarshalFromJSON loads a Topology from JSON data ([]byte).
func UnmarshalFromJSON(data []byte) (*Topology, error) {
	var t Topology
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
