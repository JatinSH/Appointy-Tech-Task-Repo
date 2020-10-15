package main
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
)

var Articles []Article

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
     Articles = []Article{
        Article{Id: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello2", Subtitle: "Article Description2", Content: "Article Content2"},
    }
    handleRequests()
}




func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}


type Article struct {
    Id string `json:"Id"`
    Title string `json:"Title"`
    Subtitle string `json:"Subtitle"`
    Content string `json:"Content"`
    
}



func returnAllArticles(w http.ResponseWriter, r *http.Request){
   
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnSingleArticle")
    vars := mux.Vars(r)
    key := vars["id"]
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
       
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}