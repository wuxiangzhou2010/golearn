## Go 面试基础

1. 请问下面代码输出:

   ```go
   package main

   import "fmt"

   func main() {
   	var a int = 10
   	fmt.Println(a)
   	{
   		a := 9
   		fmt.Println(a)
   		a = 8
   	}
   	fmt.Println(a)

   }
   ```

   Answer: 10, 9, 10

   2. 请问下面代码输出

   ```go
   package main

   import "fmt"

   func main() {
   	if true {
   		defer fmt.Println("1")
   	} else {
   		defer fmt.Println("2")
   	}
   	fmt.Println("3")
   }

   ```

   Answer: 3 1

2. 请问下面代码的输出

   ```go
   package main

   import "fmt"

   func main() {

       fmt.Println((1+6)/2*4 ^ 2 + 10%3<<3)
   }

   ```

   Answer: 22

3. 请问哪个声明不正确

   ```go
   A: func test(a,b int)(int, error)
   B: func test(a int, b int)(int, error)
   C: func test(a,b int)(val int, error)
   D: func test(a int, b int)(val int, err error)
   ```

   Answer: C ---> syntax error: mixed named and unnamed function parameters

4. 下面代码输出是怎样的：

   ```go
   package main

   import "fmt"

   type student struct {
   	Name string
   	Age  int
   }

   func main() {

   	m := make(map[string]*student)

   	stus := []student{
   		{Name: "zhou", Age: 22},
   		{Name: "Li", Age: 23},
   		{Name: "Wang", Age: 24},
   	}

   	for _, stu := range stus {
   		m[stu.Name] = &stu
   		// fmt.Println(stu)
   	}

   	for k, v := range m {
   		fmt.Println(k, v)
   	}
   }

   ```

   Answer:

   ```sh
   zhou &{Wang 24}
   Li &{Wang 24}
   Wang &{Wang 24}
   ```

5. 下面代码输出是：

   ```go
   package main

   import "fmt"

   func swap(a, b *int) (*int, *int) {
   	a, b = b, a
   	return a, b
   }

   func main() {
   	a, b := 3, 4
   	c, d := swap(&a, &b)
   	fmt.Println(a, b)
   	a = *c
   	b = *d
   	fmt.Println(a, b)
   }
   ```

   Answer:

   ```sh
   3, 4
   4, 4
   ```

6. 关于 init 函数， 下面说法正确的是：

   ```txt
   A. 一个包中， 可以包含多个init函数
   B. 程序编译的时候， 先执行导入包的init函数， 再执行本包内的init函数
   C. main包中不能包含init函数
   D. init函数可以被其他函数调用
   ```

   Answer： A B

7. 下面函数的执行结果是

   ```go
   package main

   import "fmt"

   func main() {

   	defer fmt.Println(1)
   	defer recover()

   	defer fmt.Println(2)

   	panic("abc")

   }

   ```

   Answer: `recover` must be called inside a function.

   ```go
   2
   1
   panic: abc

   goroutine 1 [running]:
   ```

8. 下面函数的执行结果是

   ```go
   package main

   import "fmt"

   func main() {

   	defer fmt.Println(1)
   	defer func() {
   		recover()
   	}()

   	defer fmt.Println(2)

   	panic("abc")
   }
   ```

   Answer: recover can take effect in a defer function

   ```sh
   2
   1
```
   
9. 关于循环语句， 下面说法正确的是：

   ```txt
   A. 循环语句既支持for关键字， 也支持while 和 do-while
   B. 关键字for的基本用法和C/C++中没有什么差异
   B. for 关键字支持continue/break来控制循环，但是它提供了一个更高级的label,可以选择中断哪一个循环
   D. for循环不支持以逗号为间隔的多个赋值语句
   ```

   Answer: BD

10. 下面对函数 add 的调用正确的是：

    ```go
    func add(args ...int) int {
    	sum := 0

    	for _, arg := range args {
    		sum += arg
    	}
    	return sum
    }
    ```


    A. add(1,2)
    B. add(1,2,3)
    C. add([]int{1,2,3})
    D. add([]int{1,2,3}...)
    ```
    
    Answer: ABD

12. 关于 switch 语句 下面说法正确的是：

    ```txt
    A. 条件表达式必须为常量或者整数
    B. switch 中可以存在多个条件相同case
    C. 需要break来明确退出一个case
    D. 只有在case中明确添加fallthrough关键字， 才会继续执行紧跟的下一个case
    ```

    Answer: BD

13. Go 中的引用类型包括：

    ```txt
    A. slice
    B. map
    C. channel
    D. interface
    ```

    Answer: ABC

14. 下列赋值正确的是：

    ```go
    A.  var x = nil
    B. var x interface{} = nil
    C. var x string = nil
    D. var x error = nil
    ```

    Answer: BD

15. 下列 slice 初始化方式正确的是：

    ```go
    A. make([]int)
    B. make([]int, 0)
    C. make([]int, -1)
    D. make([]int,0, 10)
    ```

    Answer: D

16. 下列不能正确定义数组的赋值语句有：

    ```go
    A. var a = [5]int{1,2,3,4,5}
    B. var a = [...]int{1,2,3,4,5}
    C. var a [5]int = {1,2,3,4,5}
    D. var a = [5]int{1:2, 4:5}
    ```

    Answer: C

17. 对局部变量整形切片 x 赋值， 下面写法正确的是

    ```go
    A. x:= []int{
      					1,2,3,
                4,5,6,
    					}
    B. x:= []int{
      			1,2,3,
      			4,5,6  //必须有逗号
    				}
    C. x := []int{
      			1,2,3,
    				4,5,6}
    D. x := []int{1,2,3,4,5,6}
    ```

    Answer: ACD

18. 下面程序的输出是

    ```go
    package main

    import "fmt"

    type Slice []int

    func NewSlice() Slice {
    	return make(Slice, 0)
    }

    func (s Slice) Add(elem int) *Slice {

    	s = append(s, elem)
    	fmt.Println(elem)
    	return &s
    }

    func main() {
    	s := NewSlice()
    	defer s.Add(1).Add(2).Add(3)
    	s.Add(4)

    }
    ```

    Answer: 1243

19. 下面代码会输出什么？

    ```go
    package main

    import "fmt"

    type people struct{}

    func (p *people) showA() {
    	fmt.Println("showA ")
    	p.showB()
    }
    func (p *people) showB() {
    	fmt.Println("showB")
    }

    type teacher struct {
    	people
    }

    func (t *teacher) showB() {
    	fmt.Println("Teacher showB")
    }
    func main() {
    	t := teacher{}
    	t.showA()
    	t.showB()
    }
    ```

    Answer:

    ```sh
    showA
    showB
    Teacher showB
    ```
