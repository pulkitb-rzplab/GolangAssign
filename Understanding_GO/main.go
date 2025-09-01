package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/exp/constraints"
)

func main() {
	for i := range 6 {
		star := strings.Repeat("*", i)
		fmt.Println(star)
	}

	//ARRAYS

	// h := "HARD"
	// m := "MEDIUM"
	// e := "EASY"
	// ch := 0
	// cm := 0
	// ce := 0

	// var arrQ = [25]string{"HARD", "EASY", "EASY", "MEDIUM", "HARD", "MEDIUM", "MEDIUM", "EASY", "EASY", "EASY", "HARD", "MEDIUM", "HARD", "EASY", "EASY", "MEDIUM", "HARD", "MEDIUM", "MEDIUM", "EASY", "EASY", "EASY", "HARD", "MEDIUM", "EASY"}
	// for i := range 25 {
	// 	switch arrQ[i] {
	// 	case h:
	// 		fmt.Println("The weightage for this Q is 3 points")
	// 		ch += 3
	// 	case m:
	// 		fmt.Println("The weightage for this Q is 2 points")
	// 		cm += 2
	// 	case e:
	// 		fmt.Println("The weightage for this Q is 1 point")
	// 		ce++

	// 	}
	// }

	// fmt.Println("The TOTAL marks to be obtained:", ce+cm+ch)

	// var TwoDArr [4][4]int
	// var i int = 0

	// for i < len(TwoDArr) {
	// 	var j int = 0
	// 	for j < len(TwoDArr) {
	// 		fmt.Print("Enter an element:")
	// 		fmt.Scan(&TwoDArr[i][j])
	// 		j++
	// 	}
	// 	i++

	// }
	// fmt.Println(TwoDArr)

	tDArr := [2][2]string{
		{"Apple", "Banana"},
		{"Fruit", "Vegetable"},
	}
	fmt.Println(tDArr)

	//SLICES

	var g []string
	fmt.Println("Slice: ", g, " len(slice): ", len(g), " cap(slice): ", cap(g))

	g = make([]string, 3)
	g[0] = "Happy"
	g[1] = "Birthday"
	g[2] = "To"

	fmt.Println(g)

	x := make([]string, 3)
	x = append(x, g...)
	fmt.Println(x)
	x[0] = "1"
	x[1] = "2"
	x[2] = "3... "
	fmt.Println(x)
	x[3] = "STOP"
	fmt.Println(x)

	//MAPS implementation

	m := make(map[string]int)

	m["Students"] = 84
	m["Teachers"] = 15
	m["Workers"] = 10
	m["Total Staff"] = m["Teachers"] + m["Workers"]

	fmt.Println(m)
	fmt.Println("Total Len of Map", len(m))
	fmt.Println("Map before deletion: ", m)
	delete(m, "Workers")
	fmt.Println("Map after deletion: ", m)

	//Functions

	ans := hcf(12, 58)
	fmt.Println("HCF of 12 and 58 is :", ans)

	a, b := 10, 81
	fmt.Printf("The initial values of 'a' and 'b' are: %d , %d \n ", a, b)
	a, b = swap(a, b)
	fmt.Println("The values of 'a' and 'b' after the Swap is: ", a, b)
	c, d := 15, 30
	_, d = swap(c, d)
	fmt.Println("Value of d after swapping:", d)

	//variadic func sum

	ans = sum(1, 2, 3, 4, 4, 5, 6)
	fmt.Println(ans)

	functon := arith()

	fmt.Println(functon())
	fmt.Println(functon())
	fmt.Println(functon())
	fmt.Println(functon())
	fmt.Println(functon())

	//recursion test
	fibonacci := fib(12 + 1)
	fmt.Println("The fibonacci sum for 12 is:", fibonacci)

	//Range over  built in type
	//Range over slice
	fmt.Println("\nRange over Slice:")

	sl := []int{23, 24, 25, 26, 27, 28, 29}
	fmt.Println("initial Slcie: ", sl)
	sl1 := []int{}
	sl2 := []int{}
	for _, v := range sl {
		if v%2 == 0 {
			sl1 = append(sl1, v)
		} else {
			sl2 = append(sl2, v)
		}
	}

	fmt.Println("slice with even no.: ", sl1)
	fmt.Println("Slice with odd nos. : ", sl2)

	//range over maps
	fmt.Println("\nRange over maps:")

	inv := map[string]int{
		"apple":       4,
		"Kiwi":        2,
		"banana":      9,
		"Grapes":      24,
		"dragonfruit": 1,
	}

	for k, v := range inv {
		if k == "Kiwi" || k == "dragonfruit" {
			fmt.Printf("This is a rare fruit %s with low stock of: %d \n", k, v)
		} else {
			fmt.Printf("This is a common fruit %s with stock of: %d \n", k, v)

		}
	}

	//Over strings
	fmt.Println("\nRange over strings:")
	for k, s := range "Happy Noise" {
		fmt.Println(k, s)
	}

	//pointers
	fmt.Println("\n Pointers: ")

	ab := 10
	bb := 18
	fmt.Printf("Initial values for both variables: ab: %d, bb: %d \n", ab, bb)
	noptr(ab)
	useptr(&bb)
	fmt.Println("Value of ab after noptr: ", ab)
	fmt.Println("Value of bb after useptr: ", bb)

	//Strings
	str1 := "Just"
	str2 := "Do"
	str3 := "It"

	fmt.Println("The Complete string together is : ", str1+str2+str3)

	//iterating over strings length
	for j := range len(str1) {
		fmt.Printf("%d : %v \n", j, string(str1[j]))
	}

	//iterating over string using range
	for indx, value := range str2 {
		fmt.Printf("%d : %U \n", indx, value)
	}

	//Structs
	fmt.Println("\nStructs: ")

	ferrari := newcar("V8", "Michellin", 0, "Carbon-Fibre", "Twin-Turbo", 762)
	fmt.Printf("A new car made using structs: \n %+v \n\n", ferrari)

	bmw := newcar("V10", "CEAT", 4, "Poslished Aluminum", "Naturally-Aspirated", 560)
	fmt.Printf("A new car made using structs: \n %+v \n\n", bmw)

	bmw.Power = 600
	fmt.Printf("An update in car  \n %+v \n\n", bmw)

	bmw.Calipers = "Carbon-Fibre"
	fmt.Printf("An update in car : \n %+v \n\n", bmw)

	//struct methods
	fmt.Println("\nStruct Methids:")
	sq := fig{look: "Square", length: 8, breadth: 8}
	rec := fig{look: "Rectangle", length: 6, breadth: 13}
	circle := fig{look: "Circle", radius: 4}

	fmt.Println("Perimeter of rectangle: ", peri(rec))
	fmt.Println("Area of rectangle: ", area(&rec))
	fmt.Println("Perimeter of circle: ", peri(circle))
	fmt.Println("Area of circle: ", area(&circle))
	fmt.Println("Square struct: ", sq)
	fmt.Println(math.Pi * 4 * 4)

	//Interfaces implemented in CalculateSalary.go

	//ENUMS

	a1 := transition(ServerStarting)
	a2 := transition(a1)
	a3 := transition(a2)
	fmt.Println("Server state: ", a1, a2, a3)

	//Struct Embedding

	cc := carContainer{
		engine: engine{
			Power:     850,
			Torque:    1350,
			Type:      "V8",
			Cylinders: 8,
			RPM:       12000,
		},
		doors:  2,
		color:  "Black",
		wheels: "Pirelli",
	}
	fmt.Println("Car Container: ", cc)

	fmt.Println("Car Engine Power: ", cc.Power)
	fmt.Println("Car Engine Description: ", cc.Describe())

	//Generic functions implementation
	intslice := []int{11, 2, 3, 4, 45, 5, 6, 67}
	fltslice := []float64{13.13, 44.2, 5.3, 66.4, 77.5, 88.6, 99.7, 100.8}
	minInt := findMin(intslice)
	minFloat := findMin(fltslice)

	fmt.Println("\n\nMinimum value in int slice: ", minInt)
	fmt.Println("\nMinimum value in float slice: ", minFloat)

	//Iterators

	//Errors
	v, e := f(5)
	if e != nil {
		fmt.Println("Error: ", e)
	} else {
		fmt.Println("Value: ", v)
	}
	//Sentinel errors
	conn, err := connect(0)
	fmt.Printf("\nTrying to connect to server: %v, %v\n", conn, err)
	conn, err = connect(2)
	fmt.Printf("\nTrying to connect to server: %v, %v\n", conn, err)

	//Custom errors using structs and satisfying error interface

	//Goroutines

	gorouteTest(5, "Hello")
	go gorouteTest(6, "World")
	go gorouteTest(6, "Go")
	time.Sleep(time.Second * 3)
	fmt.Println("function ends")

	//Playing with Channels

	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)
	go fromOne(ch1, "So", "Finesse")
	go Into2(ch1, ch2)
	fmt.Println("data", <-ch2, <-ch2)

	//select statement for goroutines and channels allow us to wait for multiple goroutines to finish first
	fmt.Println("\nSelect statement")
	ch3 := make(chan string, 2)
	ch4 := make(chan string, 2)
	go inpipe1(ch3, "Hello ", "Lebron")
	go inpipe2(ch4, "From ", "Satan")
	for range 2 {
		select {
		case msg1 := <-ch3:
			fmt.Printf("Message from ch3: %s", msg1)
			fmt.Printf(" %s ", <-ch3)
		case msg2 := <-ch4:
			fmt.Printf(" %s", msg2)
			fmt.Printf(" %s\n", <-ch4)
		}
	}

	//Timeouts can be used with channels and select easily. just add a <-time.After(time.Second * 1) case
	// and if the other goroutine takes more than that, we return timeout error
	//Eg. select{
	// 		case msg1 <- channel1:
	// 			fmt.Prinln("Go routine passed")
	//		case <-time.After(time.Second * 1):
	//			fmt.Println("Timeout")
	//	}

	// Non blocking channels

	//closing channels => close(chan)

	//Range Over Channels
	fmt.Println("\nRange Over Channels:")
	chh := make(chan int, 3)
	chh <- 3
	chh <- 2
	chh <- 1
	close(chh)

	for value := range chh {
		fmt.Println(value)
	}

	//Timers
	fmt.Println("\nTimers:")
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("Timer1 starts ticking")
	<-timer1.C
	fmt.Println("Timer1 executed")

	timer2 := time.NewTimer(time.Second * 3)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 executed")

	}()
	timer2.Stop()
	fmt.Println("Timer2 stopped before executing..")
	time.Sleep(time.Second * 3)

	//Tickers used for repeating an action in future at certain interval

	//Rate limiting
	fmt.Println("\nRate Limiting")
	req := make(chan int, 6)
	for i := range 6 {
		req <- i
	}
	close(req)
	limit := time.Tick(200 * time.Millisecond)

	for r := range req {
		<-limit
		fmt.Println("Request", r, "completed at time", time.Now())
	}

	//There is also bursty limiter which allows a burst of item to be processed at once.

	//Automic Counters allow for atomic operations on a single shared resource by many goroutines and without causing race conditions

	// sync/atomic package is used for atomic operations
	fmt.Println("\nAtomic Operations: ")
	var op atomic.Uint64
	var wg1 sync.WaitGroup
	for range 50 {
		wg1.Add(1)
		go func() {
			for range 1000 {
				op.Add(1)

			}
			wg1.Done()
		}()

	}
	wg1.Wait()
	fmt.Println("Total operations: ", op.Load())

}

