package anaconda

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func (a TwitterApi) PostStatsJobs(idaccount string, values url.Values) (job Job, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/stats/jobs/accounts/%s", AdsUrl, idaccount), values, &job, _POST, response_ch}
	return job, (<-response_ch).err
}

func (a TwitterApi) GetStatsJobs(idaccount string, values url.Values) (jobs Jobs, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{fmt.Sprintf("%s/stats/jobs/accounts/%s", AdsUrl, idaccount), values, &jobs, _GET, response_ch}
	return jobs, (<-response_ch).err
}

func (a TwitterApi) ParseStatsJobs(url string) (stats Stats, err error) {

	var (
		resp *http.Response
		archive *gzip.Reader
	)

	if resp, err = http.Get(url); err == nil {

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			err = errors.New(resp.Status)
			return
		}

		if archive, err = gzip.NewReader(resp.Body); err == nil {
			defer archive.Close()
			err = json.NewDecoder(archive).Decode(&stats)
		}

	}

	return

}