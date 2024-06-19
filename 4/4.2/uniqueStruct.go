package main

import (
	"fmt"
	"sort"
)

type structID struct {
	ID int
}

func uniqueID(s []structID) []structID {
	uniqueIDs := []structID{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i].ID != s[i-1].ID {
			uniqueIDs = append(uniqueIDs, s[i])
		}
	}
	return uniqueIDs
}

func main() {
	IDs := []structID{{6}, {3}, {2}, {5}, {9}, {4}, {2}, {6}, {2}}
	fmt.Println(IDs)

	sort.Slice(IDs, func(i, j int) bool {
		return IDs[i].ID < IDs[j].ID
	})

	uniqueIDs := uniqueID(IDs)
	fmt.Println(uniqueIDs)
}
