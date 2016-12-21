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
					case strings.Compare(message.Text, "æº«é¦¨?é?") == 0:
						outmsg.WriteString("<<<æº«é¦¨?é?>>>\r\n? ç‚º?™å€‹ç¾¤å¾ˆåµ -->\r\n?³ä?è§??¯ä»¥ ?œé??é?\r\n\r\n[?Œå­¸?ƒ] ?•ç¥¨?²è?ä¸?-->\r\n?³ä?è§?ç­†è????¯ä»¥?²è??•ç¥¨\r\n\r\n[?šè??„] ?€è¦å¤§å®¶ç??”åŠ© -->\r\n?³ä?è§?ç­†è???è«‹æ›´?°è‡ªå·±ç??¯çµ¡?¹å?")
					
					case strings.HasSuffix(message.Text, "éº¼å¸¥"):
						outmsg.WriteString(GetHandsonText(message.Text))

					case strings.Compare(message.Text, "PPAP") == 0:
						outmsg.WriteString(GetPPAPText())

					case strings.HasPrefix(message.Text, "ç¿»ç¿»"):
						outmsg.WriteString(GetTransText(gkey, strings.TrimLeft(message.Text, "ç¿»ç¿»")))

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
	outmsg.WriteString("?‘è¦ºå¾—é???)
	switch i % 20 {
	case 0:
		outmsg.WriteString("å°æ?")
	case 1:
		outmsg.WriteString("å½¬å½¬")
	case 2:
		outmsg.WriteString("?»æ¦®")
	case 3:
		outmsg.WriteString("?¯èƒ½")
	case 4:
		outmsg.WriteString("?çˆº")
	case 5:
		outmsg.WriteString("å»ºè‰¯")
	case 6:
		outmsg.WriteString("?ä?")
	case 7:
		outmsg.WriteString("å¿—å?")
	case 8:
		outmsg.WriteString("?­æ?å¦?)
	case 9:
		outmsg.WriteString("å¤§å“¥å¤?)
	case 10:
		outmsg.WriteString("ä¸‰å“¥")
	default:
		outText.WriteString(inText)
		outText.WriteString("+1")
		return outText.String()
	}
	outmsg.WriteString("æ¯”è?å¸?)
	return outmsg.String()	
}

func GetPPAPText() string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	switch i % 5 {
	case 0:
		return "I have a pencil,\r\nI have an Apple,\r\nApple pencil.\r\nI have a watch,\r\nI have an Apple,\r\nApple watch."
	case 1:
		return "?†å¸¶ä¸€?ï?è«‹ä?è¦æ?Apple Pencil?ºé€²æ°´?œè£¡ï¼Œä?ç®¡æ˜¯?‹æ??„æ˜¯é³³æ¢¨??
	case 2:
		return "?‘æ?äº†ï??™æ˜¯ä»¥æ›¸å¯«å·¥?·è?ç¨®é?é£Ÿç‰©?ºé??„é??Œæ???
	case 3:
		return "?‘ä?å¤ªæ?æ¥šPPAP?¯ä?éº¼ï?ä½†ä??¯ä»¥?æ?AAPL?„ç›¸?œè?è¨Šã€?
	case 4:
		return "?‘æ˜¯ä¸æ??¥è??±ç?ï¼?
	}
	return "?»å? siri ??
}
