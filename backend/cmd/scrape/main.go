package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StockRecord struct {
	gorm.Model
	Product string
	Stock   int
}

func stockLevel(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	r := regexp.MustCompile("(\\d+) in stock")
	m := r.FindStringSubmatch(sb)
	if len(m) > 0 {
		i, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatalln(err)
		}
		return i
	} else {
		return 0
	}
}

func recordStockLevel(db *gorm.DB, product string, url string) int {
	level := stockLevel(url)

	s := StockRecord{
		Product: product,
		Stock:   level,
	}

	r := db.Create(&s)
	if r.Error != nil {
		log.Fatalln(r.Error)
	}

	return level
}

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(StockRecord{})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		vanilla := recordStockLevel(db, "gj-vanilla-syrup", "https://www.gloriajeanscoffees.com.au/product/vanilla-syrup-750ml/")
		println("VANILLA: ", vanilla)

		sfVanilla := recordStockLevel(db, "gj-vanilla-sugar-free-syrup", "https://www.gloriajeanscoffees.com.au/product/vanilla-sugar-free-syrup-750ml/")
		println("SF VANILLA: ", sfVanilla)

		caramel := recordStockLevel(db, "gj-caramel-syrup", "https://www.gloriajeanscoffees.com.au/product/caramel-syrup-750ml/")
		println("CARAMEL: ", caramel)

		sfCaramel := recordStockLevel(db, "gj-caramel-sugar-free-syrup", "https://www.gloriajeanscoffees.com.au/product/caramel-sugar-free-syrup-750ml/")
		println("SF CARAMEL: ", sfCaramel)

		time.Sleep(5 * time.Minute)
	}
}
