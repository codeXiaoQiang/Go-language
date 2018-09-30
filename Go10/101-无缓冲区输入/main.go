package main
/*
#include <stdio.h>
char lowerCase(char ch){
    // 1.判断当前是否是小写字母
    if(ch >= 'a' && ch <= 'z'){
        return ch;
    }
    // 注意点: 不能直接编写else, 因为执行到else不一定是一个大写字母
    else if(ch >= 'A' && ch <= 'Z'){
        return ch + ('a' - 'A');
    }
    return ' ';
}
char getCh(){
    // 1.接收用户输入的数据
    char ch;
    scanf("%c", &ch);
    setbuf(stdin, NULL);
    // 2.大小写转换
    ch = lowerCase(ch);
    // 3.返回转换好的字符
    return ch;
}
 */
import "C"
import "fmt"

func main() {
	for {
		fmt.Println("请输入一个字符")
		ch := C.getCh()
		fmt.Printf("%c\n", ch)
	}
}
