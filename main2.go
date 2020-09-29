//Calculo de e

package main

import "fmt"

func main(){
	var n float64
	var j float64
	var i float64
	var e float64
	var temp float64 //factorial

	e = 0	
	fmt.Scan(&n)
	for j = 0; j < n; j++ {
		temp = 1 //  se usara como el factorial
		for i = 1; i <= n - j; i++ {
			temp = temp * i
		} 
		e = e + 1.0 / temp
	}
	fmt.Println(e + 1)
}