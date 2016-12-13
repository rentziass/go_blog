package main

import "blog_sample/router"

func main() {
	r := router.RegisterRoutes()

	r.Run(":3000")
}
