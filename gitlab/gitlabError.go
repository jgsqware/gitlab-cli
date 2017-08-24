package gitlab

import (
	"fmt"
)

type gitlabError struct {
	Message interface{}
}

func (ge gitlabError) String() string {
	return fmt.Sprintf("message: %v", ge.Message)
}

func isKeyAlreadyExists(ge gitlabError) bool {

	return "map[key:[has already been taken]]" == fmt.Sprintf("%v", ge.Message)
}
