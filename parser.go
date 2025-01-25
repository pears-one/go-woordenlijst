package woordenlijst

import (
	"encoding/xml"
	"fmt"
)

func parseXML(data []byte) (root, error) {
	var root root
	err := xml.Unmarshal(data, &root)
	if err != nil {
		return root, fmt.Errorf("failed to parse xml: %w", err)
	}
	return root, nil
}
