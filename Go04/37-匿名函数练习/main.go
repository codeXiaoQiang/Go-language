package main

import (
	"fmt"
)

func main() {

	//work(func() {
	//	fmt.Println("参与部门会议")
	//	fmt.Println("修改BUG")
	//	fmt.Println("完成老大安排的任务")
	//	fmt.Println("... ...")
	//	fmt.Println("提交代码")
	//})

	work(func() {
		fmt.Println("组织部门会议")
		fmt.Println("给部门的员工分配任务")
		fmt.Println("检查部门员工昨天提交的代码")
		fmt.Println("... ...")
	})
}
// 攻城狮的一天
func work(method func()){
	// 1.上班前
	func(){
		fmt.Println("起床")
		fmt.Println("刷牙洗脸")
		fmt.Println("吃早餐")
		fmt.Println("做地铁")
		fmt.Println("打卡")
		fmt.Println("打开电脑")
	}()

	// 2.上班中
	method()

	// 3.下班之后
	func(){
		fmt.Println("关电脑")
		fmt.Println("打卡")
		fmt.Println("坐地铁")
		fmt.Println("到家")
		fmt.Println("吃饭")
		fmt.Println("娱乐")
		fmt.Println("刷牙洗脸")
		fmt.Println("睡觉")
	}()

}

// 项目经理的一天
func work2(){
	fmt.Println("起床")
	fmt.Println("刷牙洗脸")
	fmt.Println("吃早餐")
	fmt.Println("做地铁")
	fmt.Println("打卡")
	fmt.Println("打开电脑")

	fmt.Println("组织部门会议")
	fmt.Println("给部门的员工分配任务")
	fmt.Println("检查部门员工昨天提交的代码")
	fmt.Println("... ...")

	fmt.Println("关电脑")
	fmt.Println("打卡")
	fmt.Println("坐地铁")
	fmt.Println("到家")
	fmt.Println("吃饭")
	fmt.Println("娱乐")
	fmt.Println("刷牙洗脸")
	fmt.Println("睡觉")
}
// 程序员的一天
func work1()  {
	fmt.Println("起床")
	fmt.Println("刷牙洗脸")
	fmt.Println("吃早餐")
	fmt.Println("做地铁")
	fmt.Println("打卡")
	fmt.Println("打开电脑")

	fmt.Println("参与部门会议")
	fmt.Println("修改BUG")
	fmt.Println("完成老大安排的任务")
	fmt.Println("... ...")
	fmt.Println("提交代码")

	fmt.Println("关电脑")
	fmt.Println("打卡")
	fmt.Println("坐地铁")
	fmt.Println("到家")
	fmt.Println("吃饭")
	fmt.Println("娱乐")
	fmt.Println("刷牙洗脸")
	fmt.Println("睡觉")
}
