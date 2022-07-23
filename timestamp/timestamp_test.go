package timestamp

import (
	"testing"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func TestTruncateMinutes(t *testing.T) {
	test, err := time.Parse(timeLayout, "2022-05-30 12:38:51")
	if err != nil {
		t.Error(err)
	}
	want, err := time.Parse(timeLayout, "2022-05-30 12:00:00")
	if err != nil {
		t.Error(err)
	}
	res := TruncateMinutes(test)
	if res != want {
		t.Errorf("expected value is: %v", want)
	}
}

func TestTruncateHours(t *testing.T) {
	test, err := time.Parse(timeLayout, "2022-05-30 12:38:51")
	if err != nil {
		t.Error(err)
	}
	want, err := time.Parse(timeLayout, "2022-05-30 00:00:00")
	if err != nil {
		t.Error(err)
	}
	res := TruncateHours(test)
	if res != want {
		t.Errorf("expected value is: %v, got: %v", want, res)
	}
}

func TestTruncateWeek(t *testing.T) {

	tc := []struct {
		test string
		want string
	}{
		{
			test: "2022-01-01 12:38:51",
			want: "2021-12-27 00:00:00",
		},
		{
			test: "2022-07-23 12:38:51",
			want: "2022-07-18 00:00:00",
		},
		{
			test: "2022-07-25 19:24:45",
			want: "2022-07-25 00:00:00",
		},
	}

	for _, c := range tc {
		test, err := time.Parse(timeLayout, c.test)
		if err != nil {
			t.Error(err)
		}
		want, err := time.Parse(timeLayout, c.want)
		if err != nil {
			t.Error(err)
		}
		res := TruncateWeek(test)
		if res != want {
			t.Errorf("expected value is: %v, got: %v", want, res)
		}
	}

}

func TestTruncateDays(t *testing.T) {
	test, err := time.Parse(timeLayout, "2022-05-30 12:38:51")
	if err != nil {
		t.Error(err)
	}
	want, err := time.Parse(timeLayout, "2022-05-01 00:00:00")
	if err != nil {
		t.Error(err)
	}
	res := TruncateDays(test).UTC()
	if res != want {
		t.Errorf("expected value is: %v, got: %v", want, res)
	}
}
