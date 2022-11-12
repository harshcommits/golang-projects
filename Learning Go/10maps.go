package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string) //format for maps: map[key_datatype]value_datatype; eg. map[int]string
	fmt.Println(states)
	states["RJ"] = "Rajasthan"
	states["DL"] = "Delhi"
	states["BR"] = "Bihar"
	fmt.Println(states) //sorts by key by default

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println("keys: ", keys)

	//print states
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

	delhi := states["DL"]
	fmt.Println(delhi)

	delete(states, "BR")
	fmt.Println(states)
}
