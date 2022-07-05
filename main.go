package main

import (
	"fmt"
)

// собака - їсть 10кг корму на кожні 5кг власної вагиs
const FeedingFactorDog = 2

// кішка  - 7кг на кожен кілограм власної ваги
const FeedingFactorCat = 7

// корова - 25кг на кожен кілограм власної ваги
const FeedingFactorCow = 25

const (
	NameDog = "Dog"
	NameCat = "Cat"
	NameCow = "Cow"
)

type AnimalInterface interface {
	Name() string
	Weight() float64
	FoodForMonth() float64
	IsValid() bool
}

type Animal struct {
	name          string
	weight        float64
	feedingFactor float64
}

func (a Animal) Name() string {
	return a.name
}
func (a Animal) Weight() float64 {
	return a.weight
}
func (a Animal) FoodForMonth() float64 {
	return a.weight * a.feedingFactor
}
func (a Animal) IsValid() bool {
	return a.weight > 0
}

// Dog
type Dog struct {
	Animal
}

func NewDog(weight float64) Dog {
	item := Dog{}
	item.name = NameDog
	item.weight = weight
	item.feedingFactor = FeedingFactorDog

	return item
}

// Cat
type Cat struct {
	Animal
}

func NewCat(weight float64) Cat {
	item := Cat{}
	item.name = NameCat
	item.weight = weight
	item.feedingFactor = FeedingFactorCat

	return item
}

// Cow
type Cow struct {
	Animal
}

func NewCow(weight float64) Cow {
	item := Cow{}
	item.name = NameCow
	item.weight = weight
	item.feedingFactor = FeedingFactorCow

	return item
}

// Farm
type Farm struct {
	animals []AnimalInterface
}

func (f *Farm) Add(animal AnimalInterface) {
	if !animal.IsValid() {
		return
	}

	f.animals = append(f.animals, animal)
}

func (f *Farm) AllAnimals() []AnimalInterface {
	return f.animals
}

// написати динамічну функцію, яка буде ВИВОДИТИ в консоль
// для кожної тварини на фермі
//  * її назву,
//  * вагу,
//  * та скільки їжі треба для того щоб її прогодувати
// ця функція також моє ПОВЕРТАТИ сумму кг корму на всю ферму
// для подальшого виводу у консоль
func FarmReport(farm Farm) float64 {
	var allFood float64
	for _, animal := range farm.AllAnimals() {
		fmt.Println(fmt.Sprintf("name: %s weight: %.2f food per month: %.2f", animal.Name(), animal.Weight(), animal.FoodForMonth()))
		allFood += animal.FoodForMonth()
	}

	return allFood
}

func main() {
	farm := Farm{}
	animals := []AnimalInterface{
		NewDog(1),
		NewCat(1),
		NewCow(1),

		NewDog(10.5),
		NewCat(5.2),
		NewCow(255.4),

		NewDog(-15.0),
		NewCat(0),
		NewCow(-300.4),

		NewDog(2.6),
		NewCat(1.5),
		NewCow(45.2),

		NewDog(23.4),
		NewCat(6.3),
		NewCow(230.9),
	}

	for _, v := range animals {
		farm.Add(v)
	}

	fmt.Println("Food for all animals per month:", FarmReport(farm))
}
