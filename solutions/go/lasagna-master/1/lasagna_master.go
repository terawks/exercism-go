package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgT int) (resT int) {
	if avgT == 0 {
		avgT = 2
	}
	return len(layers) * avgT
}

// TODO: define the 'Quantities()' function
func Quantities(q []string) (noodles int, sauce float64) {
	for _, i := range q {
		if i == "noodles" {
			noodles += 50
		} else if i == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberPortions int) (res []float64) {
	factor := float64(numberPortions) / 2.0
	for idx, _ := range quantities {
		res = append(res, quantities[idx]*factor)
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
