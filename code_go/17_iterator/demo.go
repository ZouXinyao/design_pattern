package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/17_iterator/iterator"
)

func main() {

	user1 := iterator01.NewUser("a", 10)
	user2 := iterator01.NewUser("b", 20)

	set := iterator01.NewUserSet()
	set.Add(user1)
	set.Add(user2)

	iterator := set.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}