/*
 * Helper functions for testing. Includes a little gratuitous reflection.
 */

package assertions

import (
  "testing"
  "runtime"
  "reflect"
  "fmt"
  "bytes"
  "subakva/matasano/utils"
)

func AssertEquals(t *testing.T, actual interface{}, expected interface{}, message string) {
  if actual != expected {
    t.Errorf("Assertion failed: %v", message)
    t.Errorf("\tExpected: %v", expected)
    t.Errorf("\tActual  : %v", actual)
  }
}

func AssertStringArraysEqual(t *testing.T, actual [][]byte, expected []string) {
  actualStrings := utils.BytesToStrings(actual)
  if len(actual) != len(expected) {
    t.Errorf("Actual length does not match expected: %v != %v", len(actual), len(expected))
    t.Errorf("Actual  : %v\n", actualStrings)
    t.Errorf("Expected: %v\n", expected)
  } else {
    for i := 0; i < len(actual); i++ {
      if ! bytes.Equal(actual[i], []byte(expected[i])) {
        t.Errorf("Actual does not match expected at index: %v", i)
        t.Errorf("Actual  : %v\n", actualStrings)
        t.Errorf("Expected: %v\n", expected)
        t.Errorf("Actual[%v]  : %v\n", i, actualStrings[i])
        t.Errorf("Expected[%v]: %v\n", i, expected[i])
      }
    }
  }
}

func mapArgsToValues(args []interface{}) []reflect.Value {
  argValues := make([]reflect.Value, len(args))
  for i := 0; i < len(args); i++ {
    argValues[i] = reflect.ValueOf(args[i])
  }
  return argValues
}

func mapValuesToResults(values []reflect.Value) []interface{} {
  results := make([]interface{}, len(values))
  for i := 0; i < len(values); i++ {
    results[i] = values[i].Interface()
  }
  return results
}

func formatMessage(f interface{}, args []interface{}) string {
  funcUnderTest := reflect.ValueOf(f)
  funcName      := runtime.FuncForPC(funcUnderTest.Pointer()).Name()
  return fmt.Sprintf("%s(%v)", funcName, args)
}

func callFunction(f interface{}, args []interface{}) []interface{} {
  funcUnderTest := reflect.ValueOf(f)
  argValues     := mapArgsToValues(args)
  resultValues := funcUnderTest.Call(argValues)
  return mapValuesToResults(resultValues)
}

func CallAndAssertEquals(t *testing.T, f interface{}, args []interface{}, expectedResults []interface{}) {
  message       := formatMessage(f, args)
  actualResults := callFunction(f, args)

  for i := 0; i < len(actualResults); i++ {
    actual    := actualResults[i]
    expected  := expectedResults[i]
    AssertEquals(t, actual, expected, fmt.Sprintf("%v did not match.", message))
  }
}
