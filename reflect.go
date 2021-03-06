// Notes on reflect package

package main

import (
  "reflect"
  "log"
  "os"
)

type Person struct {
  Name string
  Age int
  IsMarried bool
}

var logger = log.New(os.Stdout, "mylog: ", log.Ldate+log.Ltime+log.Lshortfile)

func main() {
  var p Person
  GenricFunc(&p)  
}

func GenricFunc (structPtr interface{}) error {
  // check whether structPtr is a pointer
  value := reflect.ValueOf(structPtr)
  if value.Kind() == reflect.Ptr {
    logger.Println("%v is a pointer \n", structPtr)
  }
  return nil
}

func generic(strct interface{}){
  val := reflect.TypeOf(strct)
	value := reflect.New(val) // creates a pointervalue for the given type
}

// useful utility taken from https://github.com/jmoiron/sqlx package
func baseType(t reflect.Type, expected reflect.Kind) (reflect.Type, error) {
	t = deref(t)
	if t.Kind() != expected {
		return nil, fmt.Errorf("expected %s but got %s", expected, t.Kind())
	}
	return t, nil
}

func deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
