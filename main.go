package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	deleteIndex()
	createIndex()
	createMapping()
}

func deleteIndex() {
	fmt.Println("Deleting index...")
	url := "http://localhost:9200/rustest"

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Define request error: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Make request error: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read resp body err: %v", err)
	}
	fmt.Println("response Body:", string(body))
}

func createIndex() {
	fmt.Println("Creating index with settings...")
	url := "http://localhost:9200/rustest"

	jsonStr := `{
        "settings": {
            "analysis": {
                "analyzer": {
                    "ru": {
                        "type": "custom",
                        "tokenizer": "standard",
                        "filter": ["lowercase", "russian_morphology", "english_morphology", "my_stopwords"]
                    }
                },
                "filter": {
                    "my_stopwords": {
                        "type": "stop",
                        "stopwords": "а,без,более,бы,был,была,были,было,быть,в,вам,вас,весь,во,вот,все,всего,всех,вы,где,да,даже,для,до,его,ее,если,есть,еще,же,за,здесь,и,из,или,им,их,к,как,ко,когда,кто,ли,либо,мне,может,мы,на,надо,наш,не,него,нее,нет,ни,них,но,ну,о,об,однако,он,она,они,оно,от,очень,по,под,при,с,со,так,также,такой,там,те,тем,то,того,тоже,той,только,том,ты,у,уже,хотя,чего,чей,чем,что,чтобы,чье,чья,эта,эти,это,я,a,an,and,are,as,at,be,but,by,for,if,in,into,is,it,no,not,of,on,or,such,that,the,their,then,there,these,they,this,to,was,will,with"
                    }
                }
            }
        }
    }`

	var jsonBytes = []byte(jsonStr)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatalf("Define request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Make request error: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read resp body err: %v", err)
	}
	fmt.Println("response Body:", string(body))
}

func createMapping() {
	fmt.Println("Creating mapping for types...")
	url := "http://localhost:9200/rustest/type1/_mapping"

	jsonStr := `{
        "type1": {
            "_all" : {"analyzer" : "russian_morphology"},
            "properties" : {
                "body" : { "type" : "string", "analyzer" : "ru" }
            }
        }
    }`

	var jsonBytes = []byte(jsonStr)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatalf("Define request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Make request error: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read resp body err: %v", err)
	}
	fmt.Println("response Body:", string(body))
}
