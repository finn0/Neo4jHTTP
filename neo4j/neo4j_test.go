package neo4j

import (
	"testing"
)

var cli = New()

func TestQueryJSON(t *testing.T) {
	cypher := `
		MATCH (n) 
		WHERE id(n) = {nodeID} 
		RETURN n LIMIT 1
	`
	params := make(map[string]interface{})
	params["nodeID"] = 0

	graph, err := cli.Query(cypher, params)
	if err != nil {
		t.Error(err)
	}

	if len(graph.Errors) > 0 {
		t.Error(graph.Errors[0])
	}

	if len(graph.Results) == 0 {
		t.Errorf("invalid graph: %v\n", graph)
	}
}
