package main

import (
	"fmt"
	"math/rand"
	"time"
)
func guess () int {
	rand.Seed( time.Now().UnixNano())
	num:=rand.Intn(100)
	fmt.Println(num)
	left:=0
	right:=100
	mid:=left+(right-left)/2
	for {
		mid=left+(right-left)/2
		if mid==num {
			return mid
		}else if mid<num {
			left=mid
		}else{
			right=mid
		}
	}

}
func main () {
	n:=guess()
	fmt.Println(n)
}