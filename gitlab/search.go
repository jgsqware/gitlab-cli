package gitlab

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type SearchStatus struct {
	Total, TotalPages, CurrentPage, PerPage, NextPage, PrevPage int
}

type SearchObject interface {
	Endpoint() string
	setResult(SearchStatus, []byte)
	Select(interface{})
	GetStatus() SearchStatus
}

func parseSearch(resp *http.Response) SearchStatus {
	s := SearchStatus{}
	s.NextPage, _ = strconv.Atoi(resp.Header.Get("X-Next-Page"))
	s.CurrentPage, _ = strconv.Atoi(resp.Header.Get("X-Page"))
	s.PerPage, _ = strconv.Atoi(resp.Header.Get("X-Per-Page"))
	s.PrevPage, _ = strconv.Atoi(resp.Header.Get("X-Prev-Page"))
	s.Total, _ = strconv.Atoi(resp.Header.Get("X-Total"))
	s.TotalPages, _ = strconv.Atoi(resp.Header.Get("X-Total-Pages"))
	return s
}

func (s SearchStatus) String() string {
	var tpl bytes.Buffer

	err := template.Must(template.New("search").Parse(
		`Search {{. | printf "%T" }} result:
	Total Result: {{.Total}}
	Results Per Page: {{.PerPage}}
	Current Page: {{.CurrentPage}}
	Total Pages: {{.TotalPages}}

`)).Execute(&tpl, s)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}

func choose(s SearchStatus, selection func()) int {

	sTense := "Select"
	if s.CurrentPage != s.TotalPages {
		sTense += " (c to show next page)"
	}
	var selected int

	size := s.Total
	if size > s.PerPage {
		size = s.Total - (s.CurrentPage * s.PerPage)
	}
	for {
		fmt.Println()

		selection()

		fmt.Printf("%s: ", sTense)
		fmt.Scanln(&selected)
		if selected > 0 && selected <= size {
			return selected - 1
		}
	}
}
