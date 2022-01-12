package scraper

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/syrupstock/pkg/inventory"
	"github.com/jdotw/syrupstock/pkg/product"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
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

type Scraper struct {
	logger log.Factory
	tracer opentracing.Tracer

	token string
}

func NewScraper(logger log.Factory, tracer opentracing.Tracer) Scraper {
	s := Scraper{
		logger: logger,
		tracer: tracer,
	}

	s.getToken()

	return s
}

func (s *Scraper) getToken() {

	url := "https://isvanilla.au.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"B1AAk4waBvQCpm24kg2kNHE2CM5Nagry\",\"client_secret\":\"5MG88Y-SpfYrR9iaS68lRXKUFBxxJwcTtFcxW1WhUbo0JDlWMdJ6Tk6ngU00Legu\",\"audience\":\"https://api.syrupstock.io\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Bg().Fatal("failed to request token", zap.Error(err))
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var j TokenResponse
	err = json.Unmarshal(body, &j)
	if err != nil {
		s.logger.Bg().Fatal("failed to parse token response", zap.Error(err))
	}

	s.token = j.AccessToken
}

func (s *Scraper) getStockLevel(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		s.logger.Bg().Fatal("failed to get product page", zap.Error(err), zap.String("url", url))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Bg().Fatal("failed to read body", zap.Error(err), zap.String("url", url))
	}
	sb := string(body)

	r := regexp.MustCompile(`(\d+) in stock`)
	m := r.FindStringSubmatch(sb)
	if len(m) > 0 {
		i, err := strconv.Atoi(m[1])
		if err != nil {
			s.logger.Bg().Fatal("failed to convert stock level string to an integer", zap.Error(err), zap.String("url", url))
		}
		return i, nil
	} else {
		r := regexp.MustCompile(`Out of stock`)
		m := r.FindStringSubmatch(sb)
		if len(m) > 0 {
			s.logger.Bg().Info("out of stock", zap.String("url", url))
			return 0, nil
		} else {
			errStr := "failed to find stock level string"
			s.logger.Bg().Fatal(errStr, zap.String("url", url))
			return 0, errors.New(errStr)
		}
	}
}

func (s *Scraper) ScrapeStockLevel(c *inventory.ClientWithResponses, product product.Product) int {
	if product.Url != nil {
		reqEd := func(t string) inventory.RequestEditorFn {
			return func(ctx context.Context, req *http.Request) error {
				req.Header.Add("Authorization", "Bearer "+t)
				return nil
			}
		}

		l, err := s.getStockLevel(*product.Url)
		if err != nil {
			s.logger.Bg().Fatal("failed to get stock level", zap.Error(err))
		}

		_, err = c.CreateInventorySnapshotWithResponse(context.Background(),
			*product.VendorID,
			*product.ID,
			inventory.MutateInventorySnapshot{StockLevel: &l},
			reqEd(s.token))
		if err != nil {
			s.logger.Bg().Fatal("failed to create inventory snapshot", zap.Error(err))
		}

		return l
	} else {
		s.logger.Bg().Error("product has no URL", zap.String("productID", *product.ID))
		return 0
	}
}
