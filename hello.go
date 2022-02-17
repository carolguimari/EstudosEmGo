package main

import "fmt"

func main() {
	heros := make(map[string][]string)

	heros["Marvel"] = append(heros["Marvel"], "Hulk", "Homem Aranha", "Dr. Estranho", "Viúva Negra")
	heros["DC"] = append(heros["DC"], "Batman", "Superman", "Mulher Maravilha", "Flash")

    fmt.Println(heros)

	superHero := heros["DC"][2]

	powerLevel := 4

	if powerLevel > 5 {
		fmt.Printf("Super herói %s é forte \n", superHero)
	} else {
		fmt.Printf("Super herói %s é fraco(a) \n", superHero)
	}
}