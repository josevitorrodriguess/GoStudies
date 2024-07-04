package competition

import (
	"fmt"
	"time"
)

func Write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}