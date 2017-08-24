package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
)

var last int

func main() {

	/////////////
	// Запуск GTK
	gtk.Init(&os.Args)

	/////////////
	// Приложение

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
}

func addPage(notebook *gtk.Notebook) {
	dialog := gtk.NewDialog()
	dialog.SetTitle("Title?")
	dVbox := dialog.GetVBox()

	input := gtk.NewEntry()
	input.SetEditable(true)
	dVbox.Add(input)
	vbox := gtk.NewVBox(false, 1)


	input.Connect("activate", func() {
		s := input.GetText()
		if s != "" {
			notebook.InsertPage(vbox, gtk.NewLabel(s), last)
			last++
			notebook.ShowAll()
		}
		notebook.PrevPage()
		dialog.Destroy()
	})

	button := gtk.NewButtonWithLabel("OK")
	button.Connect("clicked", func() {
		input.Emit("activate")
	})
	dVbox.Add(button)
	dialog.SetModal(true)
	dialog.ShowAll()

}