package main

//00最初效果：
//func main() {
////	go printer("hello")
////	go printer("world")
////	time.Sleep(3*time.Second)
////	fmt.Println("main...over...")
////}
////
////func printer(str string){
////	for _, v := range str {
////		time.Sleep(100*time.Millisecond)
////		fmt.Printf("%c",v)
////	}
////}}
//01.使用通道的方式
//var chan1 = make(chan bool)
//
//func main() {
//	go printer1("hello")
//	go printer2("world")
//	time.Sleep(5 * time.Second)
//	fmt.Println("main...over...")
//}
//func printer1(str string) {
//
//	for _, v := range str {
//		fmt.Printf("%c", v)
//		time.Sleep(100 * time.Millisecond)
//	}
//	<-chan1
//}
//func printer2(str string) {
//
//	chan1 <- true
//	for _, v := range str {
//		fmt.Printf("%c", v)
//		time.Sleep(100 * time.Millisecond)
//	}
//
//}

//01 主程和go程之间等待，不建议使用channel，而是使用同步等待组
/*
	主程和go程直接的关系使用。
	解决问题是等待问题。
*/
//var wg=sync.WaitGroup{}
//func main() {
//
//	wg.Add(2)
//	go printer("hello")
//	go printer("world")
//	wg.Wait()
//	fmt.Println("main...over...")
//}
//
//func printer(str string){
//	for _, v := range str {
//		fmt.Printf("%c",v)
//	}
//	wg.Done()
//}

//02使用互斥锁完成同步
/*
	使用锁的思想是：
		在使用共享数据的时候使用，为了安全访问，使用互斥锁。
		缺点：数据访问效率有所下降！
		优点：增加数据安全性！
			---强制锁：
				底层操作系统时会用到。
            ---建议锁:
                操作系统提供，建议在编程时使用。
				互斥锁为建议锁。
				编程一般都是建议锁。
*/
//var mutex sync.Mutex  //单进程中，锁只有一把。
//func main() {
//	str1:="HelloWorld!!!\n"
//	str2:="Oh,MyGod!!!\n"
//	go printer1(str1)
//	go printer2(str2)
//	time.Sleep(5*time.Second)
//	fmt.Println("main...over...")
//}
//func printer1(str string) {
//	time.Sleep(10*time.Millisecond)
//	mutex.Lock()
//	for _, v := range str {
//		fmt.Printf("%c", v)
//		time.Sleep(100 * time.Millisecond)
//	}
//	mutex.Unlock()
//}
//func printer2(str string) {
//	mutex.Lock()
//	for _, v := range str {
//		fmt.Printf("%c", v)
//		time.Sleep(100 * time.Millisecond)
//	}
//	mutex.Unlock()
//}
