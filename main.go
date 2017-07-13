package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Print("CREATE DATABASE IF NOT EXISTS test;")
	fmt.Print("USE test;")
	fmt.Print("DROP TABLE IF EXISTS test;")
	fmt.Print("CREATE TABLE test (")
	fmt.Print("id INT NOT NULL AUTO_INCREMENT,")
	fmt.Print("amount FLOAT(10,4) NOT NULL,")
	fmt.Print("PRIMARY KEY(id)) ENGINE=InnoDB;\n")

	for i := 0; i < 10000; i++ {
		fmt.Print("INSERT INTO test (amount) VALUES ")
		for y :=0 ; y < 999; y++ {
			fmt.Printf("('%f'),", rand.Float32())
		}
		fmt.Printf("('%f');\n", rand.Float32())
	}
}
