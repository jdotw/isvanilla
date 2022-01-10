package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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

type TokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

func getToken() string {

	url := "https://isvanilla.au.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"B1AAk4waBvQCpm24kg2kNHE2CM5Nagry\",\"client_secret\":\"5MG88Y-SpfYrR9iaS68lRXKUFBxxJwcTtFcxW1WhUbo0JDlWMdJ6Tk6ngU00Legu\",\"audience\":\"https://api.syrupstock.io\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var j TokenResponse
	err = json.Unmarshal(body, &j)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("J: %+v\n", j.AccessToken)

	return j.AccessToken
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

	r := regexp.MustCompile(`(\d+) in stock`)
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

func tokenRequestEditor(t string) inventory.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer "+t)
		return nil
	}
}

func recordStockLevel(c *inventory.ClientWithResponses, token string, product product.Product) int {
	if product.Url != nil {

		l := stockLevel(*product.Url)
		r, err := c.CreateInventorySnapshotWithResponse(context.Background(), *product.VendorID, *product.ID, inventory.MutateInventorySnapshot{StockLevel: &l}, tokenRequestEditor(token))
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("RECORDED: %+v", r)
		return l
	} else {
		log.Printf("Product %v has no URL", product.ID)
		return 0
	}
}

func main() {

	t := getToken()
	fmt.Printf("TOKEN: %+v", t)

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

			l := recordStockLevel(ic, t, p)
			log.Printf("%v: %v", *p.Name, l)
		}
	}

	// for {
	// 	vanilla := recordStockLevel(ic, "gj-vanilla-syrup", "https://www.gloriajeanscoffees.com.au/product/vanilla-syrup-750ml/")
	// 	println("VANILLA: ", vanilla)

	// 	sfVanilla := recordStockLevel(ic, "gj-vanilla-sugar-free-syrup", "https://www.gloriajeanscoffees.com.au/product/vanilla-sugar-free-syrup-750ml/")
	// 	println("SF VANILLA: ", sfVanilla)

	// 	caramel := recordStockLevel(ic, "gj-caramel-syrup", "https://www.gloriajeanscoffees.com.au/product/caramel-syrup-750ml/")
	// 	println("CARAMEL: ", caramel)

	// 	sfCaramel := recordStockLevel(ic, "gj-caramel-sugar-free-syrup", "https://www.gloriajeanscoffees.com.au/product/caramel-sugar-free-syrup-750ml/")
	// 	println("SF CARAMEL: ", sfCaramel)

	// 	time.Sleep(5 * time.Minute)
	// }
}
