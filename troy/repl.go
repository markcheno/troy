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
				query = troy.V(args[1].(string))
				continue
			}
			if args[0] == "has" {
				query.Has(args[1].(string), args[2].(string))
				continue
			}
			if args[0] == "out" {
				query.Out(args[1].(string))
				continue
			}
			if args[0] == "in" {
				query.In(args[1].(string))
				continue
			}
		}
		v, _ := vm.ToValue(query.Result)
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
            var g = {}
            g.v = function(start) {
                var instructions = []
                instructions.push(["start", start])
                return {
                    all : function() {
                        return getExec(instructions);
                    },
                    out : function(p) {
                        instructions.push(["out", p])
                        return this;
                    },
                    in: function(p) {
                        instructions.push(["in", p])
                        return this;
                    },
                    has : function(p, o) {
                        instructions.push(["has", p, o])
                        return this
                    }
                }
            }

            g.update = function(start) {
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
