// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"bytes"
	"math/rand"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var gkey string

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	gkey = os.Getenv("GOOGLEAPIKEY")
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				var outmsg bytes.Buffer

				switch {
					case strings.Compare(message.Text, "溫馨?��?") == 0:
						outmsg.WriteString("<<<溫馨?��?>>>\r\n?�為?�個群很吵 -->\r\n?��?�??�以 ?��??��?\r\n\r\n[?�學?�] ?�票?��?�?-->\r\n?��?�?筆�????�以?��??�票\r\n\r\n[?��??�] ?�要大家�??�助 -->\r\n?��?�?筆�???請更?�自己�??�絡?��?")
					
					case strings.HasSuffix(message.Text, "麼帥"):
						outmsg.WriteString(GetHandsonText(message.Text))

					case strings.Compare(message.Text, "PPAP") == 0:
						outmsg.WriteString(GetPPAPText())

					case strings.HasPrefix(message.Text, "翻翻"):
						outmsg.WriteString(GetTransText(gkey, strings.TrimLeft(message.Text, "翻翻")))

					default:
						continue
				}
				
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outmsg.String())).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func GetHandsonText(inText string) string {
	var outmsg bytes.Buffer	
	var outText bytes.Buffer
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	outmsg.WriteString("?�覺得�???)
	switch i % 20 {
	case 0:
		outmsg.WriteString("小�?")
	case 1:
		outmsg.WriteString("彬彬")
	case 2:
		outmsg.WriteString("?�榮")
	case 3:
		outmsg.WriteString("?�能")
	case 4:
		outmsg.WriteString("?�爺")
	case 5:
		outmsg.WriteString("建良")
	case 6:
		outmsg.WriteString("?��?")
	case 7:
		outmsg.WriteString("志�?")
	case 8:
		outmsg.WriteString("?��?�?)
	case 9:
		outmsg.WriteString("大哥�?)
	case 10:
		outmsg.WriteString("三哥")
	default:
		outText.WriteString(inText)
		outText.WriteString("+1")
		return outText.String()
	}
	outmsg.WriteString("比�?�?)
	return outmsg.String()	
}

func GetPPAPText() string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	switch i % 5 {
	case 0:
		return "I have a pencil,\r\nI have an Apple,\r\nApple pencil.\r\nI have a watch,\r\nI have an Apple,\r\nApple watch."
	case 1:
		return "?�帶一?��?請�?要�?Apple Pencil?�進水?�裡，�?管是?��??�是鳳梨??
	case 2:
		return "?��?了�??�是以書寫工?��?種�?食物?��??��??��???
	case 3:
		return "?��?太�?楚PPAP?��?麼�?但�??�以?��?AAPL?�相?��?訊�?
	case 4:
		return "?�是不�??��??��?�?
	}
	return "?��? siri ??
}
