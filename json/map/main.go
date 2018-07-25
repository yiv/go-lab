package main

import (
	"encoding/json"
	"fmt"
)
const TaskConfig = `[{"id":1,"times":5,"coin":3000,"gem":0,"growth":5},{"id":2,"times":5,"coin":4000,"gem":0,"growth":5},{"id":3,"times":5,"coin":5000,"gem":0,"growth":5},{"id":4,"times":1,"coin":5000,"gem":0,"growth":5},{"id":5,"times":50000,"coin":10000,"gem":0,"growth":5},{"id":6,"times":5,"coin":3000,"gem":0,"growth":0},{"id":7,"times":5,"coin":3000,"gem":0,"growth":0},{"id":8,"times":1,"coin":0,"gem":5,"growth":0},{"id":9,"times":3,"coin":5000,"gem":0,"growth":3},{"id":10,"times":5,"coin":5000,"gem":0,"growth":3},{"id":11,"times":10,"coin":10000,"gem":0,"growth":3},{"id":12,"times":20,"coin":10000,"gem":0,"growth":10},{"id":13,"times":1,"coin":5000,"gem":0,"growth":0},{"id":14,"times":1,"coin":3000,"gem":0,"growth":0}]`
const TaskConfig2 = `{"1":{"id":1,"times":5,"coin":3000,"gem":0,"growth":5},"2":{"id":2,"times":5,"coin":4000,"gem":0,"growth":5},"3":{"id":3,"times":5,"coin":5000,"gem":0,"growth":5},"4":{"id":4,"times":1,"coin":5000,"gem":0,"growth":5},"5":{"id":5,"times":50000,"coin":10000,"gem":0,"growth":5},"6":{"id":6,"times":5,"coin":3000,"gem":0,"growth":0},"7":{"id":7,"times":5,"coin":3000,"gem":0,"growth":0},"8":{"id":8,"times":1,"coin":0,"gem":5,"growth":0},"9":{"id":9,"times":3,"coin":5000,"gem":0,"growth":3},"10":{"id":10,"times":5,"coin":5000,"gem":0,"growth":3},"11":{"id":11,"times":10,"coin":10000,"gem":0,"growth":3},"12":{"id":12,"times":20,"coin":10000,"gem":0,"growth":10},"13":{"id":13,"times":1,"coin":5000,"gem":0,"growth":0},"14":{"id":14,"times":1,"coin":3000,"gem":0,"growth":0}}`
type Task struct {
	//任务ID
	Id int32 `json:"id"`
	//任务状态
	Status int32 `json:"status"`
	//任务目标数量
	Times int32 `json:"times"`
	//奖励
	Gem    int32 `json:"gem"`
	Coin   int64 `json:"coin"`
	Growth int32 `json:"growth"`
}

type DayTask map[int32]*Task

func main() {
	dt := make(DayTask)
	err := json.Unmarshal([]byte(TaskConfig2), &dt)
	if err != nil {
		panic(err.Error())
	}

	for _, tk := range dt {
		fmt.Println(tk.Coin)
	}

}
