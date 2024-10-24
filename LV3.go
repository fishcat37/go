package main
import (
	"fmt"
	"math"
)
func judge (a int) {
	for i:=1;i<=int(math.Sqrt(float64(a)));i++ {
		if a%i==0 {
			fmt.Printf("No")
			return 
		}
	}
	fmt.Printf("Yes")

}

func main () {
	judge(7498134739)
}