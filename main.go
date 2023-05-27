package main

import "GolangSnippets/codewarsSolutions"

func main() {
	// url_shortener.UrlShortener()
	// concurrency.ConcurrencySample()
	testArray := [][]int{
		{1, 2, 3, 1, 2, 3},
		{4, 5, 6, 1, 2, 3},
		{7, 8, 9, 1, 2, 3}}
	codewarsSolutions.RotatingArray(testArray)
}
