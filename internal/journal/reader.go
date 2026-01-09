package journal

import (
	"sort"
	"strings"
	"time"

	"github.com/bricefrisco/journalctl-gui/internal/util"
	"github.com/coreos/go-systemd/v22/sdjournal"
)

type LogsPage struct {
	Items      []LogEntry `json:"items"`
	NextCursor string     `json:"nextCursor,omitempty"`
	HasMore    bool       `json:"hasMore"`
}

type LogEntry struct {
	Timestamp  time.Time `json:"timestamp"`
	PID        int       `json:"pid"`
	Unit       string    `json:"unit"`
	Message    string    `json:"message"`
	Priority   int       `json:"priority"`
	Command    string    `json:"command"`
	Executable string    `json:"executable"`
	Hostname   string    `json:"hostname"`
	UserID     int       `json:"uid"`
	GroupID    int       `json:"gid"`
}

func ListLogsPage(limit int, cursor string) (*LogsPage, error) {
	j, err := sdjournal.NewJournal()
	if err != nil {
		return nil, err
	}
	defer j.Close()

	if cursor != "" {
		if err := j.SeekCursor(cursor); err != nil {
			return nil, err
		}
	} else {
		if err := j.SeekTail(); err != nil {
			return nil, err
		}
	}

	page := &LogsPage{
		Items: make([]LogEntry, 0, limit),
	}

	for len(page.Items) < limit {
		n, err := j.Previous()
		if err != nil {
			return nil, err
		}
		if n == 0 {
			break
		}

		e, err := j.GetEntry()
		if err != nil {
			return nil, err
		}

		page.Items = append(page.Items, LogEntry{
			Timestamp:  time.UnixMicro(int64(e.RealtimeTimestamp)),
			PID:        util.Atoi(e.Fields["_PID"]),
			Unit:       e.Fields["_SYSTEMD_UNIT"],
			Message:    e.Fields["MESSAGE"],
			Priority:   util.Atoi(e.Fields["PRIORITY"]),
			Command:    e.Fields["_COMM"],
			Executable: e.Fields["_EXE"],
			Hostname:   e.Fields["_HOSTNAME"],
			UserID:     util.Atoi(e.Fields["_UID"]),
			GroupID:    util.Atoi(e.Fields["_GID"]),
		})

		page.NextCursor = e.Cursor
	}

	// lookahead
	n, err := j.Previous()
	if err != nil {
		return nil, err
	}
	page.HasMore = n > 0

	return page, nil
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

	sort.Slice(values, func(i, j int) bool {
		return strings.ToLower(values[i]) < strings.ToLower(values[j])
	})
	return values, nil
}
