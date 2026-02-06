package main

import "fmt"

type LLM interface {
	chat(prompt string) string
}

type OpenAIModel struct {
	APIKey string
}

func (o OpenAIModel) chat(prompt string) string {
	fmt.Printf("Calling OpenAI with key %s", o.APIKey)
	return "OpenAI 回复: " + prompt
}

type DeepSeekModel struct {
	Baseurl string
}

func (d DeepSeekModel) chat(prompt string) string {
	fmt.Printf("Calling DeepSeek at url %s", d.Baseurl)
	return "DeepSeek 回复: " + prompt
}

func RunAgent(model LLM, question string) {
	answer := model.chat(question)
	fmt.Println("最终结果:", answer)
}

func main() {
	gpt := OpenAIModel{
		APIKey: "sk-123456789",
	}

	ds := DeepSeekModel{
		Baseurl: "https://api.deepseek.com",
	}
	RunAgent(gpt, "如何学习 Go 语言？")

	RunAgent(ds, "解释一下接口是什么？")
}