func hcf(x int, y int) int {

	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func swap(a, b int) (int, int) {
	return b, a
}

//variadic functions

func sum(n ...int) int {
	fmt.Println("Total number of elements: ")
	fmt.Println(n)

	s := 0
	for i := range len(n) {
		s += n[i]

	}
	return s
}

// Closures
func arith() func() int {
	i := 1
	return func() int {
		i = i * (i + 1)
		return i
	}

}

//Recursion

func fib(i int) int {
	if i < 2 {
		return i
	}
	return fib(i-1) + fib(i-2)
}

func noptr(i int) {
	i = i + 20
}
func useptr(i *int) {

	*i += 20
}

type car struct {
	Engine   string
	Wheels   string
	Doors    int
	Calipers string
	Turbo    string
	Power    int
}

func newcar(engine string, wheels string, doors int, calipers string, turbo string, power int) *car {
	c := car{}
	if doors == 0 {
		c.Doors = 2

	} else {
		c.Doors = 4
	}
	c.Engine = engine
	c.Wheels = wheels
	c.Calipers = calipers
	c.Turbo = turbo
	c.Power = power

	return &c
}

func peri(r fig) float32 {
	if r.look == "Square" {
		return 4 * r.length
	} else if r.look == "Rectangle" {
		return 2 * (r.length + r.breadth)
	} else if r.look == "Circle" {
		return 2 * math.Pi * r.radius
	}
	return 0
}

func area(r *fig) float32 {
	switch r.look {
	case "Square", "Rectangle":
		return r.breadth * r.length
	case "Circle":
		return math.Pi * r.radius * r.radius
	}
	return 0
}

type fig struct {
	look    string
	length  float32
	breadth float32
	radius  float32
}

// Implementing ENUMS
type ServerState int

const (
	ServerStarting ServerState = iota
	ServerRunning
	ServerIdle
	ServerStopped
)

func (s ServerState) String() string {
	return []string{"ServerStarting", "ServerRunning", "ServerIdle", "ServerStopped"}[s]
}

func transition(ss ServerState) ServerState {
	switch ss {
	case ServerStarting:
		fmt.Println("Server is now Running: ", ServerRunning)
		return ServerRunning
	case ServerRunning:
		fmt.Println("Server is now Idle: ", ServerIdle)
		return ServerIdle
	case ServerIdle:
		fmt.Println("Server is now Stopped: ", ServerStopped)
		return ServerStopped
	}
	return ss
}

//Struct Embedding

type engine struct {
	Power     int
	Torque    int
	Type      string
	Cylinders int
	RPM       int
}

type carContainer struct {
	engine
	doors  int
	color  string
	wheels string
}

func (e engine) Describe() string {
	return fmt.Sprintf("This is a %s engine with %d cylinders, %d torque and %d power. It revvs up to %d RPM", e.Type, e.Cylinders, e.Torque, e.Power, e.RPM)

}

//Generic function to find minimum value in slice

func findMin[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		panic("Slice is empty")
	}

	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}

	}
	return min

}

