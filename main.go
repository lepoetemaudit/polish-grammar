package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lepoetemaudit/polishgrammar/polish"
	"math/rand"
	"net/http"
	"os"
	"time"
	"github.com/eknkc/amber"
)

var templates = amber.MustCompileDir("templates", amber.DefaultDirOptions, amber.DefaultOptions)

func GetRandomDate(c *gin.Context) {
	year := 1400 + rand.Intn(750)
	day := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	day = day.AddDate(0, 0, rand.Intn(365))

	polishYear, err := polish.GetPolishYear(year, true)
	if err != nil {
		polishYear = err.Error()
	}

	polishDay, _ := polish.GetPolishDate(day.Day(), int(day.Month()), true)

	c.JSON(200, map[string]interface{}{
		"year":        year,
		"month":       day.Month(),
		"day":         day.Day(),
		"polish_year": polishYear + " roku",
		"polish_day":  polishDay,
	})
}

func index(c *gin.Context) {
	templates["index"].Execute(c.Writer, nil)
}

func main() {
	router := gin.Default()
	router.Static("/static", "static")

	router.Use(func(c *gin.Context) {
		templates = amber.MustCompileDir("templates", amber.DefaultDirOptions, amber.DefaultOptions)
	})

	router.GET("/", index)
	// A poor man's template reloading middleware


	api := router.Group("api/1")
	{
		api.GET("random_date", GetRandomDate)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":"+port, router)
}
