package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- float64) {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFomSiteOne()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFomSiteTwo()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFomSiteThree()
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchPriceFomSiteOne() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFomSiteTwo() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFomSiteThree() float64 {
	time.Sleep(2 * time.Second)
	return rand.Float64() * 100
}
