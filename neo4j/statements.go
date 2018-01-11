package neo4j

// Result data content type
const (
	ContentGraph = "graph"
	ContentRow   = "row"
)

type statements struct {
	Statements []statement `json:"statements"`
}

type statement struct {
	Statement          string                 `json:"statement"`
	Parameters         map[string]interface{} `json:"parameters"`
	ResultDataContents []string               `json:"resultDataContents"`
}

func getSingleStatements(cypher string, params map[string]interface{}, contents ...string) statements {
	cts := contents[:]
	if len(contents) == 0 {
		// query with result data type graph as default
		cts = append(cts, ContentGraph)
	}

	st := &statement{
		Statement:          cypher,
		Parameters:         params,
		ResultDataContents: cts,
	}
	sts := &statements{}
	sts.Statements = []statement{*st}

	return *sts
}
