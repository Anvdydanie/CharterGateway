package handlers

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"
)

const htmlFiles = "./frontend/html"
const frontFiles = "./frontend"

/*
Функция отдает запрошенный html файл
*/
func WebsiteHandler(w http.ResponseWriter, req *http.Request) {
	requestedFile := req.URL.Path
	if requestedFile == "/" {
		requestedFile = "index"
	}
	// находим все html файлы
	var templates = getTemlates()
	// отдаем шаблон при наличии или возвращаем 404
	view := templates.Lookup(requestedFile + ".html")
	if view != nil {
		view.Execute(w, nil)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

/*
Функция забирает все html файлы в папке
*/
func getTemlates() *template.Template {
	viewFolder, err := os.Open(htmlFiles)
	defer viewFolder.Close()
	if err == nil {
		var viewPaths []string
		viewPathRaw, _ := viewFolder.Readdir(-1)
		for _, pathInfo := range viewPathRaw {
			if !pathInfo.IsDir() {
				viewPaths = append(viewPaths, htmlFiles+"/"+pathInfo.Name())
			}
		}
		// Подключаем css и js
		//http.HandleFunc("/css/", serveResource)
		//http.HandleFunc("/js/", serveResource)

		result, _ := template.ParseFiles(viewPaths...)
		return result
	} else {
		return nil
	}
}

/*
Функция задает content-type в header для подключаемых скриптов
*/
func serveResource(w http.ResponseWriter, req *http.Request) {
	// папка, где лежат стили
	path := frontFiles + req.URL.Path
	// задаем content-type для файлов
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/plain"
	}
	// Добавляем заголовок
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
