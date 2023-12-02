package main
import (
"fmt"
"os"
)
// This program computes the closest we can get to 1/2 at the int64 limit. go devs pls add 128bit arithmatic
func main() {
var num int64
var denum int64
const difference int = 1
var finalval int64 = 1
num = 1
denum = 2
//temp
fmt.Println(num, "/", denum)
		num = num * 2 + 1
		denum = denum * 2
		fmt.Println(num, "/", denum)
// sub 1 to denum and double both
	for {
	if denum != 4611686018427387904 {
		fmt.Println(num, "/", denum)
		num = num * 2 - 1
		denum = denum * 2
		fmt.Println(num, "/", denum)
	} else {
	fmt.Println(num * 2 - denum, "= difference")
	fmt.Println("***")
	fmt.Println("if this value doesnt equal 2 it is a bug")
	//check if finalval = 1 
	finalval = num * 2 - denum
	if finalval != 2 {
	os.Exit(1)
	}
	os.Exit(0)
	}
	}
}
