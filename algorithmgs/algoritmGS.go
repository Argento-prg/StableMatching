package algorithmgs

//проверка, есть ли элемент в массиве
func isElemInArray(container []bool, element bool) bool {
	for _, elem := range container {
		if elem == element {
			return true
		}
	}
	return false
}

//получить индекс элемента в массиве булевых переменных
func findIndexBool(container []bool, element bool) int {
	for id, elem := range container {
		if elem == element {
			return id
		}
	}
	return -1
}

//получить индекс элемента в массиве целочисленных переменных
func findIndexInt(container []int, element int) int {
	for id, elem := range container {
		if elem == element {
			return id
		}
	}
	return -1
}

//удалить элемент из среза по индексу
func remove(s [][2]int, i int) [][2]int {
	var match [][2]int
	for j := 0; j < len(s); j++ {
		if j != i {
			match = append(match, s[j])
		}
	}
	return match
}

func FindStableMatching(manPair [][]int, womenPair [][]int) [][2]int {
	//manPair - подборка мужских пар
	//womenPair - подборка женских пар
	lengthManPair := len(manPair)
	lengthWomanPair := len(womenPair)
	//срез для оценки, являются ли мужчины и женщины одинокими
	isManFreeArray := make([]bool, lengthManPair)
	isWomanFreeArray := make([]bool, lengthWomanPair)
	//по умолчанию, заполняем срезы true
	for i := 0; i < lengthManPair; i++ {
		isManFreeArray[i] = true
	}
	for i := 0; i < lengthWomanPair; i++ {
		isWomanFreeArray[i] = true
	}
	//бланк предложения мужчин женщинам
	isManProposed := make([][]bool, lengthManPair)
	for i := 0; i < lengthManPair; i++ {
		isManProposed[i] = make([]bool, lengthWomanPair)
		for j := 0; j < lengthWomanPair; j++ {
			isManProposed[i][j] = false
		}
	}
	//найденные комбинации
	match := make([][2]int, lengthManPair)
	for i := 0; i < lengthManPair; i++ {
		match[i][0] = -1
		match[i][1] = -1
	}
	//поиск паросочетаний
	for isElemInArray(isManFreeArray, true) {
		indexManFree := findIndexBool(isManFreeArray, true)
		//предложим каждой женщине, у которой нет пары, найти пары по её приоритету мужчин
		if isElemInArray(isManProposed[indexManFree], false) {
			indexWomanFree := -1
			//получить индекс свободной женщины с максимальным рейтингом
			for idx := range manPair[indexManFree] {
				w := manPair[indexManFree][idx]
				if !isManProposed[indexManFree][w] {
					indexWomanFree = w
					break
				}
			}
			//формируем первичную пару
			//перерасчёт индекса
			tempIndexWomanFree := 0
			if indexWomanFree == -1 {
				tempIndexWomanFree = len(isManProposed[indexManFree]) - 1
			} else {
				tempIndexWomanFree = indexWomanFree
			}
			isManProposed[indexManFree][tempIndexWomanFree] = true
			//перерасчёт индекса
			if indexWomanFree == -1 {
				tempIndexWomanFree = len(isWomanFreeArray) - 1
			} else {
				tempIndexWomanFree = indexWomanFree
			}
			//женщина не замужем и готова начать отношения с мужчиной
			if isWomanFreeArray[tempIndexWomanFree] {
				isWomanAccept := false
				//перерасчёт индекса
				if indexWomanFree == -1 {
					tempIndexWomanFree = len(womenPair) - 1
				} else {
					tempIndexWomanFree = indexWomanFree
				}
				for idx := range womenPair[tempIndexWomanFree] {
					if womenPair[tempIndexWomanFree][idx] == indexManFree {
						isWomanAccept = true
					}
				}
				if isWomanAccept {
					//перерасчёт индекса
					if indexWomanFree == -1 {
						tempIndexWomanFree = len(isWomanFreeArray) - 1
					} else {
						tempIndexWomanFree = indexWomanFree
					}
					isWomanFreeArray[tempIndexWomanFree] = false
					isManFreeArray[indexManFree] = false
					match[indexManFree][0] = indexManFree
					match[indexManFree][1] = indexWomanFree
				}
			} else {
				//индекс мужчины, который подходит текущей женщине
				indexManForCurrentWomen := -1
				isWomanAccept := false
				//перерасчёт индекса
				if indexWomanFree == -1 {
					tempIndexWomanFree = len(womenPair) - 1
				} else {
					tempIndexWomanFree = indexWomanFree
				}
				//может ли эта женщина выбрать этого мужчину
				for idx := range womenPair[tempIndexWomanFree] {
					if womenPair[tempIndexWomanFree][idx] == indexManFree {
						isWomanAccept = true
					}
				}
				if isWomanAccept {
					for i := 0; i < lengthManPair; i++ {
						if match[i][1] == indexWomanFree {
							indexManForCurrentWomen = i
							break
						}
					}
					if findIndexInt(womenPair[tempIndexWomanFree], indexManFree) < findIndexInt(womenPair[tempIndexWomanFree], indexManForCurrentWomen) {
						isManFreeArray[indexManForCurrentWomen] = true
						isManFreeArray[indexManFree] = false
						match[indexManFree][0] = indexManFree
						match[indexManFree][1] = indexWomanFree
					}
				}
			}
		}
	}
	//удаляем нестабильные паросочетания
	for i := len(match) - 1; i >= 0; i-- {
		if match[i][0] == -1 || match[i][1] == -1 {
			match = remove(match, i)
		}
	}
	return match
}
