package main

import (
	"strings"
	"fmt"
)

func main()  {
	s := `(76948723, 'N\'cho Denis Laroche ', 0, '02:00:00:00:00:00', '358631083789768', '612051436008540', 98000, 0, 0, 0, '1970-1-1 00:00:00', '{\"1\":5,\"2\":2,\"3\":0,\"4\":0,\"5\":0}', 0, 2, 0)`
	s = strings.Replace(s, "'N\\", "", -1)
	fmt.Println(s)
}
