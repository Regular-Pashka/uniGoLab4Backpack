package main

// реализовать программу, класс предмет(название, вес, стоимость, соотношение цена/вес),
// ввести макс вес рюкзака (класс рюкзак)
// решить задачу тремя способами (рекурсия, динамика, жадный(любая стратегия))
// ввод: вручную пользователем, через файл
// вывод: список предметов, вес рюкзака, прибыль

import (
	"fmt"
	// "math"
	// "strings"
	"bufio"
	"os"
	"strings"
	"strconv"
)

/* 
	TODO: 

	максимальный вес, текущий вес, список предметов в рюкзаке.
	Реализовать меню:
	1 Заполнение списка предметов из файла
	2 Добавление предмета
	3 Изменение предмета
	4 Удаление предмета
	5 Задание максимального веса рюкзака
	6 Просмотр содержимого рюкзака
	7 Выбор способа решения задачи
	8 Сравнение способов решения
*/

/* 
	Нужно будет реализовать функции добавления и удаления из списка products в main,
	считывания из файла
*/

type Product struct{
	Name string
	Weight int
	Value float64
}

type Backpack struct {
	maxWeight int
	currWeight int
	productList []Product
}

func addProductsFromFile(filePath string) []*Product {
    var result []*Product
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	i := 0
	j := 0
	for reader.Scan() {
		line := reader.Text()
		if i == 0 {
			result = append(result, &Product{Name: getNameFromStr(line)})
		} else if i == 1 {
			result[j].Value = getValueFromStr(line)
		} else {
			result[j].Weight = getWeightFromStr(line)
			j++
			i = -1
		}
		i++
	}
	return result
}

func getNameFromStr(line string) string {
	startIndex := strings.Index(line, ":")
	name := line[startIndex + 2 : ]
	return name
}

func getWeightFromStr(line string) int {
	startIndex := strings.Index(line, ":")
	weight := line[startIndex + 2 : ]
	weightInt, err := strconv.Atoi(weight)
	if err != nil {
        // Handle error
        fmt.Println(err)
        return 1
    }
	return weightInt
}

func getValueFromStr(line string) float64 {
	startIndex := strings.Index(line, ":")
	value := line[startIndex + 2 : ]
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 1
	}
	return valueFloat
}

func printMenu() {
    fmt.Println("Меню")
    fmt.Println("1. Заполнение списка из файла")
    fmt.Println("2. Вывести список предметов в очереди")
    fmt.Println("3. Добавить предмет")
    fmt.Println("4. Изменить предмет")
    fmt.Println("5. Удалить предмет")
    fmt.Println("6. Задать максимальный вес рюкзака")
    fmt.Println("7. Решить задачу")
	fmt.Println("8. Выход из программы")
}

func getPath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите путь к файлу/каталогу: ")
	path, _ := reader.ReadString('\n')
	return strings.TrimSpace(path)
}

func createProduct(name string, weight int, value float64) *Product {
	return &Product{
		Name: name,
		Weight: weight,
		Value: value,
	}
}

func initializeProduct() *Product{
	var name string
	var weight int
	var value float64
	fmt.Print("Введите название предмета: ")
	fmt.Scanf("%s", &name)
	// fmt.Println()
	fmt.Print("Введите вес предмета: ")
	fmt.Scanf("%d", &weight)
	// fmt.Println()
	fmt.Print("Введите цену предмета: ")
	fmt.Scanf("%f", &value)
	// fmt.Println()
	return createProduct(name, weight, value)
}

func setProductName(*Product) {
	fmt.Printf()
}

func setProductName(*Product) {

}

func setProductName(*Product) {

}

func changeProduct(products []*Product) {
	var ind int
	ans := 4
	fmt.Print("Введите индекс предмета в списке: ")
	fmt.Scanf("%s", &ind)
	for ind >= len(products) {
		fmt.Print("Нет такого предмета. Введите индекс предмета в списке: ")
		fmt.Scanf("%s", &ind)
	}
	for ans <= 0 && ans > 3 {
		fmt.Println("Меню изменения предмета:")
		fmt.Println("1. Изменить название")
		fmt.Println("2. Изменить вес")
		fmt.Println("3. Изменить цену")
		fmt.Print("Введите что вы хотите изменить: ")
		fmt.Scanf("%d", &ans)
		switch ans{
		case 1:
			setProductName()
		case 2:
			setProductWeight()
		case 3:
			setProductValue()

		default:
			
		}
	}
	
}

