// 同じディレクトリ内にあるemoji.jsonを構造体として読み込む

// Path: dict/dict.go
package dict

import (
	"encoding/json"
	"log"
)
import _ "embed"

//go:embed emoji.json
var emoji_dict string

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
	var emojis []Emoji
	if err := json.Unmarshal([]byte(emoji_dict), &emojis); err != nil {
		log.Fatal(err)
	}
	return emojis
}