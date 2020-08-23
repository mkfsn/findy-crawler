package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mkfsn/findy-crawler/database"
	"github.com/mkfsn/findy-crawler/findy/api/v1/candidate"
	"github.com/mkfsn/findy-crawler/models"
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

	db, err := database.Open("jobs")
	if err != nil {
		log.Fatalf("err: %s", err)
	}

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

	// 1. Get latest job list
	recommends, err := candidate.GetRecommends(args...)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	// 1-1. Build table for lookup recommend job by job id
	idToRecommendJob := make(map[int]candidate.Job)
	for _, job := range recommends.Jobs {
		idToRecommendJob[job.Id] = job
	}

	// 2. Get all jobs from database
	var all []models.Job
	if err := db.List(&all); err != nil {
		log.Fatalf("err: %s\n", err)
	}

	// 3. Build table for lookup existing job by job id
	idToExistingJob := make(map[int]models.Job)
	for _, job := range all {
		idToExistingJob[job.Id] = job
	}

	// 3. Patch database: update current entries if exists, otherwise insert
	for _, job := range recommends.Jobs {
		existingJob, found := idToExistingJob[job.Id]
		if !found {
			b, err := json.Marshal(models.Job{Job: job, IsViewed: false})
			if err != nil {
				continue
			}
			db.Insert(job.Id, b)
		}

		// 3-1. If exists: Compare update timestamp and set is_viewed to false if updated
		if job.UpdatedAt.After(existingJob.UpdatedAt) {
			b, err := json.Marshal(models.Job{Job: job, IsViewed: false})
			if err != nil {
				continue
			}
			db.Patch(job.Id, b)
		}
	}

	// 4. Display all jobs sorted by: a) is_viewed and b) updated timestamp.
	if err := db.List(&all); err != nil {
		log.Fatalf("err: %s\n", err)
	}
	for _, job := range all {
		fmt.Printf("[%s][%s](%d) %s\n",
			isViewed(job.IsViewed),
			job.UpdatedAt.Format(time.RFC3339),
			job.Id,
			abbr(job.Title, 50),
		)
	}

	if err := db.Commit(); err != nil {
		log.Fatalf("err: %s\n", err)
	}
}

func isViewed(b bool) string {
	if b {
		return "*"
	}
	return " "
}

func abbr(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n-1]) + "…"
}
