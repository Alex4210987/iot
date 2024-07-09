package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GPT struct {
	Choices []Choice `json:"choices"`
	Created int64    `json:"created"`
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	FinishReason *string  `json:"finish_reason,omitempty"`
	Index        *int64   `json:"index,omitempty"`
	Message      *Message `json:"message,omitempty"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type Usage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

func GetAdvice(status string) (string, error) {

	url := "https://api.chatanywhere.com.cn/v1/chat/completions"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
  "model": "gpt-3.5-turbo",
  "messages": [{"role": "user", "content": "你是一个温室的管理者下面是当前温室的状态，%s,请根据这些数据提供建议，要求对每一个数据进行分析。"}]
}`, status))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", "Bearer sk-Hrw55jIAR8DANkoXw5fCyHgrVOW18cOyR57jsE6hvee7Q0gq")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))
	//解析json
	gpt := GPT{}
	err = json.Unmarshal(body, &gpt)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(gpt.Choices[0].Message.Content)
	return gpt.Choices[0].Message.Content, nil
}
