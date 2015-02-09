package main

import (
	"fmt"
	"github.com/ironbay/troy"
	"github.com/peterh/liner"
	"github.com/robertkrimen/otto"
)

func repl() {

	line := liner.NewLiner()
	defer line.Close()

	vm := otto.New()
	vm.Set("getExec", func(call otto.FunctionCall) otto.Value {
		export, _ := call.Argument(0).Export()
		instructions := export.([]interface{})
		var query *troy.Query
		for _, n := range instructions {

			args := n.([]interface{})
			if args[0] == "start" {
				query = troy.Get(args[1].(string))
				continue
			}
			if args[0] == "v" {
				query.V(args[1].(string))
				continue
			}
			if args[0] == "out" {
				query.Out(args[1].(string))
				continue
			}
			if args[0] == "all" {
				query.All()
				continue
			}
		}
		v, _ := vm.ToValue(query.Vertices)
		return v
	})

	vm.Set("updateExec", func(call otto.FunctionCall) otto.Value {
		export, _ := call.Argument(0).Export()
		instructions := export.([]interface{})
		write := troy.Update(instructions[0].(string))
		for i, a := range instructions {
			if i == 0 {
				continue
			}
			write = write.Out(a.(string))
		}
		write.Exec()
		v, _ := vm.ToValue(true)
		return v
	})

	vm.Run(`
            var get = function(start) {
                var instructions = []
                instructions.push(["start", start])
                return {
                    v : function(v) {
                        instructions.push(["v", v])
                        return this;
                    },
                    all : function() {
                        instructions.push(["all"])
                        return this;
                    },
                    out : function(p) {
                        instructions.push(["out", p])
                        return this;
                    },
                    exec : function() {
                        return getExec(instructions);
                    }
                }
            }

            var update = function(start) {
                var instructions = [start];
                var f = function(v) {
                    instructions.push(v);
                    return r;
                }
                var r = {
                    v : f,
                    out : f,
                    exec : function() {
                        return updateExec(instructions)
                    }
                }
                return r;
            }
        `)

	for {
		l, err := line.Prompt("troy> ")
		if err != nil {
			break
		}
		value, _ := vm.Run(l)
		fmt.Println(value)
	}

}