func main() {
	// Initialize an empty list of products
	products := make([]*Product, 0)

	// Initialize a backpack as nil
	// var backpack *Backpack
	
	// Initialize ans to 0
	ans := 0
	for {
		printMenu()
		fmt.Print("> ")
		fmt.Scanln(&ans)
		switch ans {
		case 1:
			path := getPath()
			products = addProductsFromFile(path)
			// fmt.Println("Предметы добавлены", products)
		case 2:
			products = append(products, initializeProduct())
			// fmt.Println("Предметы добавлены", products)
		case 3:
			
		// case 3:
		// 	fmt.Print("Введите стоимость: ")
		// 	var cost int
		// 	fmt.Scanln(&cost)
		// 	fmt.Print("Введите вес: ")
		// 	var Weight int
		// 	fmt.Scanln(&Weight)
		// 	products = append(products, Product{cost, Weight})
		// 	fmt.Println("Предмет добавлен")
		// case 4:
		// 	fmt.Print("Введите индекс предмета для изменения от 1 до ", len(products), ": ")
		// 	// changeProduct(&products[in.nextInt()-1])
		// case 5:
		// 	fmt.Print("Введите номер в списке от 1 до ", len(products), ": ")
		// 	// products = append(products[:in.nextInt()-1], products[in.nextInt():]...)
		// case 6:
		// 	fmt.Print("Введите максимальный вес рюкзака: ")
		// 	var maxWeight int
		// 	fmt.Scanln(&maxWeight)
		// 	backpack = &Backpack{maxWeight: maxWeight}
		// case 7:
		// 	if backpack == nil {
		// 		fmt.Println("Рюкзак не инициализирован!")
		// 		break
		// 	}
		// 	if len(products) == 0 {
		// 		fmt.Println("Список предметов не инициализирован!")
		// 		break
		// 	}
		// 	solver := KnapsackSolver{products: products, backpack: backpack}
		// 	fmt.Println("1. Решение методом динамического программирования")
		// 	fmt.Println("2. Решение жадным алгоритмом (макс. вес)")
		// 	fmt.Println("3. Решение рекурсией")
		// 	fmt.Print("> ")
		// 	var localAns int
		// 	fmt.Scanln(&localAns)
		// 	switch localAns {
		// 	case 1:
		// 		backpack.clearContents()
		// 		backpack = solver.findAnsDP(len(products), backpack.maxWeight)
		// 	case 2:
		// 		backpack.clearContents()
		// 		backpack = solver.findAnsGreedy(len(products), backpack.maxWeight)
		// 	case 3:
		// 		backpack.clearContents()
		// 		backpack = solver.findAnsRec(len(products), backpack.maxWeight)
		// 	default:
		// 		fmt.Println("Неверная команда!")
		// 	}
		case 8:
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Неверная команда!")
		}
	}

}





/* TODO:

Лабораторная работа №4
«Задача о заполнении рюкзака»
В задаче о рюкзаке есть набор предметов. Каждый предмет имеет
название, вес и ценность. Требуется сложить вещи с максимальной
стоимостью в рюкзак, имеющий ограничение по весу.
Реализовать программу решающую задачу о заполнении рюкзака с
1 Рекурсивного метода;
2 Метода динамического программирования;
3 Жадного алгоритма.
Для жадного алгоритма реализовать стратегии:
заполняем в первую очередь предметами с максимальным весом
(нечетные варианты)
заполняем в первую очередь предметами с максимальным
соотношением цена/вес (четные варианты)
В программе должен присутствовать класс «Предмет», обладающий
полями: название, вес, цена; и класс «Рюкзак», обладающим полями:
максимальный вес, текущий вес, список предметов в рюкзаке.
Реализовать меню:
1 Заполнение списка предметов из файла
2 Добавление предмета
3 Изменение предмета
4 Удаление предмета
5 Задание максимального веса рюкзака
6 Просмотр содержимого рюкзака
7 Выбор способа решения задачи
8 Сравнение способов решения
В качестве дополнительного задания можно реализовать заполнение рюкзака
с учетом габаритов вещей. Пользователь может задать размер рюкзака mxn и
размер каждого предмета. В результате работы помимо максимального веса
учитывается геометрическое наполнение рюкзака. 

*/