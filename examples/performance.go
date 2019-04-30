package main

import (
    "fmt"
    "time"

    "github.com/ypeng7/simgo"
)

func test(env *simgo.Environment, pc *simgo.ProcComm) interface{} {
	for {

		to := simgo.NewTimeout(env, 1, nil)
		to.Schedule(env)
		pc.Yield(to.Event)
	}
}

func main() {
    var start = time.Now()
	env := simgo.NewEnvironment()
	p := simgo.NewProcess(env, simgo.ProcWrapper(env, test))
	p.Init()
	env.Run(10000000)
    fmt.Println("Duration time is ", time.Now().Sub(start))
}
