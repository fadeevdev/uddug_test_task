package transaction

import (
	"testing"
	"time"
)

func TestGeneratePlot(t *testing.T) {
	tc := []struct {
		interval string
		test     []*Transaction
		want     []*Transaction
	}{
		{
			interval: "day",
			test: []*Transaction{
				{
					4456,
					time.Unix(1616026248, 0).UTC(),
				},
				{
					4231,
					time.Unix(1616022648, 0).UTC(),
				},
				{
					5212,
					time.Unix(1616019048, 0).UTC(),
				},
				{
					4321,
					time.Unix(1615889448, 0).UTC(),
				},
				{
					4567,
					time.Unix(1615871448, 0).UTC(),
				},
			},
			want: []*Transaction{
				{
					4456,
					time.Unix(1616025600, 0).UTC(),
				},
				{
					4231,
					time.Unix(1615939200, 0).UTC(),
				},
				{
					4321,
					time.Unix(1615852800, 0).UTC(),
				},
			},
		},
		{
			interval: "month",
			test: []*Transaction{
				{
					6483,
					time.Unix(1616343607, 0).UTC(),
				},
				{
					3283,
					time.Unix(1613498619, 0).UTC(),
				},
				{
					9865,
					time.Unix(1639748224, 0).UTC(),
				},
				{
					3471,
					time.Unix(1614426305, 0).UTC(),
				},
				{
					4186,
					time.Unix(1636823949, 0).UTC(),
				},
				{
					6783,
					time.Unix(1623360214, 0).UTC(),
				},
				{
					5983,
					time.Unix(1617795919, 0).UTC(),
				},
			},
			want: []*Transaction{
				{
					9865,
					time.Unix(1638316800, 0).UTC(),
				},
				{
					4186,
					time.Unix(1635724800, 0).UTC(),
				},
				{
					6783,
					time.Unix(1622505600, 0).UTC(),
				},
				{
					5983,
					time.Unix(1617235200, 0).UTC(),
				},
				{
					6483,
					time.Unix(1614556800, 0).UTC(),
				},
				{
					3471,
					time.Unix(1612137600, 0).UTC(),
				},
			},
		},
	}

	for _, testCase := range tc {
		test := testCase.test
		want := testCase.want

		_, err := GeneratePlot(test, "bob")
		if err == nil {
			t.Errorf("expected error on wrong interval input parameter")
			return
		}

		res, err := GeneratePlot(test, testCase.interval)
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		if len(res) != len(want) {
			t.Errorf("expected value is: %v, got: %v", want, res)
			return
		}

		for i, r := range res {
			if want[i].Value != r.Value || want[i].Timestamp != r.Timestamp {
				t.Errorf("expected value is: %v, got: %v", want[i], res[i])
			}
		}
	}

}
