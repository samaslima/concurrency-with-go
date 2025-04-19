package main

import (
	"fmt"
	"search-engine/internal/fetcher"
	"search-engine/internal/processor"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done

	fmt.Printf("\n time: %s", time.Since(start))
}
