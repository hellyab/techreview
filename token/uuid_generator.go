package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	
	for (true){
	id := uuid.New()

		fmt.Printf("new UUID:  %s\n", id)

	}

}