package gitlab

import "encoding/json"

type GroupSearch struct {
	SearchStatus
	Groups []Group
}

type Group struct {
	ID   int
	Name string
}

func (p *GroupSearch) Endpoint() string {
	return "/groups"
}

func (g *GroupSearch) setResult(s SearchStatus, body []byte) {
	g.SearchStatus = s
	err := json.Unmarshal(body, &g.Groups)
	if err != nil {
		panic(err)
	}
}

func (g *GroupSearch) GetStatus() SearchStatus {
	return g.SearchStatus
}
