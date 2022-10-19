package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	convert "github.com/basgys/goxml2json"
)

type Logic struct {
	Plist struct {
		Dict struct {
			Array []struct {
				Dict []struct {
					Key     []string `json:"key"`
					Integer []string `json:"integer"`
					String  string   `json:"string"`
					Dict    struct {
						Key     []string `json:"key"`
						Integer []string `json:"integer"`
						String  string   `json:"string"`
					} `json:"dict"`
				} `json:"dict"`
			} `json:"array"`
			String  string   `json:"string"`
			Integer string   `json:"integer"`
			Key     []string `json:"key"`
		} `json:"dict"`
		Version string `json:"-version"`
	} `json:"plist"`
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
	Color int    `json:"color"`
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
	data, err := os.ReadFile("/Users/johngoldsmith/code/AbletonExpressionMaps/NISE String Ensemble.plist")
	xml := strings.NewReader(string(data))
	articulation, err := convert.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	dataStruct := Logic{}
	json.Unmarshal([]byte(articulation.String()), &dataStruct)
	numOfSwitches := len(dataStruct.Plist.Dict.Array[1].Dict)
	for i := 0; i < numOfSwitches; i++ {
		keySwitch, _ := strconv.Atoi(dataStruct.Plist.Dict.Array[1].Dict[i].Integer[1])
		convertJson(keySwitch)
	}
}

func convertJson(keySwitch int) {
	newDataStruct := Ableton{}
	newDataStruct.Name = "-"
	newDataStruct.Key = keySwitch
	newDataStruct.Key0 = "-"
	newDataStruct.Bnk = "-"
	newDataStruct.Sub = "-"
	newDataStruct.Pgm = "-"
	newDataStruct.Ccn = "-"
	newDataStruct.Ccv = "-"
	newDataStruct.Chn = "-"
	newDataStruct.Color = 1
	output, err := json.Marshal(newDataStruct)
	if err != nil {

	}
	fmt.Println(string(output))
	// json to json
	// convert output of struct to Ableton

}

func makeFile() {
	// make file for each file traveresd
	// make the json file.
}
