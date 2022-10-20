package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

func main() {
	makeJson()
}

// func getDir() {
// 	// get current Dir from the walk.
// 	path, _ := os.Getwd()
// 	fmt.Println(path)
// }

// func walkPath() { // maybe move to main?
// 	path, err := os.Getwd()
// 	if err != nil {
// 		fmt.Printf("cannot get current dir: %v\n", err)
// 		return
// 	}
// 	os.Chdir(path)

// 	subDirToSkip := "skip"

// 	fmt.Println("On Unix:")
// 	// ensure the files are only plist files.
// 	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
// 		if err != nil {
// 			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
// 			return err
// 		}
// 		if info.IsDir() && info.Name() == subDirToSkip {
// 			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
// 			return filepath.SkipDir
// 		}
// 		fmt.Printf("visited file or dir: %q\n", path)
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Printf("error walking the path %q: %v\n", path, err)
// 		return
// 	}
// }

func makeJson() {
	// xml is an io.Reader
	xmlFile, err := os.Open("/Users/johngoldsmith/code/AbletonExpressionMaps/NISE String Ensemble.plist")
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var articulation Logic
	xml.Unmarshal(byteValue, &articulation)
	if err != nil {
		panic("That's embarrassing...")
	}
	var getName []string
	numOfSwitches := len(articulation.Dict.Array[1].Dict)
	for i := 0; i < numOfSwitches; i++ {
		getName = articulation.Dict.Array[0].Dict[i].String
		keySwitch, _ := strconv.Atoi(articulation.Dict.Array[1].Dict[i].Integer[1])
		output := convertJson(keySwitch, getName[0])
		fmt.Println(output)
	}
	defer xmlFile.Close()
}

func convertJson(keySwitch int, getName string) string {
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
	output, err := json.Marshal(newDataStruct)
	if err != nil {

	}
	return string(output)
}

func makeFile() {
	// make file for each file traveresd
	// make the json file.
}
