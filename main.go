package main

import "fmt"

func main() {
	desiredRecipes := GetDesiredRecipes()
	recipes := Scrape(desiredRecipes)
	fmt.Println(recipes)
	SendEmail(recipes)
}
