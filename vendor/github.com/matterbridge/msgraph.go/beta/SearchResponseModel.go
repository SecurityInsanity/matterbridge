// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SearchResponse undocumented
type SearchResponse struct {
	// Object is the base model of SearchResponse
	Object
	// SearchTerms undocumented
	SearchTerms []string `json:"searchTerms,omitempty"`
	// HitsContainers undocumented
	HitsContainers []SearchHitsContainer `json:"hitsContainers,omitempty"`
}