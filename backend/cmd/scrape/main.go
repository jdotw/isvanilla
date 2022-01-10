package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/jdotw/syrupstock/pkg/inventory"
	"github.com/jdotw/syrupstock/pkg/product"
	"github.com/jdotw/syrupstock/pkg/vendor"
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

func recordStockLevel(c *inventory.ClientWithResponses, product string, url string) int {
	level := stockLevel(url)

	// c.CreateInventorySnapshot(context.Background(), vendorID, productID, &MutateInventorySnapshot{StockLevel: level})

	// s := StockRecord{
	// 	Product: product,
	// 	Stock:   level,
	// }

	return level
}

func main() {
	clientHost := "http://localhost:8080"

	vc, err := vendor.NewClientWithResponses(clientHost)
	if err != nil {
		log.Fatal(err)
	}

	pc, err := product.NewClientWithResponses(clientHost)
	if err != nil {
		log.Fatal(err)
	}

	ic, err := inventory.NewClientWithResponses(clientHost)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	vs, err := vc.GetVendorsWithResponse(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("VENDORS RESP: %+v", vs)
	log.Printf("VENDORS BODY: %+v", string(vs.Body))

	for i := 0; i < len(*vs.JSON200); i++ {
		v := (*vs.JSON200)[i]
		log.Printf("VENDOR: %+v", *v.Name)

		ps, err := pc.GetProductsWithResponse(ctx, *v.ID)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("PRODUCT RESP: %+v", ps)
		log.Printf("PRODUCT BODY: %+v", string(ps.Body))

		for i := 0; i < len(*ps.JSON200); i++ {
			p := (*ps.JSON200)[i]
			log.Printf("PRODUCT: %+v", *p.Name)
		}
	}

	for {
		vanilla := recordStockLevel(ic, "gj-vanilla-syrup", "https://www.gloriajeanscoffees.com.au/product/vanilla-syrup-750ml/")
		println("VANILLA: ", vanilla)

		sfVanilla := recordStockLevel(ic, "gj-vanilla-sugar-free-syrup", "https://www.gloriajeanscoffees.com.au/product/vanilla-sugar-free-syrup-750ml/")
		println("SF VANILLA: ", sfVanilla)

		caramel := recordStockLevel(ic, "gj-caramel-syrup", "https://www.gloriajeanscoffees.com.au/product/caramel-syrup-750ml/")
		println("CARAMEL: ", caramel)

		sfCaramel := recordStockLevel(ic, "gj-caramel-sugar-free-syrup", "https://www.gloriajeanscoffees.com.au/product/caramel-sugar-free-syrup-750ml/")
		println("SF CARAMEL: ", sfCaramel)

		time.Sleep(5 * time.Minute)
	}
}
