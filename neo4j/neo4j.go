package neo4j

import (
	"Neo4jHTTP/conf"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Neo4j is http client for communicating with Graph Neo4j
type Neo4j struct {
	cfg *conf.Config
}

// New a Neo4j instance
func New() *Neo4j {
	return &Neo4j{
		cfg: conf.GetConfig(),
	}
}

// QueryJSON returns the json data responsed
func (cli *Neo4j) QueryJSON(cypher string, params map[string]interface{}, contents ...string) ([]byte, error) {
	bs, err := cli.post(cypher, params, contents...)
	if err != nil {
		return nil, err
	}
	return bs, err
}

// Query returns a graph from unmarshal json
func (cli *Neo4j) Query(cypher string, params map[string]interface{}, contents ...string) (*GraphRoot, error) {
	graph := &GraphRoot{}
	bs, err := cli.post(cypher, params, contents...)
	if err != nil {
		return graph, err
	}
	err = json.Unmarshal(bs, graph)
	if err != nil {
		return graph, err

	}
	return graph, nil
}

func (cli *Neo4j) post(cypher string, params map[string]interface{}, contents ...string) ([]byte, error) {
	body, err := cli.singleStatements(cypher, params, contents...)
	if err != nil {
		return nil, err
	}

	httpcli := new(http.Client)
	req, err := http.NewRequest("POST", cli.cfg.URL(), body)
	req.Header = cli.header()
	req.SetBasicAuth(cli.cfg.Username, cli.cfg.Password)

	res, err := httpcli.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return result, err
}

func (cli *Neo4j) header() http.Header {
	header := &http.Header{}
	header.Set("Accept", "application/json; charset=UTF-8")
	header.Set("X-Stream", "true")
	return *header
}

func (cli *Neo4j) singleStatements(cypher string, params map[string]interface{}, contents ...string) (io.Reader, error) {
	sts := getSingleStatements(cypher, params, contents...)
	bs, err := json.Marshal(sts)
	return bytes.NewBuffer(bs), err
}
