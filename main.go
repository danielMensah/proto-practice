package main

import (
	"fmt"

	pb "github.com/danielMensah/proto-practice/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         123,
		IsSimple:   true,
		Name:       "Daniel",
		SimpleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		Dummy: &pb.DummyMessage{
			Id:   123,
			Name: "Daniel",
		},
		Dummies: []*pb.DummyMessage{
			{
				Id:   123,
				Name: "Daniel",
			},
			{
				Id:   123,
				Name: "Daniel",
			},
		},
	}
}

func doEnumeration() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BLUE,
	}
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doComplex())
	fmt.Println(doEnumeration())
}
