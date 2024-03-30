package hookcore

import (
	"encoding/json"
	"os"
)

type ltestcmd struct {
	configFile string `json:"configFile"`
	logFile    string `json:"logFile"`
	rFlag      bool   `json:"random"`
}

func serialize(cmd *ltestcmd) ([]byte, error) {
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func deserialize(data []byte) (*ltestcmd, error) {
	var cmd ltestcmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func isFileEmpty() (bool, error) {
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
