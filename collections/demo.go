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

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	telescopeView := planets[0:3]
	fmt.Println(telescopeView)
	telescopeView[0] = "destroyed"
	x := append(telescopeView, "planet-x")
	fmt.Println(x)
	fmt.Println(planets)

	originalSlice := make([]string, 5, 10)
	originalSlice = append(originalSlice, "0", "1", "2", "3")
	slice1 := append(originalSlice, "4")
	slice2 := append(originalSlice, "4", "5")
	fmt.Println(originalSlice)
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice2[1] = "x"
	fmt.Println(originalSlice)
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice1[1] = "y"
	fmt.Println(originalSlice)
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := []string{"a", "b", "c"}
	prefix(slice3)
	fmt.Println(slice3)

	args := []string{"apply", "-f", "test.yaml"}
	kubectl(args...)
	fmt.Println(args)
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

func prefix(slice []string) {
	for idx, _ := range slice {
		slice[idx] = "p_" + slice[idx]
	}
}

func kubectl(args ...string) {
	fmt.Printf("%d args received\n", len(args))
	fmt.Println(args)
	args[0] = "delete"
}
