package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	pokecache "github.com/raainshe/pokedexcli/internal"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "blue greeN Yellow  ",
			expected: []string{"blue", "green", "yellow"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		fmt.Println("Length of actual: " + strconv.Itoa(len(actual)))
		fmt.Printf("%v\n", actual)
		fmt.Println("Length of expected: " + strconv.Itoa(len(c.expected)))
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match")
			t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Got '%s' but expected '%s'", word, expectedWord)
				t.Fail()
			}
		}
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
			val: []byte("testdata"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Second
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area?offset=0&limit=20", []byte("testdata"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
