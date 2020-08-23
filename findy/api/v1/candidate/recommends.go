package candidate

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mkfsn/findy-crawler/findy/api/v1/candidate/legacy"
)

var (
	apiClient = &http.Client{}
)

type Recommends struct {
	Jobs []legacy.Job
}

func GetRecommends(setters ...Argument) (*Recommends, error) {
	page := 1
	args := &arguments{}
	for _, setter := range setters {
		setter(args)
	}

	origin, err := getRecommends(page, args)
	if err != nil {
		return nil, err
	}

	all := &Recommends{Jobs: origin.JobList}
	for page++; page <= origin.Pagination.TotalPage; page++ {
		data, err := getRecommends(page, args)
		if err != nil {
			continue
		}
		all.Jobs = append(all.Jobs, data.JobList...)
	}
	return all, nil
}

func getRecommends(page int, args *arguments) (*legacy.Recommends, error) {
	req, err := buildRequest(page, args)
	if err != nil {
		return nil, err
	}
	res, err := apiClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	d := json.NewDecoder(res.Body)
	var data legacy.Recommends
	if err := d.Decode(&data); err != nil {
		return nil, err
	}
	if data.Error != "" {
		return nil, errors.New(data.Error)
	}
	return &data, nil
}

func buildRequest(page int, args *arguments) (*http.Request, error) {
	u, _ := url.Parse("https://findy-code.io/api/v1/candidate/recommends")
	q := u.Query()
	q.Add("page", fmt.Sprintf("%d", page))
	if args.skillId != nil {
		q.Add("skill_id", fmt.Sprintf("%d", *args.skillId))
	}
	if args.sort != nil {
		q.Add("sort", *args.sort)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "_excalibur_session", Value: args.session})
	return req, nil
}
