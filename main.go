package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"./workers"
)

func contains(s []string, e string) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}

func checkWordInLists(check map[rune][]string, first rune, word string) bool {

	alias := map[rune]string{
		'-': "exclude words",
		'!': "essential words",
		'#': "keywords",
	}

	for key, valmap := range check {
		ok, _ := contains(valmap, word)
		if ok {
			log.Printf("%s: Word is in %s already.\n", string(word), alias[key])
			return true
		}
	}
	return false
}

func parseWord(str string) (first rune, word string) {
	runes := []rune(str)
	first = runes[0]
	// Смотрим на первый символ ключевого слова
	switch first {
	case '-':
		word = string(runes[1:])
	case '!':
		word = string(runes[1:])
	default:
		// Обычное ключевое слово, обозначим через '#'
		first = '#'
		word = string(runes[:])
	}

	return first, word
}

func addWords(reference map[rune][]string, words []string) {

	for _, r := range words {

		first, word := parseWord(r)
		f := checkWordInLists(reference, first, word)
		if f {
			log.Printf("Keyword error:\n")
			continue
		}
		_, ok := reference[first]
		if !ok {
			reference[first] = []string{}
		}
		reference[first] = append(reference[first], word)
	}
}

func delWords(reference map[rune][]string, words []string) {

	alias := map[rune]string{
		'-': "exclude words",
		'!': "essential words",
		'#': "keywords",
	}

	for _, r := range words {

		first, word := parseWord(r)
		// _, ok := reference[first][string(word)]
		// if !ok {
		// 	log.Fatalf("Keyword error:\n%s: No such word in %s\n", string(word), alias[first])
		// }
		ok, i := contains(reference[first], word)
		if ok {
			reference[first] = append(reference[first][:i], reference[first][i+1:]...)
		} else {
			log.Printf("Keyword error:\n%s: No such word in %s\n", string(word), alias[first])
		}
	}
}

func getNewsSubjects(keywords []string) []string { //сделать структуру: заголовок, первое предложение из статьи, ссылка
	ret := []string{}
	return ret
}

func main() {

	commandsList := []string{
		"start",
		"stop",
		"topic",
		"add",
		"del",
		"sort",
	}

	argsLenDict := map[string]int{
		"start": 0,
		"stop":  0,
		"topic": 1,
		"add":   -1,
		"del":   -1,
		"sort":  1,
	}

	// const maxArgs = 10
	//
	// if len(str) < 1 {
	// 	fmt.Printf("help")
	// 	return
	// }

	pool := workers.NewPool(5)

	reference := make(map[rune][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		str := strings.Split(scanner.Text(), " ")
		var command = str[0]
		ok, _ := contains(commandsList, command)
		if !ok {
			panic("Wrong command!") // Заменить панику на что-то более спокойное.
		}

		//if len(str) > 15 {
		//	panic("Too much arguments")
		//}

		if argsLenDict[command] >= 0 {
			if len(str) != argsLenDict[command]+1 {
				panic("Wrong number of arguments")
			}
		}

		/*	for i, r := range str[2:] {
				parse_arg(command, r, i)
			}
		*/
		switch command {
		case "start":
			res, err := http.Get("http://otvet.mail.ru/search/" + strings.Join(reference['#'], " "))
			if err != nil {
				log.Fatal(err)
			}
			robots, err := ioutil.ReadAll(res.Body)
			res.Body.Close()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", robots)
			// pool.Run()
		case "stop":
			pool.Stop()
		case "topic":
			//setTopic(str[2])
		case "add":
			addWords(reference, str[1:])
			break
		case "del":
			delWords(reference, str[1:])
			break
		case "sort":
			break
		}
	}
	for first, words := range reference {
		fmt.Printf("list %q:\n", first)
		for _, word := range words {
			fmt.Printf("%s, ", word)
		}
		fmt.Printf("\n\n")
	}
}
