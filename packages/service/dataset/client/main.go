package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Dataset struct {
	DataID int    `json:"dataId"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Meta   []Meta `json:"meta"`
}

type Meta struct {
	Fid          string `json:"fid"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	SemanticType string `json:"semanticType"`
	GeoRole      string `json:"geoRole"`
	DataType     string `json:"dataType"`
	Order        int64  `json:"order"`
}

// QueryMeta
func QueryhMeta(dataId int) ([]Meta, error) {
	resp, err := http.Get("http://localhost:23402/query/" + strconv.Itoa(dataId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var meta []Meta
	err = json.Unmarshal(body, &meta)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func main() {
	metas, err := QueryhMeta(1)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(metas)
}
