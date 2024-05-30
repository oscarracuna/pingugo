package main

import (
	"encoding/csv"
	"fmt"
	"os"
	//probing "github.com/prometheus-community/pro-bing"
)

type Record struct {
	LAN string
	CER string
	PER string
}

func main() {
	file, err := os.Open("scheme.csv")
	if err != nil {
		fmt.Println("Unabled to open file", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Unable to read file", err)
		return
	}

	data := make(map[string]Record)

	for _, row := range records[1:] {
		if len(row) < 4 {
			fmt.Println("Unable to get row", row)
			continue
		}
		cap := row[0]
		lan := row[1]
		cer := row[2]
		per := row[3]
		data[cap] = Record{LAN: lan, CER: cer, PER: per}
	}

	var cap string
	fmt.Print("Enter the cap number: ")
	fmt.Scan(&cap)
	if record, found := data[cap]; found {
		fmt.Printf("For cap %s: lan = %s, cer = %s, per = %s\n", cap, record.LAN, record.CER, record.PER)
	} else {
		fmt.Printf("Cap %s not found\n", cap)
	}
}
