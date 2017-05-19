package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type config struct {
	GitlabURL  string
	PrivateKey string
}

var Config = config{}

type Project struct {
	Id   int
	Name string
}

func (c config) apiURL() string {
	return fmt.Sprintf("%v/api/v4", Config.GitlabURL)
}

func GetProject(name string) Project {
	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/projects", Config.apiURL()), nil)
	values := req.URL.Query()
	values.Add("search", name)
	req.URL.RawQuery = values.Encode()

	fmt.Println(req.URL.String())
	req.Header.Set("PRIVATE-TOKEN", Config.PrivateKey)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	ps := []Project{}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	json.Unmarshal(bodyBytes, &ps)
	if err != nil {
		panic(err)
	}

	fmt.Println(ps)
	return Project{}
}

func (p Project) addVariable(key, value string) {

}
