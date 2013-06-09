package assertions

import (
  "testing"
  "runtime"
  "reflect"
  "fmt"
)

func AssertEquals(t *testing.T, f interface{}, args []interface{}, expectedResults []interface{}) {
  argValues     := make([]reflect.Value, len(args))
  for i := 0; i < len(args); i++ {
    argValues[i] = reflect.ValueOf(args[i])
  }

  funcUnderTest := reflect.ValueOf(f)
  funcName      := runtime.FuncForPC(funcUnderTest.Pointer()).Name()
  message       := fmt.Sprintf("%s(%v)", funcName, args)

  actualResults := funcUnderTest.Call(argValues)
  for i := 0; i < len(actualResults); i++ {
    actual    := actualResults[i]
    expected  := expectedResults[i]
    if actual.Interface() != expected {
      t.Errorf("%v did not match.", message)
      t.Errorf("Expected: %v", expected)
      t.Errorf("Actual  : %v", actual.Interface())
    }
  }
}

