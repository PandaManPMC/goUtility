package util

import (
	"fmt"
	"testing"
)

func TestNewNickNameGenerator(t *testing.T) {
	// https://gist.github.com/hugsy/8910dc78d208e40de42deb29e62df913
	adjs := []string{
		"Lucky", "Happy", "Cool", "Magic", "Sweet",
	}

	nouns := []string{
		"Girl", "Boy", "Cat", "Tiger", "Mike", "Anna",
	}

	gen, err := NewNickNameGenerator(
		adjs,
		nouns,
		2,    // 最短长度
		12,   // 最长长度
		1000, // 000~999
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("total =", gen.Total())

	for i := uint64(0); i < 10; i++ {
		name, _ := gen.Nickname(i)
		fmt.Println(name)
	}
}
