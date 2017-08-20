package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"../workers"
	"github.com/gorilla/mux"
	"github.com/moovweb/gokogiri"
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

func editTopic(wr http.ResponseWriter, req *http.Request) {
	fmt.Printf("Message /topic/ received:\n%s\n", req.Body)

}

func editWordsDelete(wr http.ResponseWriter, req *http.Request) {

}
func editWordsAdd(wr http.ResponseWriter, req *http.Request) {

	// body, _ := ioutil.ReadAll(req.Body)
	// req.Body.Close()
	fmt.Printf("Message /words.add/ received:\n%s\n", req.Body)

	handler := func() interface{} {

		/*		commandsList := []string{
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
		*/
		body, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		fmt.Printf("Message received:\n%s\n", body)

		if err != nil {
			log.Fatal(err)
		}

		reference := make(map[rune][]string)
		str := strings.Split(string(body), " ")
		var command = str[0]
		ok, _ := contains(commandsList, command)
		if !ok {
			log.Fatalf("Wrong command!")
		}

		//if len(str) > 15 {
		//	panic("Too much arguments")
		//}

		if argsLenDict[command] >= 0 {
			if len(str) != argsLenDict[command]+1 {
				log.Fatalf("Wrong number of arguments")
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
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()

			if err != nil {
				log.Fatal(err)
			}

			doc, _ := gokogiri.ParseHtml(body)
			panelPath := "[@id=\"ColumnCenter\"]/div[2]/div/div[3]/div"
			doc.XPathCtx(panelPath)
			fmt.Printf("%s", body)

		case "stop":

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

		for first, words := range reference {
			fmt.Printf("list %q:\n", first)
			for _, word := range words {
				fmt.Printf("%s, ", word)
			}
			fmt.Printf("\n\n")
		}
		return nil

	}

	pool.AddTaskSyncTimed(handler, time.Second)

}

var pool workers.Pool

// Start a server with @parameter concurency pool size
func Start(concurency int, addr string) {

	// const maxArgs = 10
	//
	// if len(str) < 1 {
	// 	fmt.Printf("help")
	// 	return
	// }
	router := mux.NewRouter()
	router.HandleFunc("/words", editWordsAdd).Methods("POST")
	router.HandleFunc("/words", editWordsDelete).Methods("DELETE")
	router.HandleFunc("add", editWords).Methods("POST")
	pool := workers.NewPool(concurency)
	pool.Run()
	log.Fatal(http.ListenAndServe(addr, router))
}
