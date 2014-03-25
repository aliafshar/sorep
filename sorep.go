// Sorep calculates the reputation for a user over a period of time.
package sorep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const URI_TEMPLATE = "https://stackoverflow.com/users/%v?tab=reputation&sort=graph"

// stripRawData strips the raw reputation data out of the profile page html.
func stripRawData(bs []byte) string {
	for _, line := range strings.Split(string(bs), "\n") {
		if strings.Contains(line, "var rawData =") {
			return line[18 : len(line)-2]
		}
	}
	return ""
}

// fetchReputationTimeline fetches a list of [y, m, d, rep_gain] for a user.
func fetchReputationTimeline(user string) ([][]int, error) {
	var v [][]int
	u := fmt.Sprintf(URI_TEMPLATE, user)
	resp, err := http.Get(u)
	if err != nil {
		return v, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return v, err
	}
	ds := stripRawData(bs)
	if ds == "" {
		return v, errors.New("No data?")
	}
	err = json.Unmarshal([]byte(ds), &v)
	if err != nil {
		return v, err
	}
	return v, nil
}

// TotalRepFor gets the total reputation gain for a user after an optional time.
func TotalRepFor(user string, after time.Time) (int, error) {
	ts, err := fetchReputationTimeline(user)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, v := range ts {
		d := time.Date(v[0], time.Month(v[1]), v[2], 0, 0, 0, 0, time.UTC)
		if d.After(after) {
			total += v[3]
		}
	}
	return total, err
}
