package main

/* GetStockInfo - Retrieve info about stocks in a similar format to the old yahoo finance apu
Scrape from nasdaq.com information related to dividends
Dividend Amount
Ex Date
Pay Date
52 Week Low
52 Week High
Market Cap

also retrieves information from iex trading.com


*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getTickers(m map[string]string) {
	/* get tickers from file stockprices.csv
	creates a map entry for each ticker retrieved
	*/
	var ctr int

	file, err := os.Open("c:/data/personal/excel/stockprices.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var ticker = strings.Split(scanner.Text(), ",")[0]
		m[ticker] = "0"
		ctr++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	tickers := make(map[string]string)
	getTickers(tickers)

	keys := make([]string, 0, len(tickers))
	for k := range tickers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, tickers[k])
	}
}
