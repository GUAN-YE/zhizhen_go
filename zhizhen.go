// package main


// import ."fmt"

// func main(){
// 	var a int = 4
// 	var b *int
// 	b = &a
// 	Println(*b)
// }




// if(ptr != nil)     /* ptr 不是空指针 */
// if(ptr == nil)    /* ptr 是空指针 */

// for (指针)

package main
import "fmt"

func main(){
	const max int = 3
	aa := []int{100,200,3333}
	var i int
	var bb [max]*int
	var cc []int
	for i=0;i<max;i++ {
		bb[i]=&aa[i]
		cc=append(cc,*bb[i])

	}
	fmt.Println(cc)
	
}
// func data(max int)int{
// 	aa := []int{100,200,3333}
// 	var i int
// 	var bb [max]*int
// 	for i=0;i<max;i++ {
// 		bb[i]=&aa[i]
// 	}
// 	bb = *bb
// 	return bb
// }