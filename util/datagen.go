package util

import (
	"math/rand"
	"sort"
)

func MonthsForYear(yearstr string) []string {
	// var months []string
	monthNames := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	// for _, month := range monthNames {
	// 	months = append(months, month+" "+yearstr)
	// }
	return monthNames
}

func RoomNightsPerMonth() []int {
	roomNights := make([]int, 12)
	for i := range roomNights {
		roomNights[i] = 200 + rand.Intn(1601) // 1601 because rand.Intn(n) generates a number from 0 to n-1
	}
	return roomNights
}

type HotelAndSavings struct {
	ShortName string
	Savings   float64
}

func SavingsLost() []HotelAndSavings {
	hotelNames := []string{
		"Marriott", "Hilton", "Hyatt", "Sheraton", "Westin",
		"Four Seasons", "Ritz-Carlton", "InterContinental",
		"Holiday Inn", "Best Western", "Radisson", "Wyndham",
		"Fairmont", "DoubleTree", "Crowne Plaza", "Embassy Suites",
		"Comfort Inn", "La Quinta", "Red Roof Inn", "Motel 6",
	}

	length := 5 + rand.Intn(5) // 13 because rand.Intn(n) generates a number from 0 to n-1
	hotelsAndSavings := make([]HotelAndSavings, length)
	for i := range hotelsAndSavings {
		hotelsAndSavings[i].Savings = 200 + rand.Float64()*(3000-200)
		hotelsAndSavings[i].ShortName = hotelNames[i]
	}
	sort.Slice(hotelsAndSavings, func(i, j int) bool {
		return hotelsAndSavings[i].Savings > hotelsAndSavings[j].Savings
	})

	return hotelsAndSavings
}

func HotelNamesOnly(hotelAndSavings []HotelAndSavings) []string {
	hotelNames := make([]string, len(hotelAndSavings))
	for i, hs := range hotelAndSavings {
		hotelNames[i] = hs.ShortName
	}
	return hotelNames
}

func SavingsOnly(hotelAndSavings []HotelAndSavings) []float64 {
	savings := make([]float64, len(hotelAndSavings))
	for i, hs := range hotelAndSavings {
		savings[i] = hs.Savings
	}
	return savings
}
