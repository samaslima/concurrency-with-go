package main

import (
	"fmt"
	"search-engine/internal/fetcher"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for price := range priceChannel {
			fmt.Printf("price received: R$ %.2f \n", price)
		}
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFomSiteOne()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFomSiteTwo()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFomSiteThree()
	}()

	wg.Wait()
	close(priceChannel)

	fmt.Printf("\n time: %s", time.Since(start))
}
