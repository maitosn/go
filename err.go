package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Интерфейс для всех животных
type Animal interface {
	Speak() (string, error)
	Move() (string, error)
	Eat() (string, error)
	Sleep() (string, error)
}

// Интерфейс для животных, которые умеют плавать
type Swimmer interface {
	CanSwim() bool
}

// Структура для неизвестного животного
type UnknownAnimal struct{}

func (u UnknownAnimal) Speak() (string, error) {
	return "", errors.New("неизвестное животное не может издавать звуки")
}

func (u UnknownAnimal) Move() (string, error) {
	return "", errors.New("неизвестное животное не может двигаться")
}

func (u UnknownAnimal) Eat() (string, error) {
	return "", errors.New("неизвестное животное не может есть")
}

func (u UnknownAnimal) Sleep() (string, error) {
	return "", errors.New("неизвестное животное не может спать")
}

type Monkey struct{}

func (t Monkey) Speak() (string, error) {
	return "Кричит", nil
}
func (t Monkey) Move() (string, error) {
	return "Прыгает по деревьям", nil
}
func (t Monkey) Eat() (string, error) {
	return "Ест бананы", nil
}
func (t Monkey) Sleep() (string, error) {
	return "Спит в вольере", nil
}
func (t Monkey) CanSwim() bool {
	return false
}
func (m Monkey) Climb() string {
	return "Лазает по деревьям"
}

type Shark struct{}

func (t Shark) Speak() (string, error) {
	return "не издаёт звук", nil
}
func (t Shark) Move() (string, error) {
	return "Плавает", nil
}
func (t Shark) Sleep() (string, error) {
	return "Спит на дне", nil
}
func (t Shark) Eat() (string, error) {
	return "Съела мясо", nil
}
func (t Shark) CanSwim() bool {
	return true
}
func (s Shark) Hunt() string {
	return "Охотится на рыбу"
}

type Eagle struct{}

func (t Eagle) Speak() (string, error) {
	return "Орёт", nil
}
func (t Eagle) Move() (string, error) {
	return "Летит", nil
}
func (t Eagle) Eat() (string, error) {
	return "Ест мышь", nil
}
func (t Eagle) Sleep() (string, error) {
	return "Спит в гнезде", nil
}
func (t Eagle) CanSwim() bool {
	return false
}
func (e Eagle) Fly() string {
	return "Летает высоко"
}

type Bear struct{}

func (t Bear) Speak() (string, error) {
	return "Рычит", nil
}
func (t Bear) Move() (string, error) {
	return "Бежит", nil
}
func (t Bear) Eat() (string, error) {
	return "Ест малину", nil
}
func (t Bear) Sleep() (string, error) {
	return "Спит в берлоге", nil
}
func (t Bear) CanSwim() bool {
	return true
}
func (b Bear) Hibernate() string {
	return "Спит зимой"
}

type Whale struct{}

func (t Whale) Speak() (string, error) {
	return "Издаёт ултразвук", nil
}
func (t Whale) Move() (string, error) {
	return "Медленно плывёт у поверхности", nil
}
func (t Whale) Eat() (string, error) {
	return "Ест планктон", nil
}
func (t Whale) Sleep() (string, error) {
	return "Спит вертикально", nil
}
func (t Whale) CanSwim() bool {
	return true
}
func (w Whale) Dive() string {
	return "Ныряет глубоко"
}

// Функция для записи ошибок в лог
func logError(err error) {
	file, _ := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	log := fmt.Sprintf("Ошибка: %s\n", err)
	file.WriteString(log)
}

func main() {
	animals := map[string]Animal{
		"monkey": Monkey{},
		"shark":  Shark{},
		"eagle":  Eagle{},
		"bear":   Bear{},
		"whale":  Whale{},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Выберите животное (monkey, shark, eagle, bear, whale):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	animal, ok := animals[input]
	if !ok {
		animal = UnknownAnimal{}
	}

	fmt.Printf("Животное: %T\n", animal)

	// Обработка ошибок при вызове методов Speak, Move, Eat, Sleep
	if sound, err := animal.Speak(); err != nil {
		fmt.Printf("Ошибка при вызове Speak: %s\n", err)
		logError(err)
	} else {
		fmt.Printf("Звук: %s\n", sound)
	}

	if movement, err := animal.Move(); err != nil {
		fmt.Printf("Ошибка при вызове Move: %s\n", err)
		logError(err)
	} else {
		fmt.Printf("Движение: %s\n", movement)
	}

	if food, err := animal.Eat(); err != nil {
		fmt.Printf("Ошибка при вызове Eat: %s\n", err)
		logError(err)
	} else {
		fmt.Printf("Еда: %s\n", food)
	}

	if sleep, err := animal.Sleep(); err != nil {
		fmt.Printf("Ошибка при вызове Sleep: %s\n", err)
		logError(err)
	} else {
		fmt.Printf("Сон: %s\n", sleep)
	}

	// Проверка, умеет ли животное плавать
	if swimmer, ok := animal.(Swimmer); ok {
		fmt.Printf("Умеет плавать: %v\n", swimmer.CanSwim())
	}

	// Вызов специфичных методов для каждого типа животного
	switch a := animal.(type) {
	case Monkey:
		fmt.Printf("Лазание: %s\n", a.Climb())
	case Shark:
		fmt.Printf("Охота: %s\n", a.Hunt())
	case Eagle:
		fmt.Printf("Полет: %s\n", a.Fly())
	case Bear:
		fmt.Printf("Зимовка: %s\n", a.Hibernate())
	case Whale:
		fmt.Printf("Ныряние: %s\n", a.Dive())
	}

	fmt.Println("-----------------------------")
}
