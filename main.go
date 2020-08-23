package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mkfsn/findy-crawler/findy/api/v1/candidate"
)

type Options struct {
	Skill   int
	Sort    string
	Session string
}

func main() {
	opts := Options{}
	flag.IntVar(&opts.Skill, "skill", -1, "Skill Id.")
	flag.StringVar(&opts.Session, "session", "", "Value of cookie `_excalibur_session`.")
	flag.StringVar(&opts.Sort, "sort", "", "Sort method, options are: newest and like.")
	flag.Parse()

	var args []candidate.Argument
	if opts.Skill != -1 {
		args = append(args, candidate.WithSkillId(opts.Skill))
	}
	if opts.Sort != "" {
		args = append(args, candidate.WithSort(opts.Sort))
	}
	if opts.Session != "" {
		args = append(args, candidate.WithSession(opts.Session))
	}

	recommends, err := candidate.GetRecommends(args...)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	for _, job := range recommends.Jobs {
		fmt.Printf("[%s] %s\n", job.UpdatedAt.Format(time.RFC3339), job.Title)
	}
}
