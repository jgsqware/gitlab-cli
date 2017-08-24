package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

var client = http.DefaultClient

type Gitlab struct{}

var DefaultClient = &Gitlab{}

func (g Gitlab) Search(s SearchObject, searchString string) SearchObject {
	req, err := apiRequest(http.MethodGet, nil, s.Endpoint())

	if err != nil {
		panic(err)
	}

	values := req.URL.Query()
	values.Add("search", searchString)
	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	parseSearch(resp)

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	s.setResult(parseSearch(resp), bodyBytes)

	return s
}

func apiRequest(method string, body io.Reader, format string, a ...interface{}) (*http.Request, error) {
	apiURL := fmt.Sprintf("%v/api/v4", viper.GetString("url"))
	req, err := http.NewRequest(method, apiURL+fmt.Sprintf(format, a...), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", viper.GetString("pkey"))

	return req, nil
}

func parseGitlabError(body io.ReadCloser) gitlabError {
	var ge gitlabError
	if err := json.NewDecoder(body).Decode(&ge); err != nil {
		panic(err)
	}
	return ge
}
