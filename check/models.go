package check

import (
	"github.com/RachelW258/gitlab-merge-request-resource"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version,omitempty"`
}

type Response []resource.Version
