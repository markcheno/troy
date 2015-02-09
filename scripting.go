package troy

import (
	"github.com/robertkrimen/otto"
)

var vm *otto.Otto

func Script(script string) interface{} {
	value, _ := vm.Run(script)
	if value != otto.UndefinedValue() {
		exp, _ := value.Export()
		return exp
	}
	return nil
}

func init() {
	vm = otto.New()
	vm.Set("getExec", func(call otto.FunctionCall) otto.Value {
		export, _ := call.Argument(0).Export()
		instructions := export.([]interface{})
		var query *Query
		for _, n := range instructions {

			args := n.([]interface{})
			if args[0] == "start" {
				arr := []string{}
				for _, s := range args[1].([]interface{}) {
					arr = append(arr, s.(string))
				}
				query = V(arr...)
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
			if args[0] == "group" {
				v, _ := vm.ToValue(query.Group())
				return v

			}
		}
		v, _ := vm.ToValue(query.Result)
		return v
	})

	vm.Set("updateExec", func(call otto.FunctionCall) otto.Value {
		export, _ := call.Argument(0).Export()
		instructions := export.([]interface{})
		write := Update(instructions[0].(string))
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
        g.v = function() {
            var instructions = []
            instructions.push(["start", Array.prototype.slice.call(arguments)])
            return {
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
                },
                all : function() {
                    return getExec(instructions);
                },
                group : function() {
                    instructions.push(["group"])
                    return getExec(instructions)
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
}
