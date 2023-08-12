// 同じディレクトリ内にあるemoji.jsonを構造体として読み込む

// Path: dict/dict.go
package dict

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Emoji struct {
	Emoji string `json:"emoji"`
	Description string `json:"description"`
	Category string `json:"category"`
	Aliases []string `json:"aliases"`
	Tags []string `json:"tags"`
	UnicodeVersion string `json:"unicode_version"`
	IosVersion string `json:"ios_version"`
}

func Load() []Emoji {
	bytes, err := ioutil.ReadFile("dict/emoji.json")
	if err != nil {
		log.Fatal(err)
	}
	var emojis []Emoji
	if err := json.Unmarshal(bytes, &emojis); err != nil {
		log.Fatal(err)
	}
	return emojis
}