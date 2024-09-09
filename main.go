package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing game")
	fmt.Println("Um número inteiro aleatório entre 0 e 100 será sorteado. Tente acertar")

	x := rand.Int64N(101)
	x = 10
	scanner := bufio.NewScanner(os.Stdin)
	userGuesses := [10]int64{}

	for i := range userGuesses {
		fmt.Println("Qual o seu palpite?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("O seu palpite tem que ser um número inteiro")
			return
		}

		switch {
		case guessInt < x:
			fmt.Println("Você errou! O número é maior que ", guessInt)
		case guessInt > x:
			fmt.Println("Você errou! O número é menor que ", guessInt)
		case guessInt == x:
			fmt.Printf(
				"Parabéns, você acertou! O número era: %d\n"+
					"Você acertou em %d tentativas.\n"+
					"Essas foram as suas tentativas: %v\n",
				x, i+i, userGuesses[:i],
			)
			return
		}

		userGuesses[i] = guessInt
	}

	fmt.Printf(
		"Infelizmente, você não acertou o número, que era: %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foram as suas tentativas: %v\n",
		x, userGuesses,
	)
}
