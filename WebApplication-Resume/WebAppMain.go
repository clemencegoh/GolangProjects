package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"path/filepath"
	"io"
	"log"
	"sync"
	"strconv"
	"github.com/gorilla/mux"
)


type Page struct{
	Title string
	Body []byte
}


// following is for testing only, example webpage
func createTestWebpage() string{
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	return "TestPage"
}

func printWebpage(title string){
	p2, _ := loadPage(title)
	fmt.Println("Start of printing page:\n================")
	fmt.Println(string(p2.Body))
	fmt.Println("================")
}

func cleanupFiles(list []string){
	for _,item := range list{
		var err = os.Remove(item)
		if err != nil{
			fmt.Printf("---> Unsuccessful in removing %s\n", item)
		}
	}
	fmt.Println("Successfully removed all files")
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}


// lookForResume looks through all files for resumeID specified
func lookForResume(resumeID int) string{
	currentDir := getCurrentDir() + "\\WebApplication-Resume\\Resumes"
	filename := ""
	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error{
		resume,err := loadResume(path)
		if err != io.EOF{
			if resume.compareID(resumeID){
				filename = path
				return io.EOF
			}
		}
		return nil
	})
	if err != io.EOF{
		fmt.Println(err.Error())
	}
	if filename == ""{
		fmt.Println("File not found")
	}
	return filename
}

// loadResume loads ResumeDetails from a file
func loadResume(file string) (ResumeDetails, error){
	//fmt.Println(file)
	raw, err := ioutil.ReadFile(file)
	if len(raw) <= 0{
		return ResumeDetails{},io.EOF
	}
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var resumepage ResumeDetails
	json.Unmarshal(raw, &resumepage)

	return resumepage,nil
}

func loadAndDislayJson(file string, display bool) ResumeDetails{
	raw, err := ioutil.ReadFile(file)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var resumepage ResumeDetails
	json.Unmarshal(raw, &resumepage)
	bytes, err := json.Marshal(resumepage)

	if err!= nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if display{
		fmt.Println(string(bytes))
	}

	return resumepage
}

// returns error if exists
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/api/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}


func getExecutableDir() string{
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	return exPath
}

func getCurrentDir() string{
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(pwd)
	return pwd
}

// global
var counter = IDCounter{
	ID:0,
	mu:sync.Mutex{},
}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var rd ResumeDetails
		err := decoder.Decode(&rd)
		if err!=nil{
			panic(err)
		}
		defer r.Body.Close()
		log.Println(rd.Name + "Acquired")

		// determine id
		id := counter.GetAndIncrement()

		// save resume and id
		saveResumeDetails(rd,id)

		log.Println("Done saving")

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}


// saveResumeDetails returns err, id
// method call for all ResumeDetails Types
func saveResumeDetails(rd ResumeDetails, resumeID int) (error, int){
	filename := rd.Name + ".txt"
	fmt.Println(filename + " generated")

	rd.ResumeID = resumeID
	res, err := json.Marshal(rd)
	if err!=nil{
		fmt.Println("Error in processing resume")
		fmt.Println(err)
		return err, -1
	}

	// open file and save to file
	saveToPath := getCurrentDir() + "\\WebApplication-Resume\\Resumes\\" + filename
	//fmt.Println(saveToPath)
	f, err := os.OpenFile(saveToPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err,-1
	}
	n, err := f.Write(res)
	if err == nil && n < len(res) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err, resumeID
}

func GetHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		// parse params
		vars := mux.Vars(r)
		id := vars["id"]

		resID, err := strconv.ParseInt(id,10,64)
		if err!=nil{
			panic(err)
		}
		filename := lookForResume(int(resID))

		p := loadAndDislayJson(filename,false)

		fmt.Fprintf(w, "<h1>%s</h1><div>Name: %s\n" +
			"Current Job Description: %s\n" +
			"Current Job Title: %s\n" +
			"Current Job Company: %s</div>",
			p.Name, p.Name,
			p.CurrentJobDescription,
			p.CurrentJobTitle,
			p.CurrentJobCompany)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	r := mux.NewRouter()

	// endpoints here
	r.HandleFunc("/api/getResume/{id}",GetHandler)
	r.HandleFunc("/api/uploadResumeDetails",PostHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
