package lasagna

func PreparationTime(layers []string, averagePreparationTimePerLayer int) int {
	if averagePreparationTimePerLayer == 0 {
		averagePreparationTimePerLayer = 2
	}
	return len(layers) * averagePreparationTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friend, own []string) {
	own[len(own)-1] = friend[len(friend)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities = []float64{}
	for _, q := range quantities {
		scaledQuantities = append(scaledQuantities, q*float64(portions)/2)
	}
	return scaledQuantities
}
