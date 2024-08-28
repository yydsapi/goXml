package main

import (
	"sort"
)

type Pair struct {
	Key   string
	Value string
}

func Sort(m map[string]string) []Pair {
	pairs := make([]Pair, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Value > pairs[j].Value })
	return pairs
	//fmt.Println(pairs) // [{a 1} {b 2} {c 3} {d 4}]
}
