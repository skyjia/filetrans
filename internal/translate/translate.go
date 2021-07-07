package translate

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/jinmukeji/go-pkg/v2/crypto/hash"
)

type Translator struct {
	appID string
	key   string

	delay time.Duration // 毫秒

	client *http.Client
}

func NewTranslator(appID, key string, delay time.Duration) *Translator {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return &Translator{
		appID:  appID,
		key:    key,
		client: client,
		delay:  delay,
	}
}

func (t *Translator) Translate(query string) string {
	time.Sleep(t.delay)

	const (
		from = "auto"
		to   = "zh"
	)

	salt := strconv.FormatInt(time.Now().Unix(), 10)
	str := fmt.Sprintf("%s%s%s%s", t.appID, query, salt, t.key)
	sign := md5(str)

	params := url.Values{}
	params.Add("q", query)
	params.Add("appid", t.appID)
	params.Add("salt", salt)
	params.Add("from", from)
	params.Add("to", to)
	params.Add("sign", sign)

	u := fmt.Sprintf("https://fanyi-api.baidu.com/api/trans/vip/translate?%s", params.Encode())
	resp, err := http.Get(u)
	if err != nil {
		log.Println(err)
		return query
	}
	defer resp.Body.Close()

	var tResp transResp
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&tResp); err != nil {
		log.Println(err)
		return query
	}

	if tResp.ErrorCode != "" {
		log.Println("ErrorCode:", tResp.ErrorCode)
		return query
	}

	// fmt.Println(tResp)

	if len(tResp.TransResult) > 0 {
		txt := repaceChars(tResp.TransResult[0].Dst)
		return txt
	}

	return query
}

func md5(s string) string {
	r := hash.MD5String(s)
	return hash.HexString(r)
}

func repaceChars(s string) string {
	r := strings.ReplaceAll(s, "，", "_") // 中文逗号
	r = strings.ReplaceAll(r, ",", "_")  // 英文逗号
	r = strings.ReplaceAll(r, " ", "_")  // 空格
	r = strings.ReplaceAll(r, "<", "_")  // < (less than)
	r = strings.ReplaceAll(r, ">", "_")  // > (greater than)
	r = strings.ReplaceAll(r, ":", "_")  // : (colon)
	r = strings.ReplaceAll(r, "\"", "_") // " (double quote)
	r = strings.ReplaceAll(r, "/", "_")  // / (forward slash)
	r = strings.ReplaceAll(r, "\\", "_") // \ (backslash)
	r = strings.ReplaceAll(r, "|", "_")  // | (vertical bar or pipe)
	r = strings.ReplaceAll(r, "?", "_")  // ? (question mark)
	r = strings.ReplaceAll(r, "*", "_")  // * (asterisk)

	return r
}

type transResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type transResp struct {
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []transResult `json:"trans_result"`
	ErrorCode   string        `json:"error_code"`
	// SrcTTS      string        `json:"src_tts"`
	// DstTTS      string        `json:"dst_tts"`
	// Dict        string        `json:"dict"`
}