//ERRORS

func f(i int) (int, error) {
	if i == 2 {
		return -1, errors.New("Invalid input :2")
	} else if i == 8 {
		return -1, fmt.Errorf("Invalid input : %d", i)
	} else {
		return i * i, nil
	}
}

// Sentinel errors
var noFreeNode = fmt.Errorf("No free nodes available")
var noFreeMemory = fmt.Errorf("No free memory available")
var networkError = fmt.Errorf("Network error")

func connect(input int) (int, error) {
	if input == 0 {
		return 100, nil
	} else if input == 1 {
		return -1, noFreeNode
	} else if input == 2 {
		return 101, noFreeMemory
	} else if input == 3 {
		return 104, networkError
	}
	return 0, nil
}

func gorouteTest(i int, s string) {
	for x := range i {
		fmt.Println("No: ", x, "String:", s)
	}
}

func fromOne(tube1 chan<- string, s1 string, s2 string) {
	fmt.Println("Entering from one... ")
	tube1 <- s1
	tube1 <- s2
}

// Directional channels, if tried flowing data in other way, will give compiletime error
func Into2(tube1 <-chan string, tube2 chan<- string) {
	fmt.Println("Going from one to two...")
	s1 := <-tube1
	s2 := <-tube1
	tube2 <- s1
	tube2 <- s2
}

func inpipe1(pip1 chan<- string, str1, str2 string) {
	pip1 <- str1
	pip1 <- str2
	time.Sleep(time.Second)
}

func inpipe2(pip1 chan<- string, str1, str2 string) {
	time.Sleep(time.Second * 2)
	pip1 <- str1
	pip1 <- str2

}
