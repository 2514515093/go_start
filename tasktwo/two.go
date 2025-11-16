package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//one()
	//two()
	//three()
	//four()
	//five()
	//six()
	//seven()
	//eight()
	//nine()
	ten()
}

func one() {
	//题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
	//在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	num := 20
	oneZs(&num)
	fmt.Println(num)
}

func oneZs(num *int) {
	*num += 10
}

func two() {
	//实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	//考察点 ：指针运算、切片操作。
	nums := []int{1, 2, 3}
	twoSlice(nums)
	fmt.Println(nums)
}

func twoSlice(nums []int) {
	for i, v := range nums {
		nums[i] = v * 2
	}
}

func three() {
	//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println("协程一奇数打印===", i)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println("协程二偶数打印===", i)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
}

func four() {
	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	task := func() []int {
		return []int{2, 2, 3, 4, 8, 3, 5}
	}()
	fmt.Println(task)
	for _, v := range task {
		go func() {
			yxsj := fourtask(v)
			fmt.Println("=========", v, "运行时间", yxsj)
		}()
	}
	time.Sleep(5 * time.Second)
}

func fourtask(num int) int {
	start := time.Now()
	time.Sleep(time.Duration(num) * 100 * time.Millisecond)
	end := time.Now()
	return int(end.Sub(start).Milliseconds())
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

func (r *Rectangle) Area() {
	fmt.Println("我是Rectangle,我再执行Area方法")
}
func (r *Circle) Area() {
	fmt.Println("我是Circle,我再执行Area方法")
}

func (r *Rectangle) Perimeter() {
	fmt.Println("我是Rectangle,我再执行Perimeter方法")
}

func (r *Circle) Perimeter() {
	fmt.Println("我是Circle,我再执行Perimeter方法")
}

func five() {
	//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
	//然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	//在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	rectangle := &Rectangle{}
	circle := &Circle{}

	rectangle.Area()
	rectangle.Perimeter()
	circle.Area()
	circle.Perimeter()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (emp *Employee) PrintInfo() {
	fmt.Println("employeeId:", emp.EmployeeId)
	fmt.Println("employeeName:", emp.Name)
	fmt.Println("employeeIAge:", emp.Age)
}

func six() {
	//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
	//再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
	//为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	//考察点 ：组合的使用、方法接收者
	emp := &Employee{Person{"Jack", 18}, 1}
	emp.PrintInfo()
}

func seven() {
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
	//并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	cn := make(chan int)
	timecs := time.After(5 * time.Second)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("发送===", i)
			cn <- i
		}
	}()
	go func() {
		for {
			select {
			case i := <-cn:
				fmt.Println("接受====", i)
			case <-timecs:
				fmt.Println("超时关闭")
				return
			default:
				//fmt.Println("等待中")
			}
		}
	}()
	time.Sleep(8 * time.Second)
}

func eight() {
	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	cn := make(chan int, 50)
	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("发送===", i)
			cn <- i
			//time.Sleep(100 * time.Millisecond)
		}
	}()
	timecs := time.After(5 * time.Second)
	for {
		select {
		case i := <-cn:
			fmt.Println("接受====", i)
		case <-timecs:
			fmt.Println("超时关闭")
			return
		default:
			//fmt.Println("等待中")
		}
	}
	time.Sleep(8 * time.Second)

}

func nine() {
	//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	count := 0
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				countjia(&count)
			}
		}()
	}
	//fmt.Println("最终计数值为===", count)
	time.Sleep(3 * time.Second)
	fmt.Println("最终计数值为===", count)
}

var mt sync.Mutex

func countjia(num *int) {
	mt.Lock()
	defer mt.Unlock()
	*num++
}

func ten() {
	//：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	count := 0
	fs := int64(count)
	for i := 0; i < 10; i++ {
		go func(j int) {
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&fs, 1)
			}
			fmt.Println("任务", j, "=====完成")
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(fs)
}
