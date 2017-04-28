package controller

import(
	"net/http"
	"html/template"
	"log"
	"path/filepath"
	"runtime"
	"path"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Host)
	// 设置返回状态码
	rw.WriteHeader(http.StatusOK)
	// 读取主页模板
	// 返回对应调用文件的程序指针，文件名和行号信息
	// The return values report the
	// program counter, file name, and line number within the file of the corresponding
	// call.
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No runtime.Caller information")
	}
	currentDir := path.Dir(filename)
	indexViewFile := filepath.Join(currentDir, "../view/home/index.html")
	tpl := template.Must(template.ParseFiles(indexViewFile))

	// 将模板写入响应
	err := tpl.Execute(rw, nil)

	if err != nil {
		log.Fatalf("template execution failed: %s", err)
	}
}

func GetName(rw http.ResponseWriter, r *http.Request) {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	//resp,_ := json.Marshal(str)
	log.Println(r.Cookies())
	// normal header
	//rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Write([]byte(str))
}