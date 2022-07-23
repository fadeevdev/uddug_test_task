package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fadeevdev/uddug_test_task/transaction"
)

type UnixTimestamp struct {
	time.Time
}

type UnixTimestampTransaction struct {
	Value     int           `json:"value"`
	Timestamp UnixTimestamp `json:"timestamp"`
}

func main() {

	help := flag.Bool("help", false, "Show help")
	var inputfile string
	var interval string
	var outfile string

	flag.StringVar(&inputfile, "inputfile", "in.json", "Input filename")
	flag.StringVar(&interval, "interval", "month", "Interval")
	flag.StringVar(&outfile, "outfile", "out.json", "Interval")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	jf, err := os.Open(fmt.Sprintf("%s", inputfile))
	if err != nil {
		fmt.Printf("Error opening %s file: %v", inputfile, err.Error())
		return
	}

	defer jf.Close()

	byteValue, err := ioutil.ReadAll(jf)

	if err != nil {
		fmt.Printf("Error reading %s file: %v", inputfile, err.Error())
		return
	}

	uts := []*UnixTimestampTransaction{}

	json.Unmarshal(byteValue, &uts)

	ts := make([]*transaction.Transaction, len(uts))

	for i, t := range uts {
		time := time.Unix(t.Timestamp.UTC().Unix(), 0)
		ts[i] = &transaction.Transaction{
			Value:     t.Value,
			Timestamp: time.UTC(),
		}
	}

	res, err := transaction.GeneratePlot(ts, fmt.Sprint(interval))
	if err != nil {
		fmt.Println(err)
		return
	}

	utsToMarshal := make([]*UnixTimestampTransaction, len(res))

	for i, r := range res {
		u := UnixTimestampTransaction{
			Value:     r.Value,
			Timestamp: UnixTimestamp{Time: time.Unix(r.Timestamp.UTC().Unix(), 0).UTC()},
		}
		utsToMarshal[i] = &u
	}

	file, err := json.MarshalIndent(utsToMarshal, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(outfile, file, 0644)
	if err != nil {
		fmt.Printf("Error opening %s file: %v", outfile, err.Error())
		return
	}

	fmt.Printf("Data processed and saved in: %s\n", outfile)

}

//Implement JSON Unmarshar/Marshal Interface for UnixTimestamp
func (u *UnixTimestamp) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}

func (u UnixTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", (u.Time.UTC().Unix()))), nil
}
