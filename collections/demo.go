package main

import (
	"fmt"
	"strings"
)

func main() {
	var alphabets [26]string
	alphabets[0] = "a"
	alphabets[1] = "b"
	fmt.Println(alphabets[1])
	fmt.Println(alphabets[3] == "")
	fmt.Println(len(alphabets))

	trafficLights := [3]string{"red", "green"}
	fmt.Println(trafficLights)

	compilerDesidedLen := [...]string{"a", "b"}
	fmt.Println(compilerDesidedLen)

	printAlphabets(alphabets)
	toUpperCase(alphabets)
	printAlphabets(alphabets)

	lights := [2]light{{"off"}, light{"off"}}
	printLightsStatus(lights)
	turnAllLightsOn(lights)
	printLightsStatus(lights)
}

func toUpperCase(alphabets [26]string) {
	for idx, alphabet := range alphabets {
		alphabets[idx] = strings.ToUpper(alphabet)
	}
}

func printAlphabets(alphabets [26]string) {
	for idx, alphabet := range alphabets {
		fmt.Println(idx, alphabet)
	}
}

func printLightsStatus(lights [2]light) {
	for idx, l := range lights {
		fmt.Println(idx, l)
	}
}

func turnAllLightsOn(lights [2]light) {
	for _, l := range lights {
		l.Status = "on"
	}
}

type light struct{ Status string }
