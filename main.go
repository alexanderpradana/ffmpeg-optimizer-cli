package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Source     string `json:"source-folder"`
	Dest       string `json:"dest-folder"`
	CodecVideo string `json:"compression-video"`
	CodecAudio string `json:"compression-audio"`
	AudioSize  string `json:"audio-size"`
	CRF        string `json:"crf-quality"`
	OutputExt  string `json:"output-extension"`
	Preset     string `json:"preset"`
	MoovAtom   string `json:"moovatom"`
}

func main() {
	fmt.Println("Loading configuration...")
	config := LoadConfig("config.json")
	fmt.Println(config)
	var sb strings.Builder
	sb.WriteString("ffmpeg -i video_2021-06-06_00-40-01.mp4")
	sb.WriteString(" -" + config.MoovAtom)
	sb.WriteString(" -c:v " + config.CodecAudio)
	sb.WriteString(" -crf " + config.CRF)
	sb.WriteString(" -c:a " + config.CodecAudio)
	sb.WriteString(" -b:a " + config.AudioSize)
	sb.WriteString(config.Dest + "video_2021-06-06_00-40-01.mp4")
	cmd := sb.String()
	fmt.Println("Calling ffmpeg with: " + cmd)

}

func LoadConfig(file string) *Config {
	var config Config
	absPath, _ := filepath.Abs(file)
	configFile, err := os.Open(absPath)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	//byteValue, _ := ioutil.ReadAll(configFile)
	//json.Unmarshal(byteValue, &config)
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config
}
