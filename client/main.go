package main
import (
	"github.com/gotk3/gotk3/gtk"
	// "github.com/go-kit/kit/log"
	// "github.com/go-kit/kit/log/level"

	// "encoding/json"
	// "fmt"
	// "net/http"
	// "os"
	// "path/filepath"
	// "time"
)
func main() {
	gtk.Init(nil)
	b, _ := gtk.BuilderNew()
	b.AddFromFile("MainWindow.glade")
	obj, _ := b.GetObject("MainWindow")
	win := obj.(*gtk.Window)
	win.Connect("destroy", func(){
		gtk.MainQuit()
	})
	win.ShowAll()
	gtk.Main()
}