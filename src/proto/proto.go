// golang protobuf
// refered to: https://tutorialedge.net/golang/go-protocol-buffer-tutorial/

package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {

	elliot := &Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2323,
			Twitter: 32323,
		},
	}
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)
	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetSocialFollowers())

}
