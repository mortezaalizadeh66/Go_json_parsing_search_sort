package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Person struct {
	Name      string `json:"name"`
	Family    string `json:"family"`
	Timestamp string `json:"timestamp"`
}

func main() {
	// Generate 20 people with random names and families
	people := make([]Person, 20)
	families := []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Miller"}
	rand.Seed(time.Now().UnixNano())

	for i := range people {
		name := fmt.Sprintf("Person%d", i+1)
		family := families[rand.Intn(len(families))]
		timestamp := time.Now().Add(-time.Duration(rand.Intn(365*24*60*60)+rand.Intn(24*60*60)+rand.Intn(60*60)+rand.Intn(60)) * time.Second).Format("2006-01-02 15:04:05")

		people[i] = Person{Name: name, Family: family, Timestamp: timestamp}
	}

	// Save the people to a JSON file
	f, err := os.Create("people.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(people); err != nil {
		panic(err)
	}

	// Sort the data by timestamp and name
	sortedDataTime := make([]Person, len(people))
	copy(sortedDataTime, people)
	sort.Slice(sortedDataTime, func(i, j int) bool {
		return sortedDataTime[i].Timestamp < sortedDataTime[j].Timestamp
	})

	fmt.Println("Sorted by timestamp:")
	for _, p := range sortedDataTime {
		fmt.Println(p)
	}

	sortedDataName := make([]Person, len(people))
	copy(sortedDataName, people)
	sort.Slice(sortedDataName, func(i, j int) bool {
		return strings.ToLower(sortedDataName[i].Family) < strings.ToLower(sortedDataName[j].Family)
	})

	fmt.Println("Sorted by name:")
	for _, p := range sortedDataName {
		fmt.Println(p)
	}

	// Search for a name in the data
	var family string
	fmt.Print("Enter name to search for: ")
	fmt.Scanln(&family)

	for _, p := range people {
		if fuzzy.Match(strings.ToLower(family), strings.ToLower(p.Family)) {
			fmt.Println(p)
			return
		}
	}

	fmt.Printf("No record found for name '%s'\n", family)
}
