// main.go - Get Stock Prices

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

// Sorting keys in a map

func keysort() {
	m := map[string]int{"Tim": 3, "Hannah": 1, "Ricky": 6}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

func processCsvFile(csvFile string, m map[string]string, t map[string]float64) {
	/* reads a CSV file from Vanguard and updates the prices
	found in the map
	*/
	var ctr int
	var acs, ac string

	file1, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		var record = strings.Split(scanner.Text(), ",")
		// This is the section for stock totals
		if len(record) == 7 && record[2] != "Symbol" {

			symbol := record[2]
			price := record[4]
			m[symbol] = price
			//fmt.Println(symbol, price)
			if record[2] == "VMFXX" { // settlement record
				tot, _ := strconv.ParseFloat(record[5], 64)
				t[getAccName(record[0])] = tot

			}
			// Section for Detailed Transactions
			/*****************************************************
			 * 0-Account
			 * 1-Trade Date
			 * 2-Settlement Date
			 * 3-Transaction Type
			 * 4-Transaction Description
			 * 5-Investment Name
			 * 6-Symbol
			 * 7-Shares
			 * 8-Share Price
			 * 9-Principal Amount
			 *10-Commission Fees
			 *11-Net Amount
			 *12-Accrued Interest
			 *13-Account Type
			*******************************************************/

		} else if len(record) == 15 && record[1] != "Trade Date" {
			if acs != record[1] {
				ac = getAccName(record[0])
			}
			// date is in format mm/dd/yyyy. Need it in yyyy-mm-dd format
			settleDate := record[2]
			//fmt.Println("SettleDate: ", settleDate)
			sp := strings.Split(settleDate, "/")
			settleDatex := fmt.Sprintf("%s-%s-%s", sp[2], sp[0], sp[1])
			transType := record[3]

			if diffDays(settleDatex) < 8 && (transType == "Dividend" || transType == "Reinvestment" || transType == "Partnership") {
				acs = ac
				//tradeDate := record[1]
				//transDesc := record[4]
				investName := record[5]
				sym := record[6]
				noShares := record[7]
				net := record[11]

				fmt.Printf("%5s %s %12s %5s %35s %7s %5s\n", ac, settleDate, transType, sym, investName, net, noShares)
			}
		}
		ctr++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func copyToCsvFile(m map[string]string) {
	fmt.Println("==============================================================>")
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	f, err := os.Create("C:/data/personal/excel/stockprices.csv")
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)

	for _, k := range keys {
		fmt.Println(k, m[k])
		str := fmt.Sprintf("%s,%s\n", k, m[k])
		_, err := w.WriteString(str)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}

func scrub(m map[string]string, p string) {
	/*
	   /  scrub - cleans up any untidy data
	*/
	// add price for pogrx
	m["POGRX"] = p
	// fix RDS B to RDS.B
	holdPrice := m["RDS B"]
	delete(m, "RDS B")
	m["RDS.B"] = holdPrice
}

func getAccName(a string) string {
	//fmt.Printf("\nAccount Number: %s, Last 4: %s", a, a[4:])
	switch a[4:] {
	case "3856":
		return "SReg"
	case "1600":
		return "SSep"
	case "0761":
		return "SRoth"
	case "9332":
		return "DReg"
	case "8015":
		return "DRoth"
	}
	return fmt.Sprintf("Not Defined Account: %s", a)
}

func printSummary(m map[string]float64) {
	// Sort totals map and print
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("\nSettlement Account Summary")
	for _, k := range keys {
		fmt.Printf("%s %.2f\n", k, m[k])
	}
}
func diffDays(s string) int {
	now := time.Now()
	format := "2006-01-02 15:04:05"
	//fmt.Println(now)
	then, _ := time.Parse(format, fmt.Sprintf("%s %s", s, "11:58:06"))
	diff := now.Sub(then)
	//fmt.Println(then)
	return int(diff.Hours() / 24)
}

/*
	Main Routine
*/

func main() {
	var pogrxPrice string
	tickers := make(map[string]string)
	totals := make(map[string]float64)
	//println(diffDays("2018-05-31"))
	fmt.Print("Enter today's price for Odyssey Growth Fund(POGRX): ")
	fmt.Scan(&pogrxPrice)
	fmt.Println(pogrxPrice)
	getTickers(tickers)

	processCsvFile("C:/users/stephen/downloads/ofxdownload.CSV", tickers, totals)
	processCsvFile("C:/users/stephen/downloads/ofxdownload(1).CSV", tickers, totals)

	scrub(tickers, pogrxPrice)

	copyToCsvFile(tickers)
	printSummary(totals)
}
