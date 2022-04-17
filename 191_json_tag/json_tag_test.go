package _91_json_tag

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Max Book Pro",
		Price:    0,
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Max Book Pro","price":0,"image_url":"http://example.com/image.png"}`
	jsonByte := []byte(jsonString)

	product := Product{}
	json.Unmarshal(jsonByte, &product)

	fmt.Println(product)
}
