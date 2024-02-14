package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var score = 99.5

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("good bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// 뒤에 float64는 return type
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string){
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], " "

}

func updateName(x *string) {
	// pointer를 받겠다는 의미
	*x = "wedge"
}

func updateMenu(y map[string]float64) {
	y["pie"] = 5.99
	y["coffee"] = 2.99
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Create a new bill name: ")
	//name, _ := reader.ReadString('\n')
	//name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill =", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip: ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name:", reader)
		price, _ := getInput("item price:", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println(tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you save teh file", b)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is : %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}


type salesRecord interface {
	getMargin() float64
	getItem() string
}

type imported struct {
	name string // 제품명
	customerPrice float64 // 소비자 판매가
	venderPrice float64 // 매입가
	duty float64 // 관세
	forwarder float64 // 운송비
}

type exported struct {
	name string // 제품명
	customerPrice float64 // 소비자 판매가
	sourcePrice float64 // 원자재가
	forwarder float64 // 운송비
}

func (i imported) getMargin() float64 {
	return i.customerPrice - i.venderPrice - i.duty - i.forwarder
}

func (e exported) getMargin() float64 {
	return e.customerPrice - e.sourcePrice - e.forwarder
}

func (i imported) getItem() string {
	return i.name
}

func (e exported) getItem() string {
	return e.name
}

func printSalesInfo(s salesRecord) {
	fmt.Println("item name: ", s.getItem())
	fmt.Println("margin: ", s.getMargin())

}


func main() {
	/*
	*******************# 2~5 Strings, Numbers, Formatting & Array*********************************
	 */
	// strings
	//var nameOne string = "mario"
	//var nameTwo = "luigi"
	//var nameThree string
	//
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//nameOne = "peach"
	//nameThree = "bowser"
	//
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//nameFour := "yoshi"
	//fmt.Println(nameFour)

	//// ints
	//var ageOne int = 20
	//var ageTwo = 30
	//ageThree := 40
	//fmt.Println(ageOne, ageTwo, ageThree)
	//
	//// bits & memory
	////var numOne int8 = 25
	////var numTwo int8 = -128
	////var numThree uint16 = 256
	//
	//var scoreOne float32 = 25.98
	//var scoreTwo float64 = 546548473245465464.7

	//age := 35
	//name := "shaun"
	//// print
	//fmt.Print("hello, ")
	//fmt.Print("world \n")
	//fmt.Print("new line \n")
	//
	//// println
	//fmt.Println("hello ninjas!")
	//fmt.Println("good bye ninjas!")
	//fmt.Println("my age is", age, "and my name is", name)
	//
	//// Printf (formatted strings)
	//fmt.Printf("my age is %v and my name is %v \n", age, name)
	//fmt.Printf("my age is %q and my name is %q \n", age, name)
	//
	//fmt.Printf("age is of type %T \n", age)
	//fmt.Printf("you scored %f points \n", 225.55)
	//
	//// Sprintf (save formatted strings)
	//var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	//fmt.Println("the saved string is:", str)

	//var ages [3]int = [3]int{20, 25, 30}
	//var ages = [3]int{20, 25, 30}
	//names := [4]string{"yoshi", "mario", "peach", "bowser"}
	//names[1] = "luigi"
	//
	//fmt.Println(names, len(names))
	//
	//// slices (use arrays under the hood)
	//var scores = []int{100, 50, 60}
	//
	//scores[2] = 25
	//scores = append(scores, 85)
	//
	//fmt.Println(scores, len(scores))
	//
	//// slice ranges
	//rangeOne := names[1:3]
	//rangeTwo := names[2:]
	//rangeThree := names[:3]
	//
	//fmt.Println(rangeOne, rangeTwo, rangeThree)
	//
	//rangeOne = append(rangeOne, "koopa")
	//fmt.Println(rangeOne)

	/*
	******************* #6 strings, Ints, sort*********************************
	 */

	//greeting := "hello there friends!"
	//fmt.Println(strings.Contains(greeting, "hello"))
	//fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	//fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "th"))
	//fmt.Println(strings.Split(greeting, " "))

	// 이렇게한다고 원본이 바뀌진 않음
	//fmt.Println("original:", greeting)

	//ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	//sort.Ints(ages)
	//fmt.Println(ages)
	//
	//index := sort.SearchInts(ages, 30)
	//fmt.Println(index)
	//
	//names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	//
	//sort.Strings(names)
	//fmt.Println(names)
	//
	//fmt.Println(sort.SearchStrings(names, "bowser"))

	/*
	******************* #7 Loops *********************************
	*/
	//x := 0
	//for x < 5 {
	//	fmt.Println("value of x is :", x)
	//	x++
	//}

	//for i := 0; i < 5; i++ {
	//	fmt.Println("value of i is :", i)
	//}

	//names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}

	//for index, value := range names {
	//	fmt.Printf("the value at index %v is %v \n", index, value)
	//}
	//
	//for _, value := range names {
	//	fmt.Printf("the value is %v \n", value)
	//}
	//
	//fmt.Println(names)

	/*
	******************* #8 Booleans & conditionals *********************************
	*/

	//age := 45
	//
	//fmt.Println(age <= 50)
	//fmt.Println(age >= 50)
	//fmt.Println(age == 45)
	//fmt.Println(age != 50)
	//
	//if age < 30 {
	//	fmt.Println("age is less than 30")
	//} else if age < 40 {
	//	fmt.Println("age is less than 40")
	//} else {
	//	fmt.Println("age is not less than 45")
	//}
	//
	//names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}
	//
	//for index, value := range names {
	//	if index == 1 {
	//		fmt.Println("continuing at pos", index)
	//		continue
	//	//	다음 index로 for loop을 진행시키는 것
	//	}
	//	if index > 2 {
	//		fmt.Println("breaking at pos", index)
	//		break
	//	//	 아예 for loop을 나가는 것
	//	}
	//	fmt.Printf("the value at pos %v is %v \n", index, value)
	//}

	/*
	******************* #9 Using Functions *********************************
	*/

	//sayGreeting("mario")
	//sayGreeting("luigi")
	//sayBye("mario")
	//
	//cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	//
	//a1 := circleArea(10.5)
	//a2 := circleArea(15)
	//fmt.Println(a1, a2)
	//fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1,a2)

	/*
	******************* #10 Multiple Return Values *********************************
	*/

	//fn1, sn1 := getInitials("tifa lockhart")
	//fmt.Println(fn1, sn1)
	//fn2, sn2 := getInitials("cloud strife")
	//fmt.Println(fn2, sn2)
	//fn3, sn3 := getInitials("barret")
	//fmt.Println(fn3, sn3)

	/*
	******************* #11 Package Scope *********************************
	*/
	//sayHello("mario")
	//
	//for _, v := range points {
	//	fmt.Println(v)
	//}
	//
	//showScore()

	/*
	******************* #12 Maps *********************************
	*/
	//menu := map[string]float64{
	//	"soup":4.99,
	//	"pie":7.99,
	//	"salad":6.99,
	//	"toffee pudding": 3.55,
	//}
	//fmt.Println(menu)
	//fmt.Println(menu["pie"])
	//
	//// looping maps
	//for k, v := range menu {
	//	fmt.Println(k, "-", v)
	//}
	//
	//// ints as key type
	//phonebook := map[int]string{
	//	267584967: "mario",
	//	984759373: "luigi",
	//	845775485: "peach",
	//}
	//
	//fmt.Println(phonebook)
	//fmt.Println(phonebook[984759373])
	//
	//phonebook[984759373] = "bowser"
	//fmt.Println(phonebook)

	/*
	******************* #13-14 Pass By Value, pointers *********************************
	*/

	// group A types -> strings, ints, bools, floats, arrays, structs
	//name := "tifa"
	//
	////updateName(name)
	////fmt.Println("memory address of name is: ", &name)
	//
	//m := &name
	////fmt.Println("memory address: ", m)
	////fmt.Println("value at memory address:", *m)
	//fmt.Println(name)
	//updateName(m)
	//fmt.Println(name)
	//
	//updateName(m)
	//
	//// group B types -> slices, maps, functions
	//menu := map[string]float64{
	//	"pie": 5.95,
	//	"ice cream": 3.99,
	//}
	//
	//updateMenu(menu)
	//fmt.Println(menu)

	/*
	******************* #15 structs and custom types *********************************
	*/
	//mybill := newBill("mario's bill")
	//fmt.Println(mybill)

	/*
	******************* #16-17 Receiver Functions & with pointers*********************************
	 */
	//mybill := newBill("mario's bill")
	//
	//mybill.addItem("onion soup", 4.50)
	//mybill.addItem("veg pie", 8.95)
	//mybill.addItem("toffee pudding", 4.95)
	//mybill.addItem("coffee", 3.25)
	//mybill.updateTip(10)
	////format()
	//fmt.Println(mybill.format())

	/*
	******************* #18, 19, 20,21 user input & switch statements & parsing floats *********************************
	 */
	//mybill := createBill()
	//promptOptions(mybill)

	/*
	******************* #22 Interfaces *********************************
	 */
	//shapes := []shape{
	//	square{length: 15.2},
	//	circle{radius: 7.5},
	//	circle{radius: 12.3},
	//	square{length: 4.9},
	//}
	//
	//for _, v := range shapes {
	//	printShapeInfo(v)
	//	fmt.Println("---")
	//}

	/*
	******************* #22 번외 *********************************
	 */
	mySales := []salesRecord{
		imported{name: "apple", customerPrice: 10.5, venderPrice: 3, duty: 0.08, forwarder: 2.5},
		exported{name: "grape", customerPrice: 8.99, sourcePrice: 2, forwarder: 3},
	}

	for _, v := range mySales {
		printSalesInfo(v)
		fmt.Println("---")
	}

}

