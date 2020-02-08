package main

import (
	"fmt"
	"time"
)

func main() {

	/*go count("sheep")
	go count("fish")*/

	//time.Sleep(time.Millisecond * 2000)

	//fmt.Scanln()

	/*var wg sync.WaitGroup
	wg.Add(1)
	*go func() {
		count1("sheep")
		wg.Done()
	}()

	wg.Wait()*/

	/*c := make(chan string)
	go count2("sheep", c)*/

	/*for {
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println(msg)
	}*/

	/*for msg := range c {
		fmt.Println(msg)
	}*/

	/*c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	c <- "third"
	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)*/

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two secons"
			time.Sleep(time.Second * 2)
		}
	}()

	/*for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}*/

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func count1(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func count2(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
