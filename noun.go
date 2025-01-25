package woordenlijst

import (
	"fmt"
)

func (l *Lemma) IsNoun() bool {
	return l.PartOfSpeech()[:3] == Noun.String()
}

func (l *Lemma) PluralForm() (string, error) {
	if !l.IsNoun() {
		return "", fmt.Errorf("word: %s is a %s, must be a noun", l.Lemma, l.PartOfSpeech())
	}
	for _, p := range l.P.Paradigms {
		pos := p.PartOfSpeechDetails()
		if pos.Details["number"] == "pl" {
			return p.Wordform, nil
		}
	}
	return "", fmt.Errorf("word: %s has no plural form", l.Lemma)
}

func (l *Lemma) SingularForm() (string, error) {
	if !l.IsNoun() {
		return "", fmt.Errorf("word: %s is a %s, must be a noun", l.Lemma, l.PartOfSpeech())
	}
	for _, p := range l.P.Paradigms {
		pos := p.PartOfSpeechDetails()
		if pos.Details["number"] == "sg" {
			return p.Wordform, nil
		}
	}
	return "", fmt.Errorf("word: %s has no plural form", l.Lemma)
}

func (l *Lemma) Gender() (string, error) {
	if !l.IsNoun() {
		return "", fmt.Errorf("word: %s is a %s, must be a noun", l.Lemma, l.PartOfSpeech())
	}
	pos := l.PartOfSpeechDetails()
	gender := pos.Details["gender"]
	if gender == "" {
		return "", fmt.Errorf("word: %s has no gender", l.Lemma)
	}
	return gender, nil
}

func (l *Lemma) Article() (string, error) {
	gender, err := l.Gender()
	if err != nil {
		return "", err
	}
	if gender == "n" {
		return "het", nil
	}
	return "de", nil
}
