package main
import (
	"fmt"
	"math"
)
func hwan (r float32) float32 {
	area:=math.Pi*r*r
	return area
}
func main () {
	var r float32
	fmt.Scan(&r)
	area:=hwan(r)
	fmt.Println(area)
}