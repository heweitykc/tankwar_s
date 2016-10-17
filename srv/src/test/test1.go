package main

import (
	"fmt"
	"time"
//	"flag"	
)

func main() {	
	var ticker = time.NewTicker(time.Second)
	var cnt = 0
	for{
		select{
			case <-ticker.C:
				fmt.Printf("loop: cnt=%d\n", cnt)
			default:
				cnt++
				if cnt%1e6 == 0 {
					fmt.Printf("M")
				}
		}
	}
}
