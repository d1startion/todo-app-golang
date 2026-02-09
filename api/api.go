package api

import (
	"errors"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("empty repeat")
	}

	start, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", errors.New("bad repeat")
	}

	switch parts[0] {
	case "y":
		d := start
		for {
			prev := d
			d = d.AddDate(1, 0, 0)
			if prev.Month() == 2 && prev.Day() == 29 && d.Month() == 2 && d.Day() == 28 {
				d = d.AddDate(0, 0, 1)
			}
			if d.After(now) {
				return d.Format(dateFormat), nil
			}
		}

	case "d":
		if len(parts) != 2 {
			return "", errors.New("bad d")
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil || n <= 0 || n > 400 {
			return "", errors.New("bad d")
		}
		d := start
		for {
			d = d.AddDate(0, 0, n)
			if d.After(now) {
				return d.Format(dateFormat), nil
			}
		}

	case "w":
		if len(parts) != 2 {
			return "", errors.New("bad w")
		}
		ok := map[int]bool{}
		for _, s := range strings.Split(parts[1], ",") {
			v, err := strconv.Atoi(s)
			if err != nil || v < 1 || v > 7 {
				return "", errors.New("bad w")
			}
			ok[v] = true
		}
		d := start
		for {
			d = d.AddDate(0, 0, 1)
			wd := int(d.Weekday())
			if wd == 0 {
				wd = 7
			}
			if ok[wd] && d.After(now) {
				return d.Format(dateFormat), nil
			}
		}

	case "m":
		if len(parts) < 2 {
			return "", errors.New("bad m")
		}

		days := []int{}
		negatives := []int{}
		for _, s := range strings.Split(parts[1], ",") {
			v, err := strconv.Atoi(s)
			if err != nil || v == 0 || v < -31 || v > 31 {
				return "", errors.New("bad m")
			}
			days = append(days, v)
			if v < 0 {
				negatives = append(negatives, v)
			}
		}

		nNeg := len(negatives)
		if nNeg > 1 {
			sort.Ints(negatives)
			if negatives[nNeg-1] != -1 {
				return "", errors.New("bad m")
			}
			for i := 1; i < nNeg; i++ {
				if negatives[i] != negatives[i-1]+1 {
					return "", errors.New("bad m")
				}
			}
		}

		months := map[int]bool{}
		if len(parts) == 3 {
			for _, s := range strings.Split(parts[2], ",") {
				v, err := strconv.Atoi(s)
				if err != nil || v < 1 || v > 12 {
					return "", errors.New("bad m")
				}
				months[v] = true
			}
		} else {
			for i := 1; i <= 12; i++ {
				months[i] = true
			}
		}

		d := start
		for i := 0; i < 4000; i++ {
			if d.After(now) && months[int(d.Month())] {
				last := time.Date(d.Year(), d.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
				for _, want := range days {
					day := want
					if want < 0 {
						day = last + want + 1
						if day < 1 {
							continue
						}
					}
					if day == d.Day() && day <= last {
						return d.Format(dateFormat), nil
					}
				}
			}
			d = d.AddDate(0, 0, 1)
		}

		return "", errors.New("no next date found in reasonable time")

	default:
		return "", errors.New("bad rule")
	}
}

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	if nowStr == "" {
		now = time.Now()
	} else {
		t, err := time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(w, "bad now", http.StatusBadRequest)
			return
		}
		now = t
	}

	res, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
