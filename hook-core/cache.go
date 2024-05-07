/*
Copyright © 2024 nznznz42
*/
package hookcore

import (
	"encoding/json"
	"os"
)

type Ltestcmd struct {
	ConfigFile string `json:"configFile"`
	LogFile    string `json:"logFile"`
	Rflag      bool   `json:"random"`
}

func NewCmd(configFileName string, logFileName string, rflag bool) Ltestcmd {
	return Ltestcmd{
		ConfigFile: configFileName,
		LogFile:    logFileName,
		Rflag:      rflag,
	}
}

func Serialize(cmd *Ltestcmd) ([]byte, error) {
	jsonData, err := json.Marshal(&cmd)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func Deserialize() (*Ltestcmd, error) {
	var cmd Ltestcmd
	data, err := os.ReadFile("cache.json")
	if err != nil {
		panic("nooo")
	}

	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func IsFileEmpty() (bool, error) {
	file, err := os.Open("cache.json")
	if err != nil {
		return false, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}

	if fileInfo.Size() == 0 {
		return true, nil
	}

	return false, nil
}
