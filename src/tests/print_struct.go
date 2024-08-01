package tests

import "log"

func PrintStruct(i interface{}) {
	log.Printf("%+v", i)
}
