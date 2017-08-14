package main

import (
	"os"
	"fmt"
)

func addKeyword(s map[string]bool, word string) {

}

func delKeyword(word string) {

}

func getNewsSubjects(keywords []string) []string { //сделать структуру: заголовок, первое предложение из статьи, ссылка
	ret := []string{}
	return ret
}

func parse_arg(cmd string, arg string, num int) {
	
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	commandsList := []string{
		"start",
		"end",
		"add",
		"del",
		"sort",
	}

	argsLenDict := map[string]int{
		"start": 0,
		"end": 0,
		"add": -1,
		"del": -1,
		"sort": 1,
	}

	const maxArgs = 10

	if len(os.Args) < 1 {
		fmt.Printf("help")
		return
	}

	var command string = os.Args[1]

	if !contains(commandsList, command) {
		panic("Wrong command!") // Заменить панику на что-то более спокойное.
	}

	if len(os.Args) > 10 {
		panic("Too much arguments")
	}

	if(len(os.Args) >= 0) {
		if (len(os.Args) != argsLenDict[command] + 1) {
			panic("Wrong number of arguments")
		}
	}

	keyWords := map[string]bool{}

	for _, r := range os.Args[1:argsLenDict[command] + 1] {
		parse_arg(command, r, i)
	}
}
