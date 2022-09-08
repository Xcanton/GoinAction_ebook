package main

import (
	"log"
	"os"

	"modeldemo/search_deom/search"
)

func init() {
	// 将输出重定向到系统输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("present")
}
