package sorting

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// IsSorted vérifie si la pile A est déjà triée.
func IsSorted(a *stack.Stack) bool {
	for i := 0; i < len(a.Numbers)-1; i++ {
		if a.Numbers[i] > a.Numbers[i+1] {
			return false
		}
	}
	return true
}

// rangs calcule le rang de chaque nombre.
// Le rang correspond au nombre de valeurs strictement plus petites.
func rangs(numbers []int) []int {
	result := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				result[i]++
			}
		}
	}

	return result
}

// Sort choisit la méthode selon le nombre de nombres.
func Sort(a, b *stack.Stack) []string {
	if IsSorted(a) {
		return nil
	}

	a.Numbers = rangs(a.Numbers)

	var instructions []string

	switch len(a.Numbers) {
	case 2:
		if a.Numbers[0] > a.Numbers[1] {
			operations.Apply("sa", a, b)
			instructions = append(instructions, "sa")
		}

	case 3:
		trierTrois(a, b, &instructions)

	case 4, 5:
		trierCinq(a, b, &instructions)

	default:
		trierParBlocs(a, b, &instructions)
	}

	return instructions
}

// trierTrois trie une pile de trois nombres en deux opérations maximum.
func trierTrois(a, b *stack.Stack, instructions *[]string) {
	x := a.Numbers[0]
	y := a.Numbers[1]
	z := a.Numbers[2]

	if x > y && y < z && x < z {
		operations.Apply("sa", a, b)
		*instructions = append(*instructions, "sa")

	} else if x > y && y > z {
		operations.Apply("sa", a, b)
		*instructions = append(*instructions, "sa")

		operations.Apply("rra", a, b)
		*instructions = append(*instructions, "rra")

	} else if x > y && y < z && x > z {
		operations.Apply("ra", a, b)
		*instructions = append(*instructions, "ra")

	} else if x < y && y > z && x < z {
		operations.Apply("sa", a, b)
		*instructions = append(*instructions, "sa")

		operations.Apply("ra", a, b)
		*instructions = append(*instructions, "ra")

	} else if x < y && y > z && x > z {
		operations.Apply("rra", a, b)
		*instructions = append(*instructions, "rra")
	}
}

// positionDuMinimumA trouve la position du plus petit rang dans A.
func positionDuMinimumA(a *stack.Stack) int {
	position := 0

	for i := 1; i < len(a.Numbers); i++ {
		if a.Numbers[i] < a.Numbers[position] {
			position = i
		}
	}

	return position
}

// amenerAuSommetA utilise le chemin le plus court.
func amenerAuSommetA(a, b *stack.Stack, position int, instructions *[]string) {
	size := len(a.Numbers)

	if position <= size/2 {
		for i := 0; i < position; i++ {
			operations.Apply("ra", a, b)
			*instructions = append(*instructions, "ra")
		}
	} else {
		for i := 0; i < size-position; i++ {
			operations.Apply("rra", a, b)
			*instructions = append(*instructions, "rra")
		}
	}
}

// trierCinq met de côté les plus petits nombres dans B,
// trie les trois nombres restants, puis les remet dans A.
func trierCinq(a, b *stack.Stack, instructions *[]string) {
	for len(a.Numbers) > 3 {
		position := positionDuMinimumA(a)

		amenerAuSommetA(a, b, position, instructions)

		operations.Apply("pb", a, b)
		*instructions = append(*instructions, "pb")
	}

	trierTrois(a, b, instructions)

	for !b.IsEmpty() {
		operations.Apply("pa", a, b)
		*instructions = append(*instructions, "pa")
	}
}

// pousserParBlocs envoie les rangs de A vers B.
func pousserParBlocs(a, b *stack.Stack, tailleBloc int, instructions *[]string) {
	pousses := 0

	for len(a.Numbers) > 0 {
		rank := a.Numbers[0]

		if rank <= pousses {
			operations.Apply("pb", a, b)
			*instructions = append(*instructions, "pb")
			pousses++

			if len(b.Numbers) > 1 {
				operations.Apply("rb", a, b)
				*instructions = append(*instructions, "rb")
			}

		} else if rank <= pousses+tailleBloc {
			operations.Apply("pb", a, b)
			*instructions = append(*instructions, "pb")
			pousses++

		} else {
			operations.Apply("ra", a, b)
			*instructions = append(*instructions, "ra")
		}
	}
}

// positionDuMaximumB trouve la position du plus grand rang dans B.
func positionDuMaximumB(b *stack.Stack) int {
	position := 0

	for i := 1; i < len(b.Numbers); i++ {
		if b.Numbers[i] > b.Numbers[position] {
			position = i
		}
	}

	return position
}

// amenerAuSommetB utilise rb ou rrb selon le chemin le plus court.
func amenerAuSommetB(a, b *stack.Stack, position int, instructions *[]string) {
	size := len(b.Numbers)

	if position <= size/2 {
		for i := 0; i < position; i++ {
			operations.Apply("rb", a, b)
			*instructions = append(*instructions, "rb")
		}
	} else {
		for i := 0; i < size-position; i++ {
			operations.Apply("rrb", a, b)
			*instructions = append(*instructions, "rrb")
		}
	}
}

// remonterLesPlusGrands remet les nombres de B vers A,
// du plus grand rang au plus petit.
func remonterLesPlusGrands(a, b *stack.Stack, instructions *[]string) {
	for !b.IsEmpty() {
		position := positionDuMaximumB(b)

		amenerAuSommetB(a, b, position, instructions)

		operations.Apply("pa", a, b)
		*instructions = append(*instructions, "pa")
	}
}

// tailleDeBloc choisit la taille des blocs.
func tailleDeBloc(n int) int {
	if n <= 100 {
		return 15
	}

	return 30
}

// trierParBlocs assemble les deux phases de la séance 5.
func trierParBlocs(a, b *stack.Stack, instructions *[]string) {
	taille := tailleDeBloc(len(a.Numbers))

	pousserParBlocs(a, b, taille, instructions)
	remonterLesPlusGrands(a, b, instructions)
}
