package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Komputer losuje liczbę między 1 a 100
	secretNumber := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	attempts := 0

	fmt.Println("Zgadnij liczbę między 1 a 100!")

	for {
		fmt.Print("Podaj swoją liczbę: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("To nie jest prawidłowa liczba!")
			continue
		}

		attempts++
		if guess < secretNumber {
			fmt.Println("Liczba jest za mała!")
		} else if guess > secretNumber {
			fmt.Println("Liczba jest za duża!")
		} else {
			fmt.Printf("Gratulacje! Zgadłeś liczbę w %d próbach.\n", attempts)
			break
		}
	}
}
