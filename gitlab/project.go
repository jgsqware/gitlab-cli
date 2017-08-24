package gitlab

import (
	"encoding/json"
	"fmt"
)

type Project struct {
	ID                int
	Name              string
	NameWithNamespace string `json:"name_with_namespace"`
}

func (p Project) String() string {
	return fmt.Sprintf(p.NameWithNamespace)
}

type ProjectSearch struct {
	SearchStatus
	Projects []Project
}

func (p *ProjectSearch) Endpoint() string {
	return "/projects"
}

func (p *ProjectSearch) setResult(s SearchStatus, body []byte) {
	err := json.Unmarshal(body, &p.Projects)
	if err != nil {
		panic(err)
	}

	p.SearchStatus = s

}

func (p *ProjectSearch) Select(v interface{}) {
	i := choose(p.GetStatus(), func() {
		for i, p := range p.Projects {
			fmt.Printf("[%v] %v\n", i+1, p)
		}
	})
	*v.(*Project) = p.Projects[i]
}

func (p *ProjectSearch) GetStatus() SearchStatus {
	return p.SearchStatus
}

func GetProject(name string) Project {

	// ps := ProjectSearch{}
	// pr := ps.Perform(name).([]Project)

	// if len(pr) > 1 {
	// 	return choose(pr, func(p interface{}) string {
	// 		return p.(Project).NameWithNamespace
	// 	}).(Project)
	// }
	// return pr[0]

	return Project{}
}
