package main

func main() {
	/*
		死锁：
			不是锁的一种！是一种错误使用锁导致的现象。
			1.单go程自己死锁
			2.go程间channel访问顺序导致死锁
			3.多go程,多channel交叉死锁
	*/
	//01.go至少在两个以上的go程中通信，否则死锁!!!

	//死锁1：类似队列，需要先出再进.----------channel访问顺序导致死锁
	{
		//ch := make(chan int)
		//ch <- 789
		//num := <-ch
		//fmt.Println("num=", num)
	}
	//死锁2：使用channel一端读（写），要保证另一段写（读）操作，同时有机会执行。否则死锁。
	{
		//ch := make(chan int)
		//num := <-ch //一旦阻塞了，以下代码都不会执行！！
		//fmt.Println("num=", num)
		//go func() {
		//	time.Sleep(2 * time.Second)
		//	ch <- 789
		//}()
	}
	//死锁3：多go程,多channel交叉死锁。-----逻辑错误
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() { //子
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()
	for {
		select {
		case num := <-ch2:
			ch1 <- num

		}
	}
}
