package utils

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func WriteToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't Serialize to bytes", err)
		return err
	}

	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func ReadFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read file", err)
		return err
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Could't unmarshal", err)
		return err
	}
	return nil
}
