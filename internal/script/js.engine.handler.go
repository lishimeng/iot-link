package script

import "github.com/robertkrimen/otto"

type jsEngine struct {
	script string
	vm *otto.Otto
}

func (engine *jsEngine)Invoke(method string, params ...interface{}) (otto.Value, error) {
	res, err := CallFunc(engine.vm, method, params...)
	return res, err
}

func (engine *jsEngine)Inject(name string, callback func(call otto.FunctionCall) otto.Value) {
	_ = engine.vm.Set(name, callback)
}

func (engine *jsEngine) SetValue(name string, value interface{}) error {
	return engine.vm.Set(name, value)
}