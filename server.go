package main

import . "test_gin/router"

func main() {
	router := InitROuter()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

