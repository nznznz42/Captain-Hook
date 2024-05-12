/*
Copyright © 2024 nznznz42
*/
package hookcore

import (
	"encoding/json"
	"fmt"
	"os"
)

type Ltestcmd struct {
	ConfigFile string `json:"configFile"`
	LogFile    string `json:"logFile"`
	Rflag      bool   `json:"random"`
}

type Ctestcmd struct {
	Domain string `json:"domain"`
	Port   int    `json:"port"`
}

func NewLcmd(configFileName string, logFileName string, rflag bool) Ltestcmd {
	return Ltestcmd{
		ConfigFile: configFileName,
		LogFile:    logFileName,
		Rflag:      rflag,
	}
}

func NewCcmd(domain string, port int) Ctestcmd {
	return Ctestcmd{
		Domain: domain,
		Port:   port,
	}
}

func SerializeLcmd(cmd *Ltestcmd) error {
	data := map[string]interface{}{
		"configFileName": cmd.ConfigFile,
		"logFileName":    cmd.LogFile,
		"rFlag":          cmd.Rflag,
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("Cache/Lcache.json", jsonData, 0644)
	if err != nil {
		return err
	}

	err = CacheCmd("ltest")
	if err != nil {
		println("Couldnt Cache command")
	}

	return nil
}

func DeserializeLcmd() (*Ltestcmd, error) {
	fileData, err := os.ReadFile("Cache/Lcache.json")
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

func SerializeCcmd(cmd *Ctestcmd) error {
	data := map[string]interface{}{
		"domain": cmd.Domain,
		"port":   cmd.Port,
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("Cache/Ccache.json", jsonData, 0644)
	if err != nil {
		return err
	}

	err = CacheCmd("ctest")
	if err != nil {
		println("Couldnt Cache command")
	}

	return nil
}
func DeserializeCcmd() (*Ctestcmd, error) {
	fileData, err := os.ReadFile("Cache/Ccache.json")
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	cmd := &Ctestcmd{
		Domain: data["domain"].(string),
		Port:   data["port"].(int),
	}

	return cmd, nil
}

func IsFileEmpty(filename string) (bool, error) {
	file, err := os.Open("Cache/" + filename)
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

func CacheCmd(cmdname string) error {
	data := map[string]string{
		"cmd": cmdname,
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("Cache/LastCmdCache.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadCmdCache() (string, error) {
	content, err := os.ReadFile("Cache/LastCmdCache.json")
	if err != nil {
		return "", err
	}

	var data map[string]string
	if err := json.Unmarshal(content, &data); err != nil {
		return "", err
	}
	cmdValue, ok := data["cmd"]
	if !ok {
		return "", fmt.Errorf("Cache empty")
	}

	return cmdValue, nil
}
