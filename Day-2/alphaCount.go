// package main

// import (
// 	"fmt"
// 	"sync"
// 	"unicode"
// )

// func alphaFreq(s1 string, out chan<- string) {
// 	for _, v := range s1 {
// 		v = unicode.ToLower(v)
// 		if unicode.IsLetter(v) {
// 			out <- string(v)
// 		}
// 	}

// }

// func main() {
// 	words := []string{"Worthy", "Mark", "Leave", "Yam", "zoo", "Thursday"}
// 	letter := make(chan string, 1)
// 	wg := sync.WaitGroup{}
// 	for _, s := range words {
// 		wg.Go(func() {
// 			alphaFreq(s, letter)
// 		})
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(letter)
// 	}()

// 	dict := make(map[string]int)

// 	for v := range letter {
// 		dict[v]++
// 	}

// 	fmt.Println(dict)

// }

//OR Other method

// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func countfreq(s string, ch chan<- map[string]int) {

// 	dict := make(map[string]int)
// 	for _, v := range s {
// 		v = unicode.ToLower(v)
// 		if unicode.IsLetter(v) {
// 			dict[string(v)]++

// 		}
// 	}

// 	ch <- dict
// }

// func main() {
// 	words := []string{"Worthy", "Mark", "Leave", "Yam", "zoo", "Thursday"}
// 	ch := make(chan map[string]int, 1)
// 	for i := range len(words) {
// 		go countfreq(words[i], ch)
// 	}

// 	dict := make(map[string]int)

// 	for j := range len(words) {
// 		intermediate := <-ch
// 		fmt.Printf("\n%d :  \n", j)
// 		for k, v := range intermediate {
// 			dict[k] += v
// 			fmt.Println(dict[k])
// 		}
// 	}
// 	fmt.Println(dict)
// }
