package woordenlijst

func (p *Paradigm) PartOfSpeechDetails() PartOfSpeechDetails {
	return parsePartOfSpeech(p.POS)
}

// PartOfSpeech returns the part of speech of the paradigm.
func (p *Paradigm) PartOfSpeech() string {
	return p.PartOfSpeechDetails().Part
}
