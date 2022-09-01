package main

import "fmt"

//  func fib(n int) chan int {
// 	out := make(chan int)

// 	go func() {
// 		defer close(out)
// 		for i, j := 0, 1; i < n; i, j = i+j, i {
// 			out <- i
// 		}

// 	}()

// 	return out
// }

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

func main() {
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
