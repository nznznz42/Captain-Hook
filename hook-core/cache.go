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

func Serialize(cmd *Ltestcmd) error {
	data := map[string]interface{}{
		"configFileName": cmd.ConfigFile,
		"logFileName":    cmd.LogFile,
		"rFlag":          cmd.Rflag,
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile("cache.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Deserialize() (*Ltestcmd, error) {
	fileData, err := os.ReadFile("cache.json")
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	cmd := &Ltestcmd{
		ConfigFile: data["configFileName"].(string),
		LogFile:    data["logFileName"].(string),
		Rflag:      data["rFlag"].(bool),
	}

	return cmd, nil
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
