package main

/* currently some boilerplate for linking up graphql, gin, go */

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/gin-gonic/gin"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

var products []Product

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func initProductsData(p *[]Product) {
	product1 := Product{ID: 1, Name: "Chicha Morada", Info: "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)", Price: 7.99}
	product2 := Product{ID: 2, Name: "Chicha de jora", Info: "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)", Price: 5.95}
	product3 := Product{ID: 3, Name: "Pisco", Info: "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)", Price: 9.95}
	*p = append(*p, product1, product2, product3)
}

func main() {
	// Primary data initialization
	initProductsData(&products)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			result := executeQuery(c.Query("query"), schema)
			c.JSON(http.StatusOK, result)
		})
	}

	// Start and run the server
	router.Run(":8000")
}
