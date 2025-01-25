package woordenlijst

import "fmt"

func SearchWithParams(params *SearchParameters) ([]Lemma, error) {
	data, err := fetchXML(params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch xml: %w", err)
	}
	root, err := parseXML(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse xml: %w", err)
	}
	return root.Lemmas, nil
}

func Search(word string, partsOfSpeech []PartOfSpeech) ([]Lemma, error) {
	params := NewSearchParameters(word, partsOfSpeech)
	return SearchWithParams(params)
}

func SearchNouns(word string) ([]Lemma, error) {
	return Search(word, []PartOfSpeech{Noun})
}

func SearchVerbs(word string) ([]Lemma, error) {
	return Search(word, []PartOfSpeech{Verb})
}
