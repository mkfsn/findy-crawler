package models

import (
	"github.com/mkfsn/findy-crawler/findy/api/v1/candidate/legacy"
)

type Job struct {
	legacy.Job `json:"_,inline"`
	IsViewed   bool `json:"is_viewed"`
}
