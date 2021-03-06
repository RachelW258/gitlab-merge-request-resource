package in

import (
	"github.com/RachelW258/gitlab-merge-request-resource"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
}

type Response struct {
	Version  resource.Version  `json:"version"`
	Metadata resource.Metadata `json:"metadata"`
}
