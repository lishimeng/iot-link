package script

import (
	"errors"
	"fmt"
	"github.com/robertkrimen/otto"
	"time"
)

var maxExecutionTime = 10 * time.Millisecond

type JsEngine interface {
	Invoke(method string, params ...interface{}) (otto.Value, error)
	Inject(name string, callback func(call otto.FunctionCall) otto.Value)
	SetValue(name string, value interface{}) error
}

func Create (javascript string) (engine JsEngine, err error) {

	vm := otto.New()
	vm.Interrupt = make(chan func(), 1)
	vm.SetStackDepthLimit(32)

	err = LoadScript(vm, javascript)

	if err == nil {
		js := jsEngine{
			script:javascript,
			vm: vm,
		}
		engine = &js
	}

	return engine, err
}

func LoadScript(vm *otto.Otto, script string) (err error) {
	defer func() {
		if exp := recover(); exp != nil {
			err = fmt.Errorf("%s", exp)
		}
	}()
	go func() {
		time.Sleep(maxExecutionTime)
		vm.Interrupt <- func() {
			panic(errors.New("execute javascript timeout"))
		}
	}()
	_, err = vm.Run(script)
	return err
}

func CallFunc(vm *otto.Otto, method string, params ...interface{}) (value otto.Value, err error) {
	defer func() {
		if exp := recover(); exp != nil {
			err = fmt.Errorf("%s", exp)
		}
	}()

	// 执行js前打开线程做超时检查
	go func() {
		time.Sleep(maxExecutionTime)
		vm.Interrupt <- func() {
			panic(errors.New("execute javascript timeout"))
		}
	}()

	value, err = vm.Call(method, nil, params...)
	return value, err
}

func Execute(vm *otto.Otto, script string, params ...interface{}) (value otto.Value, err error) {

	defer func() {
		if exp := recover(); exp != nil {
			err = fmt.Errorf("%s", exp)
		}
	}()

	// 执行js前打开线程做超时检查
	go func() {
		time.Sleep(maxExecutionTime)
		vm.Interrupt <- func() {
			panic(errors.New("execute javascript timeout"))
		}
	}()

	value, err = vm.Run(script)
	if err != nil {
		return value, err
	}
	if value.IsFunction() {
		return value, errors.New("not support javascript return type is 'function'")
	}

	_, _ = vm.Call("", nil, params...)

	return value, err
}