package handlers

import (
	"codeChallenge/helper"
	"encoding/json"
	"net/http"
)

type TextStruct struct {
	Text []string
}

func Index(w http.ResponseWriter, r *http.Request) {
	var outputTexts []string
	textStruct := TextStruct{}
	decoder := json.NewDecoder(r.Body)
	for decoder.More() {
		err := decoder.Decode(&textStruct)
		if err != nil {
			helper.SendError(w, err.Error())
		}
		for _, text := range textStruct.Text {
			hasRepeatedCharacters := helper.RepeatedCharacters(text)
			if hasRepeatedCharacters {
				outputTexts = append(outputTexts, text)
			}
		}
	}
	helper.SendData(w, outputTexts)
}
