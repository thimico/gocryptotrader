//+build mock_test_off

// This will build if build tag mock_test_off is parsed and will do live testing
// using all tests in (exchange)_test.go
package poloniex

import (
	"log"
	"os"
	"testing"

	"github.com/idoall/gocryptotrader/config"
	"github.com/idoall/gocryptotrader/exchanges/sharedtestvalues"
)

var mockTests = false

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	err := cfg.LoadConfig("../../testdata/configtest.json", true)
	if err != nil {
		log.Fatal("Poloniex load config error", err)
	}
	poloniexConfig, err := cfg.GetExchangeConfig("Poloniex")
	if err != nil {
		log.Fatal("Poloniex Setup() init error", err)
	}
	poloniexConfig.API.AuthenticatedSupport = true
	poloniexConfig.API.Credentials.Key = apiKey
	poloniexConfig.API.Credentials.Secret = apiSecret
	p.SetDefaults()
	err = p.Setup(poloniexConfig)
	if err != nil {
		log.Fatal("Poloniex setup error", err)
	}
	log.Printf(sharedtestvalues.LiveTesting, p.Name, p.API.Endpoints.URL)
	os.Exit(m.Run())
}