package main
/*
#include <stdio.h>
int cArray[5] = {1, 2, 3, 4, 5};
*/
import "C"
import "fmt"

func main() {
	var arr [5]C.int = C.cArray
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
}
