package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"testing"
)

// const site = "https://hub.helm.sh"
// const searchPath = "api/chartsvc/v1/charts/search"
// const version = "Helm/3.7"

const success = 200
const resultSize = 256

const searchTerm = "hello"

func TestSearch(t *testing.T) {
	t.Log("hello world")
	p, err := url.Parse(site)
	if err != nil {
		t.Error(err)
	}

	p.Path = path.Join(p.Path, searchPath)
	p.RawQuery = "q=" + url.QueryEscape(searchTerm)

	req, err := http.NewRequest("GET", p.String(), nil)
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("User-Agent", version)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if res.StatusCode != success {
		t.Errorf("failed to fetch %s: %s", p.String(), res.Status)
	}

	result := make([]byte, resultSize)
	for {
		n, err := res.Body.Read(result)
		t.Log(string(result[:n]))
		if err == io.EOF {
			break
		}
	}
}


func TestJSON(t *testing.T) {
	data := `
		{"data":
			[
				{
					"id":"hello/hello",
					"artifactHub": {
						"packageUrl":"https://artifacthub.io/packages/helm/hello/hello"
					},
					"attributes": {
						"description":"A Helm chart for my simple hello application",
						"repo":{"name":"hello","url":"https://www.kleinloog.ch/hello-helm/"}
					},
					"relationships": {
						"latestChartVersion":
							{"data":{"version":"0.4.0-rc2","app_version":"v4"}
						}
					}
				}
			]
		}`
	into := ArtifactHubResponse{}
	err := json.Unmarshal([]byte(data), &into)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", into)
}
