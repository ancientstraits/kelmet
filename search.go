package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/ancientstraits/kelmet/cmd"
)

var SearchCmd = &cmd.Command{
	Name:      "search",
	Usage:     "[hub | repo]",
	ShortDesc: "search packages",
	LongDesc:  "search for Helm charts",
	Run:       cmd.RunUseSubcommands,
}

const site = "https://hub.helm.sh"
const searchPath = "api/chartsvc/v1/charts/search"
const version = "Helm/3.7"

func searchToUrl(keyword string) (string, error) {
	p, err := url.Parse(site)
	if err != nil {
		return "", err
	}

	p.Path = path.Join(p.Path, searchPath)
	p.RawQuery = "q=" + url.QueryEscape(keyword)

	return p.String(), nil
}

func execRequest(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", version)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch %s: %s",
			url, res.Status)
	}

	return reqToStr(res), nil
}

func reqToStr(res *http.Response) string {
	ret := ""
	result := make([]byte, 1)
	for {
		_, err := res.Body.Read(result)
		if err == io.EOF {
			break
		}

		ret += string(result)
	}
	return ret
}

type ArtifactHubResponse struct {
	Data []struct {
		Id          string `json:"id"`
		ArtifactHub struct {
			PackageURL string `json:"packageUrl"`
		} `json:"artifactHub"`
		Attributes struct {
			Description string `json:"description"`
			Repo        struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"repo"`
		} `json:"attributes"`
		Relationships struct {
			LatestChartVersion struct {
				Data struct {
					Version    string `json:"version"`
					AppVersion string `json:"app_version"`
				} `json:"data"`
			} `json:"latestChartVersion"`
		} `json:"relationships"`
	} `json:"data"`
}

func parseResponse(result string) (*ArtifactHubResponse, error) {
	ret := ArtifactHubResponse{}
	err := json.Unmarshal([]byte(result), &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (r ArtifactHubResponse) String() string {
	ret := ""
	for _, pkg := range r.Data {
		ret += fmt.Sprintf(pkg.Id)
	}
	return ret
}

var SearchHub = &cmd.Command{
	Name:      "hub",
	Usage:     "[package_name]",
	ShortDesc: "search packages in Hub",
	LongDesc:  "search for Helm charts in Artifact Hub",
	Run: func(c *cmd.Command, args []string) error {
		url, err := searchToUrl(args[0])
		if err != nil {
			return err
		}

		result, err := execRequest(url)
		if err != nil {
			return err
		}

		data, err := parseResponse(result)
		if err != nil {
			return fmt.Errorf("failed to parse JSON: %s\n%s", err, result)
		}

		for _, pkg := range data.Data {
			fmt.Println(pkg.Id)
		}

		return nil
	},
}

var SearchRepo = &cmd.Command{
	Name:      "repo",
	Usage:     "[package_name]",
	ShortDesc: "search packages in repo",
	LongDesc:  "search for Helm charts in repositories",
	Run: func(c *cmd.Command, args []string) error {
		return nil
	},
}

func init() {
	SearchCmd.AddCommand(
		SearchHub,
		SearchRepo,
	)
}
