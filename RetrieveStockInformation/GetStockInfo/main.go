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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/puerkitoBio/goquery"
)

var p = fmt.Println
var f = fmt.Printf

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

func parseFinancePage(s string) string {

	doc, err := goquery.NewDocument(fmt.Sprintf("http://www.nasdaq.com/symbol/%s", s))
	if err != nil {
		log.Fatal(err)
	}

	p("scrape begin______________________________________________________________________")
	var heading string

	doc.Find(".overview-results .table-cell").Each(func(index int, item *goquery.Selection) {
		i := item.Text()
		// Removing linefeed(x'0A') and spaces(x'20')

		// if index is even, this is a text
		if index%2 == 0 {
			//heading = scrubText(i)
			heading = strings.TrimSpace(i)
		} else {
			//v := scrubText(i)
			v := strings.TrimSpace(i)
			p(heading, " ", v)
		}

		//f("index: %d i: %s Length: %d\n", index, i, len(i))

		// p("j:", j)

	})

	p("scrape end______________________________________________________________________")

	return ""
}
func scrubText(s string) string {
	// find first text char
	var b int
	var e int
	for j := 0; j < len(s); j++ {
		if s[j] == 32 || s[j] == 10 {
			continue
		} else {
			b = j - 1
			break
		}
	}
	// find last text char
	for k := len(s) - 1; k > -1; k-- {
		if s[k] == 32 || s[k] == 10 {
			continue
		} else {
			e = k + 1
			break
		}
	}
	// f("String: %X\n", s)
	// p(" Begin: ", b, " End: ", e, " Length: ", len(s), " Slice: ", s[b:e])
	return s[b:e]
}

//}
/*************************************************************************************
	M a i n


***************************************************************************************/

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

	keys = []string{"CLDT", "T"}

	baseURL := "https://api.iextrading.com/1.0/stock/market/batch"
	//symbols := []string{"ABBV", "T", "BEP", "IBM", "SEP", "VFINX"}
	url := fmt.Sprintf("%s?types=quote&symbols=%s&filter=week52Low,week52High,iexVolume,high,peRatio,latestPrice,sector,latestVolume,previousClose,low,iexLastUpdated,latestTime,close,open,changePercent,marketCap,symbol,change,companyName,ytdChange,latestUpdate", baseURL, strings.Join(keys, ","))

	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error ===>", err)
	}
	//fmt.Println("===>", res)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	type m3 map[string]interface{}
	type m2 map[string]m3

	m1 := make(map[string]m2)
	//var m1 interface{}
	json.Unmarshal(body, &m1)
	if err != nil {
		fmt.Println("JSON unmarshalling failed.....", err)
	}

	//fmt.Println(err)
	fmt.Println("==============================================================")
	//fmt.Println(s)
	fmt.Printf("%T ===> %v\n", m1, m1)
	fmt.Println("==============================================================")

	var symbol, companyName string
	var latestPrice, volume, dayHigh, dayLow, lastUpdated, pe, previousClose, week52Low, week52High, marketCap float64

	for _, k := range keys {
		if m1[k]["quote"]["symbol"] == nil {
			continue
		}
		symbol = m1[k]["quote"]["symbol"].(string)
		latestPrice = m1[k]["quote"]["latestPrice"].(float64)
		volume = m1[k]["quote"]["iexVolume"].(float64)
		dayHigh = m1[k]["quote"]["high"].(float64)
		dayLow = m1[k]["quote"]["low"].(float64)
		lastUpdated = m1[k]["quote"]["iexLastUpdated"].(float64)
		pe = -1
		if m1[k]["quote"]["peRatio"] != nil {
			pe = m1[k]["quote"]["peRatio"].(float64)

		}
		previousClose = m1[k]["quote"]["previousClose"].(float64)
		companyName = m1[k]["quote"]["companyName"].(string)
		week52Low = m1[k]["quote"]["week52Low"].(float64)
		week52High = m1[k]["quote"]["week52High"].(float64)
		marketCap = m1[k]["quote"]["marketCap"].(float64)

		fmt.Println("__________________________________________________")
		fmt.Println("Symbol: ", symbol)
		fmt.Println("latestPrice: ", latestPrice)
		fmt.Println("volume: ", volume)
		fmt.Println("dayHigh: ", dayHigh)
		fmt.Println("dayLow: ", dayLow)
		fmt.Println("lastUpdated: ", lastUpdated)
		fmt.Println("pe: ", pe)
		fmt.Println("previousClose: ", previousClose)
		fmt.Println("companyName: ", companyName)
		fmt.Println("week52Low: ", week52Low)
		fmt.Println("week52High: ", week52High)
		fmt.Println("marketCap: ", marketCap)
		fmt.Println("__________________________________________________`")

		// Retrieve Data from NASDAQ Finance Site
		var scrapedData string
		scrapedData = parseFinancePage(k)
		fmt.Println("Scraped Data: ", scrapedData)
	}
}
