package main

import "fmt"

func main() {

	grades := [3]float64{
		11,
		22,
	}
	fmt.Printf("%#v\n", grades)
	taskDone := [...]bool{
		5: true,
		false,
		true,
		11: true,
	}
	fmt.Printf("%#v\n", taskDone)
	cities := [5]string{
		"Paris",
		"London",
		"Tokyo",
		"Catamarca",
		"Cordoba",
	}
	fmt.Printf("%#v\n", cities)
	for i := 0; i < len(cities); i++ {
		fmt.Printf("cities[%d] is %s\n", i, cities[i])
	}
	for i, v := range cities {
		fmt.Printf("cities[%d] is %s\n", i, v)
	}

	for i,v := range grades {
		fmt.Printf("grades[%d] is %f\n", i, v)
	}
	nums := [...]int{30, -1, -6, 90, -6}

	acum := 0
	for i := 0; i < len(nums); i++ {
		if(nums[i]%2 == 0 && nums[i] > 0){
			acum++			
		}
	}
	fmt.Printf("acum is %d\n", acum)

}
