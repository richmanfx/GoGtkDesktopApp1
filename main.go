package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	"encoding/json"
	"fmt"
	"./models"
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
		fmt.Printf("Конфигурационные параметры: %v", config)

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

