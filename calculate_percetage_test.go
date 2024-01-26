package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePercentage(t *testing.T) {
	testCases := []struct {
		testName                string
		expectedTotalPercentage uint64
		options                 []Option
	}{
		{
			testName:                "test_1",
			expectedTotalPercentage: 100,
			options: []Option{
				{
					Id:    1,
					Title: "Ikan",
					Votes: 1,
				},
				{
					Id:    2,
					Title: "Buah",
					Votes: 5,
				},
				{
					Id:    3,
					Title: "Sayur",
					Votes: 80,
				},
				{
					Id:    4,
					Title: "Daging",
					Votes: 1000,
				},
				{
					Id:    5,
					Title: "Ayam",
					Votes: 876,
				},
			},
		},
		{
			testName:                "test_2",
			expectedTotalPercentage: 100,
			options: []Option{
				{
					Id:    1,
					Title: "Ikan",
					Votes: 287,
				},
				{
					Id:    2,
					Title: "Buah",
					Votes: 0,
				},
				{
					Id:    3,
					Title: "Sayur",
					Votes: 97,
				},
				{
					Id:    4,
					Title: "Daging",
					Votes: 378,
				},
				{
					Id:    5,
					Title: "Ayam",
					Votes: 6,
				},
			},
		},
		{
			testName:                "test_3",
			expectedTotalPercentage: 100,
			options: []Option{
				{
					Id:    1,
					Title: "Ikan",
					Votes: 8,
				},
				{
					Id:    2,
					Title: "Buah",
					Votes: 0,
				},
				{
					Id:    3,
					Title: "Sayur",
					Votes: 1,
				},
				{
					Id:    4,
					Title: "Daging",
					Votes: 487,
				},
				{
					Id:    5,
					Title: "Ayam",
					Votes: 877,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := CalculatePercentage(tc.options)
			assert.Equal(t, tc.expectedTotalPercentage, result)
		})
	}
}
