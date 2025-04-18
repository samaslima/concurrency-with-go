package main

import (
	"fmt"
	"search-engine/internal/fetcher"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var priceOne, priceTwo, priceThree float64

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		priceOne = fetcher.FetchPriceFomSiteOne()
	}()

	go func() {
		defer wg.Done()
		priceTwo = fetcher.FetchPriceFomSiteTwo()
	}()

	go func() {
		defer wg.Done()
		priceThree = fetcher.FetchPriceFomSiteThree()
	}()

	wg.Wait()

	fmt.Printf("R$ %.2f \n", priceOne)
	fmt.Printf("R$ %.2f \n", priceTwo)
	fmt.Printf("R$ %.2f \n", priceThree)

	fmt.Printf("\n time: %s", time.Since(start))
}
