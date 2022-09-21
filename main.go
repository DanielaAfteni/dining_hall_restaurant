package main

import (
	// Package json implements encoding and decoding of JSON.
	// The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.
	"encoding/json"
	// Package ioutil implements some I/O utility functions.
	"io/ioutil"
	// Package os provides a platform-independent interface to operating system functionality.
	"os"
	// importing the gin, because is a high-performance HTTP web framework written in Golang (Go).
	"github.com/DanielaAfteni/dining_hall_restaurant/basicproces"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// setting the configuration
	scfg := config()
	basicproces.SettingtheConfig(scfg)
	//we create the menu, by calling the corresponding function
	menu := basicproces.GetMenu()
	// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
	// we create a new order channel
	newOrderChan := make(chan basicproces.Order)
	// we create a rating channel
	ratingChan := make(chan int)
	// we make the table channels
	tablesChans := make([]chan basicproces.Order, 0)
	// as well as we make the waiter channels
	waitersChans := make([]chan basicproces.Distribution, 0)
	// we are looping over the number of tables (which is 10)
	for i := 0; i < scfg.NrOfTables; i++ {
		// we set a new variable table
		table := basicproces.NewTable(i, menu, newOrderChan, ratingChan)
		// we add to the table chans the received chan
		tablesChans = append(tablesChans, table.ReceiveChan)
		go table.Run()
	}
	// we are looping over the number of waiters (which is 4)
	for i := 0; i < scfg.NrOfWaiters; i++ {
		waiter := basicproces.NewWaiter(i, newOrderChan, tablesChans)
		// we add to the waiters chans the waiter distribution channel
		waitersChans = append(waitersChans, waiter.DistributionChan)
		go waiter.Run()
	}
	// call the function for making an average rating
	go rating(ratingChan)
	// Gin is a high-performance HTTP web framework written in Golang (Go).
	r := gin.Default()
	r.POST("/distribution", func(c *gin.Context) {
		var distribution basicproces.Distribution
		if err := c.ShouldBindJSON(&distribution); err != nil {
			log.Err(err).Msg("Error binding JSON")
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		waiterId := distribution.WaiterId
		waitersChans[waiterId] <- distribution
		c.JSON(200, gin.H{"message": "Order served"})
	})
	r.Run()
}

// function for making an average rating
func rating(ratingChan <-chan int) {
	// taking into account that we have number of ratings, set initially as 0
	nrOfRatings := 0
	// taking into account that we have total rating, set initially as 0
	totalRating := 0
	// infinit loop
	for {
		// we are going to take the
		rating := <-ratingChan
		nrOfRatings++
		totalRating += rating
		log.Info().Int("Rating", rating).Float64("avgRating", float64(totalRating)/float64(nrOfRatings)).Msg("Received rating")
	}
}

func config() basicproces.Config {
	// Output writes the output for a logging event
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	// To open and read the json file
	file, err := os.Open("config/scfg.json")
	// in case of an error, it returns a message
	if err != nil {
		log.Fatal().Err(err).Msg("Error appeared at opening menu.json. Try to find it.")
	}
	// close the file
	defer file.Close()
	// read the data from the file and return the data
	byteValue, _ := ioutil.ReadAll(file)
	//
	var scfg basicproces.Config
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by scfg.
	json.Unmarshal(byteValue, &scfg)

	return scfg
}
