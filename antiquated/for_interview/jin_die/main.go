package main

func main() {
	ch := make(chan int)
	exit := make(chan struct{})
	go func() {
		defer func() {
			//close(ch)
			close(exit)
		}()
		for i := 1; i < 10; i++ {
			println("g1:", <-ch)
			i++
			if i < 10 {
				ch <- i
			}
		}
	}()

	go func() {
		for i := 0; i < 9; i++ {
			i++
			ch <- i
			println("g2:", <-ch)
		}
	}()

	<-exit
}
