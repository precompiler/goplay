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

	lights := [2]light{{"off"}, {"off"}}
	printLightsStatus(lights)
	turnAllLightsOn(lights)
	printLightsStatus(lights)

	turnAllLightsOnForce(&lights)
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
	for i := 0; i < len(lights); i++ {
		lights[i].Status = "on"
	}
}

func turnAllLightsOnForce(lights *[2]light) {
	//lights[0].Status = "on"
	for i := 0; i < len(lights); i++ {
		lights[i].Status = "on"
	}
}

type light struct{ Status string }
