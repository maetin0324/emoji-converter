/*
Copyright © 2023 maetin0324 <maetin0324@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"emoji-converter/dict"

	"github.com/spf13/cobra"
)


func replaceEmojiToUnicode(str string) string {
	emojis := dict.Load()
	for _, emoji := range emojis {
		str = strings.Replace(str, ":" + emoji.Aliases[0] + ":", emoji.Emoji, -1)
	}
	return str
}

var rootCmd = &cobra.Command{
	Use:   "emoji-converter",
	Short: "convert emoji to unicode and vice versa",
	Long: `convert emoji from colon format to unicode and vice versa`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var str string
		// コマンドライン引数がない場合は標準入力から読み込む
		if len(args) == 0 {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			str = string(bytes)
		} else {
			str = args[0]
		}
		// コロンで囲まれた文字列を絵文字に変換する
		str = replaceEmojiToUnicode(str)
		fmt.Println(str)
	 },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


