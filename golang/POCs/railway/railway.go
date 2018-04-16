package main
import "fmt"

func main() {
	var trains [5]int
	trains[4] = 100
	fmt.Println(trains)

	var trainWeights [5]float64
	trainWeights[0] = 112
	trainWeights[1] = 98
	trainWeights[2] = 100
	trainWeights[3] = 105
	trainWeights[4] = 89

  var totalTrainWeight float64 = 0 
	for idx := 0; idx < 5 ; idx ++ {
			totalTrainWeight += trainWeights[idx]
	}
	fmt.Println("Total Trains Weight",totalTrainWeight,"lbs")

	// Range example
	for _, value := range trainWeights {
		fmt.Println(value,"lbs")
	}

	stations := [5]string{"Delhi","Bangalore","Kolkata","Goa","Mumbai"}

	fmt.Println("Stations are ",stations)
}