package woordenlijst

func (l *Lemma) PartOfSpeechDetails() PartOfSpeechDetails {
	return parsePartOfSpeech(l.POS)
}

func (l *Lemma) PartOfSpeech() string {
	return l.PartOfSpeechDetails().Part
}

func (l *Lemma) Word() string {
	return l.Lemma
}