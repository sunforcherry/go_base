package main 
import (
	"fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
    s := a[2:8] //切下标2-8但不包括8的slice
  	fmt.Println(s)	//输出：[3 4 5 6 7 8]

  	slice_make := make([]int,10,20) //创建初始元素10个，值为0，capacity为20的slice
  	fmt.Println(slice_make)

  	sa := a[2:7]
  	sa = append(sa,100)
  	sb := sa[3:5]
  	sb[0] = 99
  	// fmt.Println(&sb[0])
  	sb = append(sb,20,15) //切记在不超过capcity时，slice只是获取了源数组某元素的一个地址而已，所以往里append数据时，实际是修改了源数组的某个值
  	// sb = append(sb,20,15,11,23) //capcity不够，自己扩容，内存地址变了，新的slice复制了原数组的内容，但地址变了
  	// fmt.Println(&sb[0])
  	fmt.Println(a,len(a),cap(a))
  	fmt.Println(sa,len(sa),cap(sa))
  	fmt.Println(sb,len(sb),cap(sb))

  	fmt.Println(&a[0])
  	fmt.Println(&a[2],&sa[0])
  	fmt.Println(&a[5],&sb[0])

  	b := []int{1,2,3,4}
  	slice_b := b[1:3]
  	fmt.Println(slice_b,len(slice_b),cap(slice_b))
  	slice_b = append(slice_b,11,22,33)
  	fmt.Println(slice_b,len(slice_b),cap(slice_b))
}