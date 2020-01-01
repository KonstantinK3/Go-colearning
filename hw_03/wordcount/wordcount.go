package wordcount

import (
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

// Counter is a sctruct for counting words
type Counter struct {
	word  string
	count int
}

// Wordcount counts words in file and returns 5 most used
func Wordcount(file string) []string {

	// Считывание файла (текущая директория - та, откуда запускается main.go)
	text, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("Something is wrong with your file: %s, error is %v", file, err)
	}

	// Сплит нового массива по всему возможному говну. Для Windows возврат строки - /r/n !!!!
	wordArray := regexp.MustCompile("[ ,./;:\n\r`~!@#$%^&*()-=_+—]+").Split(strings.ToLower(string(text)), -1)

	//fmt.Printf("%v", wordArray)

	// Уберем пролетающие ХЗ как пустые слова
	wordArrayClear := []string{}
	for _, value := range wordArray {
		if value != "" {
			wordArrayClear = append(wordArrayClear, value)
		}

	}

	// посчитаем кол-во каждого слова через хешмап
	wordCounterDict := map[string]int{}
	for _, word := range wordArrayClear {
		wordCounterDict[word]++
	}

	// переведем в слайс структур для последующей сортировки и отсортируем
	wordsSliceStruct := []Counter{}
	for value, counter := range wordCounterDict {
		currentWord := Counter{
			word:  value,
			count: counter,
		}
		wordsSliceStruct = append(wordsSliceStruct, currentWord)
	}

	sort.Slice(wordsSliceStruct, func(i, j int) bool {
		return wordsSliceStruct[i].count > wordsSliceStruct[j].count
	})

	// вернем 5 самых частых, учитывая случай, что в слайсе меньше 5 значений
	ans := []string{}

	for _, value := range wordsSliceStruct {
		ans = append(ans, value.word)
	}

	limit := len(ans)
	if len(ans) > 100 {
		limit = 100
	}

	return ans[:limit]
}
