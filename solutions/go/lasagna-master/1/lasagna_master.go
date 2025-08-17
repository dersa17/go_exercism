package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
    if (avg == 0) {avg = 2}
    return len(layers) * avg
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var amountNoodles int
    var amountSauce float64 
    for i := 0; i < len(layers); i++ {
        if (layers[i] == "noodles") {
            amountNoodles += 50
        } else if (layers[i] == "sauce") {
            amountSauce += 0.2
        }
    }
    return amountNoodles, amountSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
    return myList
    
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
    scaled := make([]float64, len(amounts)) 
    for i, a := range amounts {
        scaled[i] = a * float64(portions) / 2.0
    }
    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
