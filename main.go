package main

import (
	"dockerfileparse/user/parser/babashka"
	"dockerfileparse/user/parser/docker"
)

func main() {
	for {
		message, err := babashka.ReadMessage()
		if err != nil {
			babashka.WriteErrorResponse(message, err)
			continue
		}

		res, err := docker.ProcessMessage(message)
		if err != nil {
			babashka.WriteErrorResponse(message, err)
			continue
		}

		describeres, ok := res.(*babashka.DescribeResponse)
		if ok {
			babashka.WriteDescribeResponse(describeres)
			continue
		}
		// TODO don't write done responses when callback is running
		babashka.WriteInvokeResponse(message, res)
	}
}

