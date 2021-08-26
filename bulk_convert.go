package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	var sb strings.Builder
	sb.WriteString("ffmpeg -i video_2021-06-06_00-40-01.mp4")
	sb.WriteString(" -" + config.MoovAtom)
	sb.WriteString(" -c:v " + config.CodecAudio)
	sb.WriteString(" -crf " + config.CRF)
	sb.WriteString(" -c:a " + config.CodecAudio)
	sb.WriteString(" -b:a " + config.AudioSize)
	sb.WriteString(config.Dest + "video_2021-06-06_00-40-01.mp4")
}

func LoadConfig(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
