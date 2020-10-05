package controllers

import "fmt"
import "net/http"
import "encoding/json"
import "time"
import "html/template"
import "MVCDatetime/app/models"

func GetData(w http.ResponseWriter, r *http.Request){
	var date = models.Datetime{
		Now : time.Now().Format("2006-01-02T15:04:05Z"),
		After : time.Now().Add(12*time.Hour).Format("2006-01-02T15:04:05Z"),
	}

	data := map[string]string{
		"Time" : date.Now,
		"TimeAfter" : date.After,
	}
	var t, err = template.ParseFiles("views/view.html")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, data)
}

func GetAfter(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		decoder := json.NewDecoder(r.Body)

		payload := struct{
			Times string `json:"times"`
		}{}
		if err := decoder.Decode(&payload);
		err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		message := fmt.Sprintf(
			`[{"Time" : %s}, {"TimeAfter" : %s}]`,
			payload.Times,
			time.Now().Add(12*time.Hour).Format("2006-01-02T15:04:05Z"),
)
w.Write([]byte(message))
return
}
http.Error(w, "Only accept POST request", http.StatusBadRequest)
}