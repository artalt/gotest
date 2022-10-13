package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countLetters("aaabbbccccc"))
	fmt.Println(countLetters("aaabbbcccccaaaaa"))
	fmt.Println(countLetters("zzzzcccUUUuu"))
	fmt.Println(countLetters("ЯЯЯБББддд"))
}

func countLetters(str string) string {
	// подсчет количества букв
	counts := make(map[string]int)
	for _, letter := range strings.Split(str, "") {
		counts[letter]++
	}

	// сортировка
	keys := make([]string, len(counts))
	i := 0
	for k := range counts {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// подготовка результата в виде строки
	result := make([]string, len(counts))
	i = 0
	for _, k := range keys {
		result[i] = k + strconv.Itoa(counts[k])
		i++
	}

	return strings.Join(result, "")
}
