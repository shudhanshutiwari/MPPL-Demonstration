package main

import (
	"fmt"
)

func main(){
	var num int = 10;
	var num2 float64 = 15.5;
	var name = "Shudhanshu";
	var boolean bool = true;
	
	fmt.Println("Value of Num",num," Type of Num",fmt.Sprintf("%T",num));
	fmt.Println("Value of Num2",num2," Type of Num2",fmt.Sprintf("%T",num2));
	fmt.Println("Value of Name",name," Type of Name",fmt.Sprintf("%T",name));
	fmt.Println("Value of Boolean",boolean," Type of Boolean",fmt.Sprintf("%T",boolean));

}
