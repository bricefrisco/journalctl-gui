package journal

import (
	"sort"

	"github.com/coreos/go-systemd/v22/sdjournal"
)

type Entry struct {
	Message   string `json:"message"`
	Priority  string `json:"priority"`
	Unit      string `json:"unit"`
	Timestamp uint64 `json:"timestamp"`
}

func ListServices() ([]string, error) {
	j, err := sdjournal.NewJournal()
	if err != nil {
		return nil, err
	}
	defer j.Close()

	values, err := j.GetUniqueValues("_SYSTEMD_UNIT")
	if err != nil {
		return nil, err
	}

	sort.Strings(values)
	return values, nil
}
