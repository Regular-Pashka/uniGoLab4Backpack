package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
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

func (b *Backpack) Clear() {
	b.CurrentWeight = 0
	b.CurrentValue = 0
	b.Products = nil
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
	// fmt.Print(weight + "WEIGHT")
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
	// fmt.Print(value + "VALUE")
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
	fmt.Println("USPEH")
	for weight <= 0 {
		fmt.Print("Максимальная грузоподъемность рюкзака не может быть отрицательной!: ")
		fmt.Scanf("%d", &weight)
	}
	fmt.Println("USPEH2")
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


func solveDynamic(products []*Product, maxWeight int) (int, []*Product, int) {
	n := len(products)
	dp := make([]int, maxWeight + 1)
	selected := make([][]bool, n + 1)
	currWeight := 0

	for i := 0; i <= n; i++ {
		selected[i] = make([]bool, maxWeight+1)
	}

	for i := 1; i <= n; i++ {
		for w := maxWeight; w >= products[i - 1].Weight; w-- {
			if dp[w] < dp[w - products[i - 1].Weight] + products[i - 1].Value {
				dp[w] = dp[w - products[i - 1].Weight] + products[i - 1].Value
				selected[i][w] = true
			}
		}
	}

	// Определяем, какие предметы были выбраны
	var chosenProducts []*Product
	w := maxWeight
	for i := n; i > 0; i-- {
		if selected[i][w] {
			chosenProducts = append(chosenProducts, products[i-1])
			w -= products[i-1].Weight
		}
	}

	
	for _, val := range chosenProducts {
		currWeight += val.Weight
	}
	return dp[maxWeight], chosenProducts, currWeight
}

func InitializeBackpack() *Backpack {
	return &Backpack{
		MaxWeight: 0,
		CurrentWeight: 0,
		CurrentValue: 0,
		Products: make([]*Product, 0),
	}
}

func (backpack Backpack) Output() {
	fmt.Printf("Максимальная ценность: %d\n", backpack.CurrentValue)
	fmt.Printf("Выбранные предметы:\n")
	for _, product := range backpack.Products {
		fmt.Printf("Название: %s, Вес: %d, Ценность: %d\n", product.Name, product.Weight, product.Value)
	}
	fmt.Printf("Общий вес: %d\n", backpack.CurrentWeight)
}

func (b *Backpack) solveGreedy(products []*Product) {
	// заполняем в первую очередь предметами с максимальным весом
	sort.Slice(products, func(i, j int) bool {
		return products[i].Weight > products[j].Weight
	})
	for _, product := range products {
		if b.CurrentWeight + product.Weight <= b.MaxWeight {
			b.Products = append(b.Products, product)
			b.CurrentWeight += product.Weight
			b.CurrentValue += product.Value
		}
	}
}

func solveRecursive(products []*Product, currentIndex, maxWeight int, selectedProducts []*Product) (int, []*Product) {
	// Базовый случай: если вес рюкзака равен 0 или нет больше предметов для рассмотрения
	if maxWeight == 0 || currentIndex < 0 {
		return 0, selectedProducts
	}

	// Если вес текущего предмета больше максимального веса рюкзака
	if products[currentIndex].Weight > maxWeight {
		return solveRecursive(products, currentIndex-1, maxWeight, selectedProducts)
	}

	// Включить текущий предмет
	includeValue, includeProducts := solveRecursive(products, currentIndex-1, maxWeight-products[currentIndex].Weight, append(selectedProducts, products[currentIndex]))

	// Исключить текущий предмет
	excludeValue, excludeProducts := solveRecursive(products, currentIndex-1, maxWeight, selectedProducts)

	// Сравниваем значения и возвращаем максимальное
	if includeValue+products[currentIndex].Value > excludeValue {
		return includeValue + products[currentIndex].Value, includeProducts
	} else {
		return excludeValue, excludeProducts
	}
}

// Вспомогательная функция для получения максимума
func max(a, b int) int {
	result := b
	if a > b {
		result = a
	}
	return result
}

func main() {
	products := make([]*Product, 0)

	// Initialize a backpack as nil
	backpack := InitializeBackpack() 
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
			if len(products) == 0 {
				fmt.Println("Список предметов не инициализирован!")
				break
			}
			fmt.Println("1. Решение методом динамического программирования")
			fmt.Println("2. Решение жадным алгоритмом (макс. вес)")
			fmt.Println("3. Решение рекурсией")
			fmt.Print("> ")
			var localAns int
			fmt.Scanln(&localAns)
			switch localAns {
			case 1:
				backpack.Clear()
				backpack.CurrentValue, backpack.Products, backpack.CurrentWeight = solveDynamic(products, backpack.MaxWeight)
				backpack.Output()
			case 2:
				backpack.Clear()
				backpack.solveGreedy(products)
				backpack.Output()
			case 3:
				backpack.Clear()
				_, selectedProducts := solveRecursive(products, len(products) - 1, backpack.MaxWeight, []*Product{})
				for _, product := range selectedProducts {
					backpack.Products = append(backpack.Products, product)
					backpack.CurrentWeight += product.Weight
					backpack.CurrentValue += product.Value
				}
				backpack.Output()
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