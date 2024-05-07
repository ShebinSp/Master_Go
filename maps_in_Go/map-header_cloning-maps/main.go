package main

import "fmt"

func main() {
	// How map is internally represented in memory

	// When declaring a map variable Go creates a pointer to a map header in memory,
	// The map references this internal data structure the map header.

	// When you copy a map to a new map, the internal data structure is not copied, just referenced.
	// Both maps have the same map header.
	friends := map[string]int{"Dan": 40, "Maria": 30}
	neighbors := friends
	friends["Dan"] = 50                       // the modification will reflect in neighbors map. neighbors is not
	fmt.Println("Map neighbors: ", neighbors) // a copy of friend, it references the same data structure
	//  in memeory.neighbors and friends have the same map header and when one changed, the other one also changed.

	// How to clone or duplicate a map
	// initialize a new map and then use a forloop to copy each element into the new map
	people := make(map[string]int)
	for key, value := range friends {
		people[key] = value
	}

	// now friends and people maps have different map header
	people["Annie"] = 40 // new entry belongs only to the map people
	fmt.Printf("people: %#v ----- friends: %#v\n",people,friends)

	delete(friends, "Dan")
	fmt.Println("map friends: ", friends)

	// it is not necessary to check if a key exist to delete.
	delete(friends, "John") // non existing key. It is not an error
	
	
}
