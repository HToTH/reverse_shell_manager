package lib

import(
	"reflect"
)

func GetAllMethods(any interface{}) []string{
	var methods []string
	anyType := reflect.TypeOf(any)
	for i := 0; i < anyType.NumMethod(); i++ {
		method := anyType.Method(i)
		methods = append(methods, method.Name)
	}
	return methods
}

func InvokeFunc(any interface{},name string,args ...interface{}){
	params := make([]reflect.Value, len(args))
	for i, _ := range args {
		params[i] = reflect.ValueOf(args[i])
	}

	reflect.ValueOf(any).MethodByName(name).Call(params)
}