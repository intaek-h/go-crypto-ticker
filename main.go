package main

import (
	"fmt"
	"sync"

	"github.com/intaek-h/go-crypto-ticker/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)

		go func(c string) {
			getCurrencyData(c)
			wg.Done()
		}(currency)
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("1 %v is %v USD\n", currency, rate.Price)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}
