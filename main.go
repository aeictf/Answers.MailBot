package main

import (
	"os"
	"fmt"
	"log"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func checkWordInLists(check map[rune]map[string]bool, first rune, word []rune) bool {

	alias := map[rune]string {
		'-': "exclude words",
		'!': "essential words",
		'#': "keywords",
	}

	for _, valmap := range check {
		_, ok := valmap[string(word)]
		if ok {
			log.Fatalf("%s is in %s already.\n", word, alias[first])
			return true
		}
	}
	return false
}

func addWords(reference map[rune]map[string]bool, words []string) {

	for _, r := range words {
		runes := []rune(r)
		first := runes[0]
		var word []rune
		// Смотрим на первый символ ключевого слова
		switch first {
		case '-':
			word = runes[1:]
		case '!':
			word = runes[1:]
		default:
			// Обычное ключевое слово, обозначим через '#'
			first = '#'
			word = runes[:]
		}
		f := checkWordInLists(reference, first, word)
		if f {
			log.Fatalf("Keyword error: ")
		}
		reference[first][string(runes[1:])] = true
	}
}

func delKeyword(word string) {

}

func getNewsSubjects(keywords []string) []string { //сделать структуру: заголовок, первое предложение из статьи, ссылка
	ret := []string{}
	return ret
}

func parse_arg(cmd string, arg string, num int) {

}

func main() {

	commandsList := []string {
		"start",
		"end",
		"topic",
		"add",
		"del",
		"sort",
	}

	argsLenDict := map[string]int {
		"start": 0,
		"end": 0,
		"topic": 1,
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

	reference := map[rune]map[string]bool{}

/*	for i, r := range os.Args[2:] {
		parse_arg(command, r, i)
	}
*/
	switch command {
	case "topic":
		setTopic(os.Args[2])
	case "add":
		addWords(reference, os.Args)
		break
	case "del":
		break
	case "sort":
		break
	}
}
