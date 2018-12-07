package main

import "github.com/arcplus/go-lib/micro"

func main() {
	m := micro.New("gateway")

	m.ServeHTTP(micro.Bind(":3000"), Router())

	m.Start()
}
