// will panic when rand.Intn(0)

package main

import (
	"fmt"
	"math/rand"	
)

func main(){
	fmt.Println(rand.Intn(0))
	// fmt.Println(rand.Intn(1))
}


