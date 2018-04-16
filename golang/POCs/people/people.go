package main
import "fmt"

var salutation string = "Hello"
const desig string = "Mr"
func main(){
	name := "Max Payne"
	fmt.Println(salutation +" "+desig+" "+ name)
	i := 1
	for i <= 10 {
		fmt.Println("Soldier ",i," says FREEZE!!!")	
		i++
	}

	fmt.Println("Max selects weapons")
	selector := 2
	switch selector {
		case 1: fmt.Println("Beretta")
		case 2: fmt.Println("Desert Eagle")
		case 3: fmt.Println("Pump Action Shotgun")
		case 4: fmt.Println("M4")
		case 5: fmt.Println("Grenade Launcher")
		default: fmt.Println("Baseball Bat")
	}
}
