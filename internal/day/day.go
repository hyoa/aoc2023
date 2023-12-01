package day

type ExecDay func(input1, input2 string) (any, any)

var DayCollection = map[string]ExecDay{}
