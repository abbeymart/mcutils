// @Author: abbeymart | Abi Akindele | @Created: 2020-12-03 | @Updated: 2020-12-03
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect go-helper functions

package mcutils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PanicOnError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Panicf("Error: %v", err)
	}
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}
