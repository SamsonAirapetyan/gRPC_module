package data

import (
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ExchangeRates struct {
	log  hclog.Logger
	rate map[string]float64
}

func NewExchangeRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rate: map[string]float64{}}
	er.getRate()

	return er, nil
}

func (e *ExchangeRates) GetRate(base, dest string) (float64, error) {
	br, ok := e.rate[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for cuurency %s", base)
	}
	dr, ok := e.rate[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for cuurency %s", dest)
	}
	return dr / br, nil
}
func (e *ExchangeRates) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates
				for k, v := range e.rate {
					// change can be 10% of original value
					change := (rand.Float64() / 10)
					// is this a postive or negative change
					direction := rand.Intn(1)

					if direction == 0 {
						// new value with be min 90% of old
						change = 1 - change
					} else {
						// new value will be 110% of old
						change = 1 + change
					}

					// modify the rate
					e.rate[k] = v * change
				}

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()

	return ret
}

func (e *ExchangeRates) getRate() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(&md)

	for _, c := range md.CobeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}
		e.rate[c.Currency] = r
	}
	e.rate["EUR"] = 1
	return nil
}

type Cubes struct {
	CobeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
