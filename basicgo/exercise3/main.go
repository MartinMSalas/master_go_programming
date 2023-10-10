package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.
	var (
		a int
		b float64
		c bool
		d string
	)

	fmt.Println("a is", a, "b is", b, "c is", c, "d is", d)
	/*
		   	Using short declaration syntax declare and assign these values to variables x, y and z:

		   - 20

		   - 15.5

		   - "Gopher!"

		x, y, z := 20, 15.5, "Gopher!"
		//fmt.Println("x is", x, "y is", y, "z is", z)
		_, _, _ = x, y, z


			daysWeek with value 7

		2. lightSpeed with value 299792458

		3. pi with value 3.14159
	*/
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)
	fmt.Println("daysWeek is", daysWeek, "lightSpeed is", lightSpeed, "pi is", pi)

	const (
		secondsInHour = 3600
		secondsInDay  = 24 * secondsInHour
		secondsInYear = 365 * secondsInDay
	)
	fmt.Println("secondsInHour is", secondsInHour, "secondsInDay is", secondsInDay, "secondsInYear is", secondsInYear)
	const (
		Jun = iota + 6
		Jul
		Aug
	)
	fmt.Println("Jun is", Jun, "Jul is", Jul, "Aug is", Aug)
	x, y, z := 10, 15.5, "Gophers"
	fmt.Printf("x type is %T and x value is %d\n", x, x)
	fmt.Printf("y type is %T and y value is %f\n", y, y)
	fmt.Printf("z type is %T and z value is %s\n", z, z)

	score := []int{10, 20, 30}
	fmt.Printf("score type is %T and score value is %v\n", score, score)
	const x2 float64 = 1.422349587101
	fmt.Printf("x2 type is %T and x2 value is %v\n", x2, x2)
	fmt.Printf("x2 type is %T and x2 value is %.4f\n", x2, x2)

	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"
	convertedI := strconv.Itoa(i)
	fmt.Println(convertedI)
	str6 := strconv.FormatFloat(f, 'g', 4, 64)
  fmt.Println(str6)
	convertedS2, _ := strconv.ParseInt(s2, 10, 64)
	convertedS1, _ := strconv.ParseFloat(s1, 64)
	fmt.Printf("i modified type is %T and i value is %v\n", float64(i), float64(i))
	fmt.Println(convertedS2)
	fmt.Printf("f modified type is %T and f value is %v\n", str6, str6)
	fmt.Println(convertedS1)
/*
	ance from the Sun to Earth is 149.6 million km and the speed of light  is 299,792,458 m / s

The formula used is: time = distance / speed
*/
	const (
		distanceKM = 149600000
		disctanceM = distanceKM * 1000
		speedMS = 299792458
	)
	time := disctanceM / speedMS
	fmt.Println("The time in s to reach the Earth from the Sun is", time)

	type duration int
	var hour duration = 3600
	fmt.Printf("hour type is %T and hour value is %v\n", hour, hour)
	type (
				mile float64
				kilometer float64)

	const m2km = 1.60934 
	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer = kilometer(mileBerlinToParis * m2km)
	fmt.Printf("mileBerlinToParis type is %T and mileBerlinToParis value is %v\n", mileBerlinToParis, mileBerlinToParis)
	fmt.Printf("kmBerlinToParis type is %T and kmBerlinToParis value is %v\n", kmBerlinToParis, kmBerlinToParis)

}
