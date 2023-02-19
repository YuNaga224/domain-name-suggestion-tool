package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *word `json:"noun"`
	Verb *word `json:"verb"`
}

type word struct {
	Syn []string `json:syn`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	responce, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("bighuge:: %qの類語検索に失敗しました: %v", term, err)
	}
	var data synonyms
	defer responce.Body.Close()
	if err := json.NewDecoder(responce.Body).Decode(&data); err != nil {
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return syns, nil
}

