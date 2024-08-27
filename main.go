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

type Product struct{
	Name string
	Weight int
	Value int
}

type Backpack struct {
	MaxWeight int
	CurrentWeight int
	CurrentValue int
	Products []*Product
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
        fmt.Println(err)
        return 1
    }
	return weightInt
}

func getValueFromStr(line string) int {
	startIndex := strings.Index(line, ":")
	value := line[startIndex + 2 : ]
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error:", err)
		return 1
	}
	return valueInt
}

func printMenu() {
    fmt.Println("Меню")
    fmt.Println("1. Заполнение списка из файла")
	fmt.Println("2. Вывести список предметов")
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

func createProduct(name string, weight int, value int) *Product {
	return &Product{
		Name: name,
		Weight: weight,
		Value: value,
	}
}

func initializeProduct() *Product{
	var name string
	var weight int
	var value int
	fmt.Print("Введите название предмета: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Введите вес предмета: ")
	fmt.Scanf("%d", &weight)
	fmt.Print("Введите цену предмета: ")
	fmt.Scanf("%d", &value)
	return createProduct(name, weight, value)
}

func setProductName(p *Product) {
	var name string
	fmt.Print("Введите новое название предмета: ")
	fmt.Scanf("%s", &name)
	p.Name = name
}

func setProductWeight(p *Product) {
	var weight int
	fmt.Print("Введите новый вес предмета: ")
	fmt.Scanf("%d", &weight)
	p.Weight = weight
}

func setProductValue(p *Product) {
	var value int
	fmt.Print("Введите новую цену предмета: ")
	fmt.Scanf("%d", &value)
	p.Value = value
}

func changeProduct(products []*Product) {
	var ind int
	var ans int
	fmt.Print("Введите индекс предмета в списке: ")
	fmt.Scanf("%d", &ind)
	// fmt.Print("%d", len(products))
	for ind >= len(products) || ind < 0 {
		fmt.Print("Нет такого предмета. Введите индекс предмета в списке: ")
		fmt.Scanf("%d", &ind)
	}
	for {
		fmt.Println("Меню изменения предмета:")
		fmt.Println("1. Изменить название")
		fmt.Println("2. Изменить вес")
		fmt.Println("3. Изменить цену")
		fmt.Println("4. Вернуться в меню")
		fmt.Print("> ")
		fmt.Scanf("%d", &ans)
		switch ans{
		case 1:
			setProductName(products[ind])
		case 2:
			setProductWeight(products[ind])
		case 3:
			setProductValue(products[ind])
		case 4:
			return
		default:
			fmt.Println("Нет такой опции!")
		}
	}
}

func (b *Backpack) SetMaxWeight() {
	var weight int
	fmt.Print("Введите максимальный вес рюкзака: ")
	fmt.Scanf("%d", &weight)
	// fmt.Print("%d", len(products))
	for weight <= 0 {
		fmt.Print("Нет такого предмета. Введите индекс предмета в списке: ")
		fmt.Scanf("%d", &weight)
	}
	b.MaxWeight = weight
}

func showList(list []*Product) {
	for _, val := range list {
		fmt.Print(val)
	}
	fmt.Println()
}

func deleteProduct(list *[]*Product) {
	var ind int
	fmt.Print("Введите индекс предмета в списке для удаления: ")
	fmt.Scanf("%d", &ind)
	for ind >= len(*list) || ind < 0 {
		fmt.Print("Нет такого предмета. Введите индекс предмета в списке: ")
		fmt.Scanf("%d", &ind)
	}
	// list = append(list[:ind], list[ind+1:]...)
	copy((*list)[ind:], (*list)[ind + 1:])
	(*list)[len(*list) - 1] = nil
	*list = (*list)[:len(*list) - 1]
	fmt.Println("Предмет удален!")
}



func (b *Backpack) Add(product *Product) {
    if b.CurrentWeight + product.Weight <= b.MaxWeight {
        b.Products = append(b.Products, product)
        b.CurrentWeight += product.Weight
        b.CurrentValue += product.Value
    }
}

func main() {
	products := make([]*Product, 0)

	// Initialize a backpack as nil
	var backpack *Backpack
	ans := 0
	for {
		printMenu()
		fmt.Print("> ")
		fmt.Scanln(&ans)
		switch ans {
		case 1:
			path := getPath()
			products = addProductsFromFile(path)
		case 2:
			showList(products)
		case 3:
			products = append(products, initializeProduct())
		case 4:
			changeProduct(products)
		case 5:
			deleteProduct(&products)
		case 6:
			backpack.SetMaxWeight()
		case 7:
			if backpack == nil {
				fmt.Println("Рюкзак не инициализирован!")
				break
			}
			if len(products) == 0 {
				fmt.Println("Список предметов не инициализирован!")
				break
			}
			solver := KnapsackSolver{products: products, backpack: backpack}
			fmt.Println("1. Решение методом динамического программирования")
			fmt.Println("2. Решение жадным алгоритмом (макс. вес)")
			fmt.Println("3. Решение рекурсией")
			fmt.Print("> ")
			var localAns int
			fmt.Scanln(&localAns)
			switch localAns {
			case 1:
				backpack.clearContents()
				backpack = solver.findAnsDP(len(products), backpack.maxWeight)
			case 2:
				backpack.clearContents()
				backpack = solver.findAnsGreedy(len(products), backpack.maxWeight)
			case 3:
				backpack.clearContents()
				backpack = solver.findAnsRec(len(products), backpack.maxWeight)
			default:
				fmt.Println("Неверная команда!")
			}
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



/* 
	осталось реализовать выбор способа решения задачи
*/