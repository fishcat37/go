package main
import "fmt"


func sum(a int,b int) int {
	return a+b
}
func main() {
	a:=100
	b:=99
	m:=sum(a,b)
	fmt.Println(m)
}