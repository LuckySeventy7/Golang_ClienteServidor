  
package main

import (
	"fmt"
	"net"
	"time"
	"encoding/gob"

)

type Proceso struct {
	ID int
	Timer int
	ChanCont  chan int 
	ChanID  chan int 
	
}

func cliente() {

	var p Proceso
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("err1")
		fmt.Println(err)
		return
	}

	err= gob.NewDecoder(c).Decode(&p)
	if err != nil {
		fmt.Println("err2")
		fmt.Println(err)
	}

		if &p!= nil {
		cl := make(chan int)
		go show(&p, cl)
		go func() {
			for{	
			t:= <- cl 
			err:= gob.NewEncoder(c).Encode(t)//enviar ID del cliente
				if err != nil {
					fmt.Println(err)
				//	return
				} 
			}

		}()
	}

}

func show(p* Proceso, cl chan int){
	j:=p.ID
	k:=p.Timer
	for{
		fmt.Println(j, k)
		time.Sleep(time.Millisecond * 500)
		k++
		cl<-j		
	}
}
func main()  {
	go cliente()	
	var input string
	fmt.Scanln(&input)
	fmt.Println("fin")
}