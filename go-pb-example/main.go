package main

import (
	"log"

	"code.google.com/p/goprotobuf/proto"
	"github.com/sbinet/pb-example/example"
)

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatalf("marshalling error: %v\n", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatalf("unmarshaling error: %v\n", err)
	}
	// now test and newTest contain the same data.
	if *test.Label != *newTest.Label {
		log.Fatalf("data mismatch %q != %q\n", *test.Label, *newTest.Label)
	}
	log.Printf("ref-test: %v\n", test)
	log.Printf("new-test: %v\n", newTest)
}
