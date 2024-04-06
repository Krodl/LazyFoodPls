package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

const RecipeTitleDivElement = "div.jeetYO"
const RecipeTitleH1Element = "h1.ceYciq"
const RecipeTitleH2Element = "h2.jECFqG"
const IngredientDivElement = "div.frRfTC"
const IngredientNameElement = "p.fLfTya"
const IngredientMeasurementElement = "p.bNkKoC"

type Recipe struct {
	title       string
	url         string
	ingredients []Ingredient
}

type Ingredient struct {
	measurement, name string
}

func Scrape(desiredRecipes []string) []Recipe {
	var recipes []Recipe
	var recipe Recipe

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		recipe.url = r.URL.String()
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML(RecipeTitleDivElement, func(e *colly.HTMLElement) {
		var title string
		h1 := e.ChildText(RecipeTitleH1Element)
		h2 := e.ChildText(RecipeTitleH2Element)
		title = fmt.Sprintf("%s %s", h1, h2)
		recipe.title = title
	})

	c.OnHTML(IngredientDivElement, func(e *colly.HTMLElement) {
		ingredient := Ingredient{}
		ingredient.measurement = e.ChildText(IngredientMeasurementElement)
		ingredient.name = e.ChildText(IngredientNameElement)
		recipe.ingredients = append(recipe.ingredients, ingredient)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	for _, v := range desiredRecipes {
		recipe = Recipe{"", "", nil}
		c.Visit(v)
		recipes = append(recipes, recipe)
	}

	return recipes
}
