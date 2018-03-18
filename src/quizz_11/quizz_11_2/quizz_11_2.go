package main
     
import "fmt"
 // ignore the problem, quizz 11_1 and quizz 11_2 are the same
func main() {
 c := make(chan string)
 
 for i := 0; i < 4; i++ {
	 go printString("Hello there!", c)
 }
 
 for {
	 fmt.Println(<- c)
 }
}
 
func printString(s string, c chan string) {
 fmt.Println(s)
 c <- "Done printing." 
}