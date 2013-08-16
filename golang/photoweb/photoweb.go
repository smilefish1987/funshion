package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
	LIST_DIR     = 0x0001
)

var templates = make(map[string]*template.Template)

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileinfoArr, err := ioutil.ReadDir("./uploads")

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	check(err)

	/*
		var listHtml string
		for _, fileinfo := range fileinfoArr {
			imgid := fileinfo.Name()
			listHtml += "<li><a href=\"/view?id=" + imgid + "\">" + imgid + "</a></li>"
		}

		io.WriteString(w, "<html><ol>"+listHtml+"</ol></html>")
	*/

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileinfoArr {
		images = append(images, fileInfo.Name())
	}

	locals["images"] = images

	/*
		t, err := template.ParseFiles("list.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, locals)
	*/

	if err = renderHtml(w, "list", locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		/*
			io.WriteString(w, "<html><form method=\"POST\" action=\"/upload\" "+
				" enctype=\"multipart/form-data\">"+
				"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
				"<input type=\"submit\" value=\"Upload\" />"+
				"</form></html>")

			return
		*/

		/*
			t, err := template.ParseFiles("upload.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			t.Execute(w, nil)
			return
		*/

		if err := renderHtml(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	/*
		t,err := template.ParseFiles(tmpl+".html")
		if err!=nil{
			return
		}
		err = t.Execute(w,locals)
	*/
	err = templates[tmpl+".html"].Execute(w, locals)
	return err
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {

	/*
	 * 模板缓存(模板预加载)
	 */

	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template: ", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & LIST_DIR) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

/*
func safeHandler(fn http.HandleFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, error.Error(), http.StatusInternalServerError)
				log.Println("WAN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
*/

func main() {

	/*
		http.HandleFunc("/", listHandler)
		http.HandleFunc("/upload", uploadHandler)
		http.HandleFunc("/view", viewHandler)

		//http.HandleFunc("/", safeHandler(listHandler))
		//.HandleFunc("/upload", safeHandler(uploadHandler))
		//http.HandleFunc("/view", safeHandler(viewHandler))
	*/

	mux := http.NewServeMux()
	staticDirHandler(mux, "/asserts/", "./public", 0)
	mux.HandleFunc("/", listHandler)
	mux.HandleFunc("/view", viewHandler)
	mux.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", mux)
	//err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
