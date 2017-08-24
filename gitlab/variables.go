package gitlab

import (
	"fmt"
	"net/http"
	"strings"
)

func (p Project) AddVariable(variables ...string) error {
	for _, s := range variables {
		kv := strings.Split(s, "=")
		req, err := apiRequest(http.MethodPost, strings.NewReader(fmt.Sprintf("key=%v&value=%v", kv[0], kv[1])), "/projects/%v/variables", p.ID)
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			ge := parseGitlabError(resp.Body)
			if isKeyAlreadyExists(ge) {
				ok := true
				for ok {
					var v string
					fmt.Printf("Key %v already exists, update it? ", kv[0])
					fmt.Scanln(&v)
					if v == "y" || v == "yes" {
						p.UpdateVariables(kv[0])
						break
					} else if v == "n" || v == "no" {
						break
					}
				}

			}
		}
	}
	return nil
}

func (p Project) UpdateVariables(keys ...string) error {
	for _, k := range keys {
		req, err := apiRequest(http.MethodPut, nil, "/projects/%v/variables/%v", p.ID, k)
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
		if resp.StatusCode != 200 {
			return fmt.Errorf(parseGitlabError(resp.Body).String())
		}
	}
	return nil
}

func (p Project) RemoveVariables(keys ...string) error {
	for _, k := range keys {
		req, err := apiRequest(http.MethodDelete, nil, "/projects/%v/variables/%v", p.ID, k)
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
		if resp.StatusCode != 204 {
			return fmt.Errorf(parseGitlabError(resp.Body).String())
		}
	}
	return nil
}
