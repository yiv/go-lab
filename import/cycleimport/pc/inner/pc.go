package inner

import (
	"fmt"

	painner "github.com/yiv/go-lab/import/cycleimport/pa/inner"
	pbinner "github.com/yiv/go-lab/import/cycleimport/pb/inner"
)

func Pcfun() {
	fmt.Println("pc")
	painner.Pafun()
	pbinner.Pbfun()
}
