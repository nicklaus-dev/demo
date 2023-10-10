package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func main() {
	type foo int32
	type bar int32

	m := make(map[any]int32)
	m[foo(1)] = 1
	m[bar(1)] = 2

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func processRequest(userID, authToken string) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	doRequest(ctx)
}

func doRequest(ctx context.Context) {
	fmt.Printf("userID: %v, authToken: %v\n", ctx.Value("userID"), ctx.Value("authToken"))
}
