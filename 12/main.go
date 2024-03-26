package main

import "fmt"

func main() {

	// создаем массив
	wordsArr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("последовательность:", wordsArr)

	// создаем мапу и массив для будущей последовательности
	hashMap := make(map[string]struct{})
	var res []string

	// проходим по каждому слову
	for _, value := range wordsArr {
		//если в мапе оно уже есть - то пропускаем
		if _, ok := hashMap[value]; ok {
			continue
		}

		// добавляем в мапу
		hashMap[value] = struct{}{}
		// пушим в массив результатов
		res = append(res, value)
	}

	fmt.Printf("множество: %v\n", res)
}
