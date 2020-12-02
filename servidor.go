package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"time"
	"sort"
	
)

type Proceso struct {
	ID int
	Timer int
	ChanCont chan int 
	ChanID chan int 
	
}

func servidor()  {


	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("err1 ")
		fmt.Println(err)
		return
	}

		
	sl:= Proceso{ID:0, ChanCont: make(chan int),  ChanID: make(chan int) }
	arr:= []int{0,1,2,3,4}
	go sl.start(&arr)
	go sl.stop()
	ind:=0 //iterador
	//x:=0
	

	for {
		if len(arr)>0{
			c, err := s.Accept()
			arr = removeIndex(arr, 0)//se elimina elemento de la primer posicion del slice
			sl.ID=ind
			ind++
			if err != nil {
				fmt.Println(err)
				continue
			}

			go handleClient(c, &sl, &arr)
			if ind==5{
				ind=0
			}
		}
	}
	
	
}

func handleClient( c net.Conn, s *Proceso, arr* []int){

	i := <-s.ChanCont
	s.Timer =i
	err := gob.NewEncoder(c).Encode(s)
	if err != nil {
		fmt.Println(err)
		//return
	} 

	var aux int
	var aux2 int
	for {
		err = gob.NewDecoder(c).Decode(&aux)
		if err != nil {//al momento que se interrupa con shift C y genere un error, agregamos de nuevo el id al servidor
			*arr = append(*arr, aux2)
			sort.Ints(*arr)
			return	
		} 
		aux2 =aux
		
	}
		//c.Close()
	
}

func removeIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

func ( s *Proceso) start(arr* []int){
	cont:=0
	
	for{
		p:= *arr // actualuzamos el ID cuando se termine de imprimir el contador para todo los ID
		if len(p)>0{
			fmt.Println("--------------------")
			sort.Ints(p)
			for z, x:=range p{		
				fmt.Println(x,":", cont)	
				if z == len(p)-1{
					//time.Sleep(time.Millisecond * 500)
					cont = <-s.ChanCont
				}			
			}
		}else{//para que se siga actualizando el contador hasta cuando no tenga ningun ID en el servidor
		
			fmt.Println(":", cont) 
			cont = <-s.ChanCont
		}
	}

}
func (s *Proceso) stop() {
	u:=0
	v:=0
	for {
		select {
		
		case <-s.ChanID:
			fmt.Println("placeholde")
		case <-s.ChanCont:
			fmt.Println("placeholder")
		
		default:

			time.Sleep(time.Millisecond * 500)
			s.ChanCont <-v
			u++
			v++
		}

	}
}

func main()  {

	//n:=0
	go servidor()
	var input string
	fmt.Scanln(&input)
	
}