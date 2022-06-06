// Golang Юниор, [06.06.2022 12:01]
// Задача. Поменяйте местами значения переменных без использования промежуточной переменной.

// Реализуйте swap(), обменивающую значения двух переменных, не используя третью переменную.

//Код с решением будет завтра.

//#задача

package main

import "fmt"

func main() {
	name := "Hello! "
	second := "World!! "
	fmt.Println(name + second)
	name, second = second, name
	fmt.Println(name + second)

}
