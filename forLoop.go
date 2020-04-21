package main

import (
	"fmt"
)

func main() {
	var i int // i available outside for-loop
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			continue
		}
		fmt.Println("... continuing ...")
	}

	var j int          // j available outide loop
	for ; j < 7; j++ { // blank statements in post-clause syntax are fine
		fmt.Println(j)
	}

	for i := 0; i < 7; i++ { // i only avialable inside loop
		fmt.Println(i)
	}

	// ugly infinite loop syntax
	var k int
	for { // actually for ; ; with no statememts makes it infinite. The ; can be omitted.
		if k == 5 {
			break // as soon as k is 5, leave that loop
		}
		fmt.Println(k)
		k++
	}

	// looping through a collection in post-clause syntax
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// indexed way of the above syntax
	slice2 := []int{12, 22, 23}
	for i, v := range slice2 { // by defining a range of a collection type i and v are set to index and value at that index
		fmt.Println(i, v)
	}

	// try this with a map
	wellKnownPorts := map[string]int{"http": 80, "https": 443} // map of strings to ints
	for _, v := range wellKnownPorts {                         // reaching into the maps values by wildcarding the index
		// panic("Doom and decay came upon us ...")			   // Code panicks here
		fmt.Println(v) // index can now safely be left out without compiler error
	}

}
