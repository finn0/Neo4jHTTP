package neo4j

// GraphRoot is struct of a neo4j HTTP API query response body
type GraphRoot struct {
	Results []Result `json:"results"`
	Errors  []Error  `json:"errors"`
}

// Result -
type Result struct {
	Columns []string `json:"columns"`
	Data    []Data   `json:"data"`
}

// Data -
type Data struct {
	Graph Graph       `json:"graph"`
	Row   interface{} `json:"row,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
}

// Graph -
type Graph struct {
	Nodes         []Node         `json:"nodes"`
	Relationships []Relationship `json:"relationships"`
}

// Relationship -
type Relationship struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	StartNode  string                 `json:"startNode"`
	EndNode    string                 `json:"endNode"`
	Properties map[string]interface{} `json:"properties"`
}

// Node -
type Node struct {
	ID         string                 `json:"id"`
	Labels     []string               `json:"labels"`
	Properties map[string]interface{} `json:"properties"`
}

// Error -
type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
