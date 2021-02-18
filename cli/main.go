package main

import "gmock/web"

func main() {
	if err := web.Run(); err != nil {
		panic(err)
	}
}
