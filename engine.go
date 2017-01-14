package otpl

import "github.com/one-go/otpl/runtime"

func New(ilpath string) *runtime.Runtime {
	return runtime.NewRuntime(ilpath, true)
}

// func (inter *Interpreter) Render(data map[string]interface{}, path string, out io.Writer) {

// }

// func Default(data map[string]interface{}) {

// }
