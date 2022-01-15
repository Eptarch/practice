package lasagna

func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2
    }
    return len(layers) * avgPrepTime
}

func Quantities(layers []string) (int, float64) {
    noodles, sauce := 0, 0.0 
    for _, layer := range layers {
        switch layer {
            case "noodles": noodles += 50
            case "sauce": sauce += 0.2
        }
    }
    return noodles, sauce
    
}

func AddSecretIngredient(friendsList, myList []string) []string {
    return append(myList[:len(myList) - 1], friendsList[len(friendsList) - 1])

}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    result := make([]float64, len(quantities))
    copy(result, quantities)
    for i, quantity := range quantities {
        result[i] = quantity / 2 * float64(portions)
    }
    return result
}
