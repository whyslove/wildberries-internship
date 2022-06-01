package main

func main() {
	//1ый способ (очень плохой)
	// go func(){
	// 	time.Sleep(time.Minute)
	// }
	// then return from main

	//2ой способ
	// abort := make(chan struct{})
	// go func() {
	// 	select {
	// 	case <-abort:
	// 		return
	// 	default:
	// 		fmt.Println("123")
	// 	}
	// }()
	// abort <- struct{}{}
	// //Если много подпрограмм, то придется тяжко

	//3ий способ

	// abort := make(chan struct{})
	// go func(){
	// 	select{
	// 	case <- abort:
	// 		return
	// 	default:
	// 		fmt.Println("123")
	// 	}
	// }
	// close(abort)
	// самый лучший метод
}
