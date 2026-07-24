package models

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	ID             string `json:"id"`
	CategoryID     string `json:"categoryId"`
	Text           string `json:"text"`
	OptionA        string `json:"optionA"`
	OptionB        string `json:"optionB"`
	OptionC        string `json:"optionC"`
	OptionD        string `json:"optionD"`
	CorrectOption  string `json:"correctOption"`
}
