package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Student представляет информацию о студенте и его оценках
type Student struct {
	ID     int
	Name   string
	Group  string
	Grades []int // оценки по 100-балльной шкале
}

// StudentGradeManager управляет списком студентов
type StudentGradeManager struct {
	students []Student
	nextID   int
}

func NewStudentGradeManager() *StudentGradeManager {
	return &StudentGradeManager{
		students: []Student{
			{ID: 1, Name: "Иван Петров", Group: "CS-301", Grades: []int{85, 90, 78, 92}},
			{ID: 2, Name: "Анна Сидорова", Group: "CS-302", Grades: []int{95, 98, 91, 100}},
			{ID: 3, Name: "Дмитрий Козлов", Group: "CS-301", Grades: []int{45, 52, 48, 60}},
			{ID: 4, Name: "Елена Смирнова", Group: "SE-303", Grades: []int{88, 76, 82, 85}},
			{ID: 5, Name: "Арман Сериков", Group: "CS-302", Grades: []int{40, 35, 48, 55}},
		},
		nextID: 6,
	}
}

// 1. AddStudent добавляет нового студента
func (m *StudentGradeManager) AddStudent(name, group string) {
	s := Student{
		ID:     m.nextID,
		Name:   strings.TrimSpace(name),
		Group:  strings.TrimSpace(group),
		Grades: []int{},
	}
	m.students = append(m.students, s)
	m.nextID++
	fmt.Printf("Студент #%d (%s, группа %s) успешно добавлен!\n", s.ID, s.Name, s.Group)
}

// 2. AddGrade добавляет оценку студенту по ID через append()
func (m *StudentGradeManager) AddGrade(id int, grade int) bool {
	if grade < 0 || grade > 100 {
		return false
	}
	for i := range m.students {
		if m.students[i].ID == id {
			m.students[i].Grades = append(m.students[i].Grades, grade)
			return true
		}
	}
	return false
}

// 3. PrintAllStudents выводит всех студентов со средним баллом
func (m *StudentGradeManager) PrintAllStudents() {
	if len(m.students) == 0 {
		fmt.Println("Список студентов пуст.")
		return
	}
	fmt.Println("\n--- СПИСОК СТУДЕНТОВ ---")
	fmt.Printf("%-4s | %-20s | %-8s | %-16s | %s\n", "ID", "ФИО", "Группа", "Оценки", "Средний балл")
	fmt.Println(strings.Repeat("-", 65))
	for _, s := range m.students {
		avg := calculateAverage(s.Grades)
		gradesStr := formatGrades(s.Grades)
		fmt.Printf("#%-3d | %-20s | %-8s | %-16s | %.2f\n", s.ID, s.Name, s.Group, gradesStr, avg)
	}
	fmt.Println(strings.Repeat("-", 65))
}

// 4. FindTopStudent находит студента с максимальным средним баллом
func (m *StudentGradeManager) FindTopStudent() *Student {
	var top *Student
	maxAvg := -1.0
	for i := range m.students {
		avg := calculateAverage(m.students[i].Grades)
		if avg > maxAvg {
			maxAvg = avg
			top = &m.students[i]
		}
	}
	return top
}

// 5. CalculateOverallAverage вычисляет общий средний балл по всей группе/базе
func (m *StudentGradeManager) CalculateOverallAverage() float64 {
	totalGrades := 0
	count := 0
	for _, s := range m.students {
		for _, g := range s.Grades {
			totalGrades += g
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(totalGrades) / float64(count)
}

// 6. FindUnderperforming находит студентов с академической задолженностью (средний балл < 50)
func (m *StudentGradeManager) FindUnderperforming() []Student {
	var list []Student
	for _, s := range m.students {
		if len(s.Grades) > 0 && calculateAverage(s.Grades) < 50 {
			list = append(list, s)
		}
	}
	return list
}

// 7. GroupAverageGrade группирует средний балл по группам через map[string]float64
func (m *StudentGradeManager) GroupAverageGrade() map[string]float64 {
	groupSums := make(map[string]int)
	groupCounts := make(map[string]int)

	for _, s := range m.students {
		for _, g := range s.Grades {
			groupSums[s.Group] += g
			groupCounts[s.Group]++
		}
	}

	result := make(map[string]float64)
	for grp, sum := range groupSums {
		if groupCounts[grp] > 0 {
			result[grp] = float64(sum) / float64(groupCounts[grp])
		}
	}
	return result
}

func calculateAverage(grades []int) float64 {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for _, g := range grades {
		sum += g
	}
	return float64(sum) / float64(len(grades))
}

func formatGrades(grades []int) string {
	if len(grades) == 0 {
		return "—"
	}
	var strList []string
	for _, g := range grades {
		strList = append(strList, strconv.Itoa(g))
	}
	return strings.Join(strList, ", ")
}

func main() {
	manager := NewStudentGradeManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== STUDENT GRADE MANAGER =====")
		fmt.Println("1. Добавить студента")
		fmt.Println("2. Добавить оценку студенту")
		fmt.Println("3. Список всех студентов")
		fmt.Println("4. Лучший студент")
		fmt.Println("5. Общий средний балл")
		fmt.Println("6. Студенты с задолженностями (< 50 баллов)")
		fmt.Println("7. Средний балл по группам")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("ФИО студента: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Группа: ")
			scanner.Scan()
			group := scanner.Text()

			manager.AddStudent(name, group)

		case "2":
			fmt.Print("ID студента: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Ошибка: некорректный ID")
				continue
			}

			fmt.Print("Оценка (0-100): ")
			scanner.Scan()
			grade, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil || grade < 0 || grade > 100 {
				fmt.Println("Ошибка: оценка должна быть в диапазоне от 0 до 100!")
				continue
			}

			if manager.AddGrade(id, grade) {
				fmt.Printf("Оценка %d успешно добавлена студенту #%d!\n", grade, id)
			} else {
				fmt.Println("Студент не найден!")
			}

		case "3":
			manager.PrintAllStudents()

		case "4":
			top := manager.FindTopStudent()
			if top == nil {
				fmt.Println("Студентов нет.")
			} else {
				fmt.Printf("\nЛучший студент: %s (группа %s) — средний балл: %.2f\n",
					top.Name, top.Group, calculateAverage(top.Grades))
			}

		case "5":
			overall := manager.CalculateOverallAverage()
			fmt.Printf("\nОбщий средний балл всех студентов: %.2f\n", overall)

		case "6":
			under := manager.FindUnderperforming()
			fmt.Printf("\nСтуденты с академической задолженностью (всего %d):\n", len(under))
			if len(under) == 0 {
				fmt.Println("Задолженностей нет! Все студенты успевают.")
			} else {
				for _, s := range under {
					fmt.Printf("- #%d %s (%s) — ср. балл: %.2f (Оценки: %s)\n",
						s.ID, s.Name, s.Group, calculateAverage(s.Grades), formatGrades(s.Grades))
				}
			}

		case "7":
			byGrp := manager.GroupAverageGrade()
			fmt.Println("\n--- СРЕДНИЙ БАЛЛ ПО ГРУППАМ ---")
			for grp, avg := range byGrp {
				fmt.Printf("Группа %-8s: %.2f\n", grp, avg)
			}

		case "0":
			fmt.Println("Работа программы Student Grade Manager завершена.")
			return

		default:
			fmt.Println("Неверный пункт меню. Введите от 0 до 7.")
		}
	}
}
