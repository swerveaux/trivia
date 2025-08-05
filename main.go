package main

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"math/rand"
	"strconv"
	"strings"
)

//go:embed birthdays.txt
var f embed.FS

const maxWeight = 10

type CelebAge struct {
	Name   string
	Age    int
	Weight int
}

func (c *CelebAge) String() string {
	return fmt.Sprintf("%s: %d", c.Name, c.Age)
}

func main() {
	celebAges := make([]CelebAge, 0)

	file, err := f.Open("birthdays.txt")
	if err != nil {
		panic(err)
	}
	defer func(fs fs.File) {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " - ")
		name := parts[0]
		age, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		celebAges = append(celebAges, CelebAge{name, age, maxWeight})
	}

	var ageInput string

	fmt.Println("Type 'quit' at any prompt to quit.")

	for {
		celebMatrix := populateMatrix(celebAges)
		celebAge := &celebMatrix[rand.Intn(len(celebMatrix))]
		fmt.Printf("%s: ", celebAge.Name)
		_, err := fmt.Scanf("%s", &ageInput)
		if err != nil {
			fmt.Println("What the heck did you type?")
			continue
		}

		if ageInput == "quit" {
			break
		}

		age, err := strconv.Atoi(ageInput)
		if err != nil {
			fmt.Printf("%s is not age valid ageInput\n", ageInput)
			continue
		}

		if age == celebAge.Age {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Sorry, %s is %d\n", celebAge.Name, celebAge.Age)
		}

		setWeight(celebAges, 0, celebAge.Name)
	}
	fmt.Println("Bye!")
}

func setWeight(celebAges []CelebAge, weight int, name string) {
	for i := range celebAges {
		if celebAges[i].Name == name {
			celebAges[i].Weight = weight
			break
		}
	}
}

func populateMatrix(celebAges []CelebAge) []CelebAge {
	celebMatrix := make([]CelebAge, 0, maxWeight*len(celebAges))
	for i := range celebAges {
		if celebAges[i].Weight < maxWeight {
			celebAges[i].Weight++
		}
		for j := 0; j < celebAges[i].Weight; j++ {
			celebMatrix = append(celebMatrix, celebAges[i])
		}
	}

	return celebMatrix
}
