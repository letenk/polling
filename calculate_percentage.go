package main

import (
	"fmt"
	"math"
)

type Option struct {
	Id      uint64
	Title   string
	Votes   uint64
	Percent float64
}

func CalculatePercentage(options []Option) uint64 {
	totalVotes := uint64(0)
	for _, option := range options {
		totalVotes += option.Votes
	}

	// Convert votes to percentage
	var totalRoundedPercent float64
	if totalVotes != 0 {
		for i, option := range options {
			percent := (float64(option.Votes) / float64(totalVotes)) * 100
			roundedPercent := math.Round(percent) // Rounded
			options[i].Percent = roundedPercent
			totalRoundedPercent += roundedPercent
		}
	}

	// Distribute rounding difference
	diff := 100 - totalRoundedPercent

	for i := range options {
		if options[i].Votes > 0 {
			options[i].Percent += diff
			break
		}
	}

	// Print result
	var totalPercentage uint64
	for _, option := range options {
		totalPercentage += uint64(option.Percent)
		fmt.Printf("%s: %.0f%%\n", option.Title, option.Percent)
	}

	fmt.Println("==========")
	fmt.Printf("Total Percentage: %d%%\n", totalPercentage)

	return totalPercentage
}

func main() {

	options := []Option{
		{
			Id:    1,
			Title: "Ikan",
			Votes: 1,
		},
		{
			Id:    2,
			Title: "Buah",
			Votes: 8,
		},
		{
			Id:    3,
			Title: "Sayur",
			Votes: 0,
		},
		{
			Id:    4,
			Title: "Daging",
			Votes: 0,
		},
		{
			Id:    5,
			Title: "Ayam",
			Votes: 0,
		},
	}

	// Count total votes
	CalculatePercentage(options)
}
