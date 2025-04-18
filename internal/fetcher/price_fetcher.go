package fetcher

import (
	"math/rand"
	"time"
)

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
