package main

import (
	"fmt"
)

func main() {
	//位运算符
	//6 ：0110
	//11：1011
	//^:  1101 相同位取0不同位取1
	//&:  0010 两位相同为1则为1
	//&^: 0100 第二个数字位为1的则第一个数字位为0
	//|：  1111 只要有一位为1则为1
	fmt.Println(6^11, 6&11, 6&^11, 6|11)
	fmt.Println(6 & 4)
	fmt.Printf("%b", 15)
}