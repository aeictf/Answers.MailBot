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
			log.Printf("%s: Word is in %s already.\n", string(word), alias[first])
			return true
		}
	}
	return false
}

func parseWord(str string) (first rune, word []rune) {
	runes := []rune(str)
	first = runes[0]
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

	return first, word
}

func addWords(reference map[rune]map[string]bool, words []string) {

	for _, r := range words {

		first, word := parseWord(r)
		f := checkWordInLists(reference, first, word)
		if f {
			log.Printf("Keyword error:\n")
			continue
		}
		_, ok := reference[first]
		if !ok {
			reference[first] = make(map[string]bool)
		}
		reference[first][string(word)] = true
	}
}

func delWords(reference map[rune]map[string]bool, words []string) {

	alias := map[rune]string {
		'-': "exclude words",
		'!': "essential words",
		'#': "keywords",
	}

	for _, r := range words {

		first, word := parseWord(r)
		_, ok := reference[first][string(word)]
		if !ok {
			log.Fatalf("Keyword error:\n%s: No such word in %s\n", string(word), alias[first])
		}
		delete(reference[first], string(word))
	}
}

func getNewsSubjects(keywords []string) []string { //сделать структуру: заголовок, первое предложение из статьи, ссылка
	ret := []string{}
	return ret
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

	//if len(os.Args) > 15 {
	//	panic("Too much arguments")
	//}

	if(argsLenDict[command] >= 0) {
		if (len(os.Args) != argsLenDict[command] + 1) {
			panic("Wrong number of arguments")
		}
	}

	reference := map[rune]map[string]bool{}
	reference = make(map[rune]map[string]bool)
/*	for i, r := range os.Args[2:] {
		parse_arg(command, r, i)
	}
*/
	switch command {
	case "topic":
		//setTopic(os.Args[2])
	case "add":
		addWords(reference, os.Args[2:])
		break
	case "del":
		delWords(reference, os.Args[2:])
		break
	case "sort":
		break
	}

	for first, valmap := range reference {
		fmt.Printf("list %q:\n", first)
		for key, _ := range valmap {
			fmt.Printf("%s, ", key)
		}
		fmt.Printf("\n\n")
	}
}
