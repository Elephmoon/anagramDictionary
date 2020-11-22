package models

type AnagramResponse struct {
	Word     string   `json:"word"`
	Anagrams []string `json:"anagrams"`
}
