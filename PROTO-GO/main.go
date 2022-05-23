package main

import (
	"fmt"
	"reflect"

	pb "example.com/m/proto"
	"example.com/m/utils"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:       42,
		IsSimple: true,
		Name:     "A name",
		SampleLists: []int32{
			1, 2, 3, 4, 5, 6,
		},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 42, Name: "My Name"},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "My Name 2"},
			{Id: 44, Name: "My name 2"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}

func doOneOf(msg interface{}) {
	switch x := msg.(type) {
	case *pb.Result_Id:
		fmt.Println(msg.(*pb.Result_Id).Id)
	case **pb.Result_Message:
		fmt.Println(msg.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("messsage has unexpected type: %v", x)
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	utils.WriteToFile(path, p)
	message := &pb.Simple{}
	utils.ReadFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := utils.ToJSON(p)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	utils.FromJSON(jsonString, message)
	return message
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doComplex())
	fmt.Println(doEnum())
	fmt.Println(doMap())

	fmt.Println("This should be an Id:")
	doOneOf(&pb.Result_Id{Id: 42})
	fmt.Println("This should be a Message:")
	doOneOf(&pb.Result_Message{Message: "the message is here"})

	doFile(doSimple())

	jsonString := doToJSON(doComplex())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	fmt.Println(doFromJSON(`{"id":42, "unknow":"test"}`, reflect.TypeOf(pb.Simple{})))
}
