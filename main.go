package main

import(
	"log"
	"os"
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Instance struct{
	gorm.Model
	Timestamp string
	User string
	Hostname string
	Hash string
	Ps1 string
}
func (i *Instance) check() error {
if i.Hash ==""||
i.Timestamp=="" ||
i.User =="" ||
i.Hostname =="" ||
i.Ps1==""{return errors.New("invalid instance")}
return nil


}

var db *gorm.DB

func getInstances(w http.ResponseWriter, r *http.Request){
	var instances []Instance
	result := db.Find(&instances)
	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		log.Panicf("unable to load data from db")
	}
	response, err := json.Marshal(instances)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panicf("unable to marshal response")
	}

	w.Write(response)

}

func addInstance(w http.ResponseWriter, r *http.Request){

	var instance Instance
	// check for post request
	if r.Method != http.MethodPost {
		http.Error(w,"no POST request",http.StatusBadRequest)
		log.Panicf("unable to decode data from request")
	}

	err := json.NewDecoder(r.Body).Decode(&instance)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Panicf("unable to decode data from request")
	}
	if err = instance.check(); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Panicf("bad json object")
	}

	result := db.Create(&instance)
	if result.Error != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panicf("unable to create instance on db")
	}

	fmt.Fprintf(w, "instance added")
}

func removeInstance(w http.ResponseWriter, r *http.Request){
	var instance Instance
	result := db.Where("hash = ?", r.PathValue("hash")).Delete(&instance)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		log.Panicf("unable to remove instance from db")

	}
	fmt.Fprintf(w, "instance removed")
}

func main(){

	// LOADING ENV VARS
	log.Println("loading vars")
	path := os.Getenv("TB_HAWKEYE_DB_PATH")
	port := os.Getenv("TB_HAWKEYE_PORT")
	address := os.Getenv("TB_HAWKEYE_ADDRESS")

	// OPEN DB
	var err error
	db ,err= gorm.Open(sqlite.Open(path+"instances.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to open database at path %s",path)
	}
	db.AutoMigrate(&Instance{})

	//start http endpoint
	http.HandleFunc("/instances", getInstances)
	http.HandleFunc("/addinstance", addInstance)
	http.HandleFunc("/removeinstance/{hash}", removeInstance)
	log.Printf("start listening on %s:%s",address,port)
	log.Fatal(http.ListenAndServe(address + ":" + port, nil))

}
