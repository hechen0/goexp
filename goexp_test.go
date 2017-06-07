package main

import (
	"testing"
	"log"
)

func TestFrequencySort(t *testing.T){
	catch := func(input, want string ){
		if got := FrequencySort(input); got != want {
			log.Fatalf("want: %s, got: %s", want, got)
		}
	}

	catch("tree", "eert")
}

