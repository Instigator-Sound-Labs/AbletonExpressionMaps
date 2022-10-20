package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Logic struct {
	XMLName xml.Name `xml:"plist"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Dict    struct {
		Text  string   `xml:",chardata"`
		Key   []string `xml:"key"`
		Array []struct {
			Text string `xml:",chardata"`
			Dict []struct {
				Text    string   `xml:",chardata"`
				Key     []string `xml:"key"`
				Integer []string `xml:"integer"`
				String  []string `xml:"string"`
				Dict    struct {
					Text    string   `xml:",chardata"`
					Key     []string `xml:"key"`
					Integer []string `xml:"integer"`
					String  string   `xml:"string"`
				} `xml:"dict"`
			} `xml:"dict"`
		} `xml:"array"`
		String  string `xml:"string"`
		Integer string `xml:"integer"`
	} `xml:"dict"`
}

type Ableton struct {
	Name  string `json:"name"`
	Key   int    `json:"key"`
	Key0  string `json:"+key"`
	Bnk   string `json:"bnk"`
	Sub   string `json:"sub"`
	Pgm   string `json:"pgm"`
	Ccn   string `json:"ccn"`
	Ccv   string `json:"ccv"`
	Chn   string `json:"chn"`
	Color string `json:"color"`
}

type AbletonStruct struct {
	Output interface{}
}

func main() {
	makeJson()
}

func getDir() {
	// get current Dir from the walk.
	path, _ := os.Getwd()
	fmt.Println(path)
}

func walkPath() {
}

func makeJson() {
	// xml is an io.Reader
	xmlFile, err := os.Open("/Users/johngoldsmith/code/AbletonExpressionMaps/SFA1 Brass High CB.plist")
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var articulation Logic
	xml.Unmarshal(byteValue, &articulation)
	if err != nil {
		panic("That's embarrassing...")
	}
	final := buildAbletonStruct(articulation)
	jsonOutput := []byte(final)
	_ = ioutil.WriteFile("outputtest.json", jsonOutput, 0644)
	defer xmlFile.Close()
}

func buildAbletonStruct(articulation Logic) string {
	var abletonStruct []string
	var getName []string
	numOfSwitches := len(articulation.Dict.Array[1].Dict)
	for i := 0; i < numOfSwitches; i++ {
		getName = articulation.Dict.Array[0].Dict[i].String
		keySwitch, _ := strconv.Atoi(articulation.Dict.Array[1].Dict[i].Integer[1])
		output := convertToJson(keySwitch, getName[0], i)
		store := map[string]interface{}{strconv.Itoa(i + 1): output}
		values, _ := json.Marshal(store)
		clipEnd := strings.TrimSuffix(string(values), "}")
		clipBeginning := strings.TrimPrefix(clipEnd, "{")
		toAppend := clipBeginning
		abletonStruct = append(abletonStruct, toAppend)
	}
	output := strings.Join(abletonStruct, ",")
	first := `{"KSEM-Version": 4.1, "ks":{` + output + "}}"
	return first
}

func convertToJson(keySwitch int, getName string, i int) Ableton {
	newDataStruct := Ableton{}
	newDataStruct.Name = getName
	newDataStruct.Key = keySwitch
	newDataStruct.Key0 = "-"
	newDataStruct.Bnk = "-"
	newDataStruct.Sub = "-"
	newDataStruct.Pgm = "-"
	newDataStruct.Ccn = "-"
	newDataStruct.Ccv = "-"
	newDataStruct.Chn = "-"
	newDataStruct.Color = "-"
	return newDataStruct
}

func makeFile() {
	// make file for each file traveresd
	// make the json file.
}
