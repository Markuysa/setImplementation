package main

import (
	"fmt"
	"github.com/Markuysa/setImplementation/set/set"
)

func main() {

	firtSet := set.New()
	secondSet := set.New()
	firtSet.Fill(1, 2, 3, 4, 5)
	secondSet.Fill(1, 2, 4, 5, 6)
	firtSet.Intersection(secondSet)
	fmt.Println(firtSet.GetItems())

}
