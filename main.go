package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	"encoding/json"
	"fmt"
	"./models"
	"github.com/Sirupsen/logrus"
	"time"
)

var last int

// Считывает конфигурационные данные из файла
func readConfigFile() (models.Config, error) {

	configFileName := "tomato.cfg"
	var config models.Config

	// Читать файл в JSON-строку
	file, err := os.Open(configFileName)
	defer file.Close()
	if err == nil {
		// Получить размер файла
		stat, err := file.Stat()
		if err == nil {
			// Буфер
			buffer := make([]byte, stat.Size())
			_, err = file.Read(buffer)
			if err == nil {
				// Десериализовать
				err = json.Unmarshal(buffer, &config)
			}
		}
	}

	return config, err
}


func main() {

	/////////////
	// Запуск GTK
	gtk.Init(&os.Args)

	/////////////
	// Приложение

	// Прочитать конфиг
	config, err := readConfigFile()
	if err == nil {
		logrus.Infof("Конфигурационные параметры: %v", config)

		// Рабочий ли день?
		var workDay bool = false
		workDay = WorkingDayWaiting(workDay)
		if workDay {
			fmt.Println("Наступил рабочий день")
		}

		/// Запустить цикл
		// Сравнить текущее время с временем начала - если совпало, то запустить отсчёт

		// Новый виджет - окно
		window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)

		// Позиция окна
		window.SetPosition(gtk.WIN_POS_CENTER)

		// Незаполняемая виджетами граница от краёв окна
		window.SetBorderWidth(20)

		// Заголовок окна
		window.SetTitle("Помидор")

		// Размер окна
		window.SetSizeRequest(700, 500)

		vbox := gtk.NewVBox(false, 1)
		window.Add(vbox)

		label := gtk.NewLabel("Это помидор")
		vbox.Add(label)

		hsep := gtk.NewHSeparator()
		vbox.Add(hsep)

		button := gtk.NewButtonWithLabel("Жми!")
		vbox.Add(button)

		// Отобразить окно
		window.ShowAll()

		// При закрытии окна выйти безопасно из приложения
		window.Connect("destroy", func() {
			gtk.MainQuit()
		})

		//////////////////////////
		// Передать управление GTK
		gtk.Main()

	} else {
		fmt.Printf("Ошибка при чтении файла конфигурации: %v", err)
	}
}

// Ожидание наступления рабочего дня недели
func WorkingDayWaiting(workDay bool) bool {
	for workDay != true {
		workDay = IsWorkingDay()
		logrus.Infof("День недели: '%v'", workDay)
		time.Sleep(time.Duration(time.Hour)) // Через 1 час снова проверить
	}
	return workDay
}

// Определить рабочий ли день в настоящий момент
func IsWorkingDay() bool {
	weekday := time.Now().Weekday().String()
	//logrus.Infof("День недели: %v", weekday)
	workdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	var status bool = false
	for _, day := range workdays {
		if weekday == day {
			status = true
			break
		}
	}
	return status
}

