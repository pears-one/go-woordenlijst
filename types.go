package woordenlijst

type SearchParameters struct {
	WordForm      string `validate:"required"`
	PartsOfSpeech []PartOfSpeech
	Paradigm      bool
	Diminutive    bool
}

type Lemma struct {
	DiminutiveInfo string `xml:"diminutive_info"`
	Diminutives    string `xml:"diminutives"`
	EntryType      string `xml:"entry_type"`
	ExternalLink   string `xml:"external_link"`
	Gloss          string `xml:"gloss"`
	Keurmerk       bool   `xml:"keurmerk"`
	Label          string `xml:"label"`
	Lemma          string `xml:"lemma"`
	LemmaID        int    `xml:"lemma_id"`
	Message        string `xml:"message"`
	Nuancerende    string `xml:"nuancerende_opmerking"`
	Online         bool   `xml:"online"`
	POS            string `xml:"lemma_part_of_speech"`
	P              P      `xml:"paradigm"`
	Parent         string `xml:"parent"`
	PartNumber     int    `xml:"part_number"`
	Pronounciation string `xml:"pronounciation"`
	Source         string `xml:"source"`
	Subset         string `xml:"subset"`
	Taaladvies     string `xml:"taaladvies"`
	Taalvariant    string `xml:"taalvariant"`
	Trademark      bool   `xml:"trademark"`
	VerbFeatures   string `xml:"verb_features"`
	Wordparts      string `xml:"wordparts"`
	WordpartsInfo  string `xml:"wordparts_info"`
}

type root struct {
	Gloss          string  `xml:"gloss"`
	Message        string  `xml:"message"`
	Lemma          string  `xml:"lemma"`
	LemmaID        int     `xml:"lemma_id"`
	Lemmas         []Lemma `xml:"found_lemmata"`
	POS            string  `xml:"part_of_speech"`
	Pronounciation string  `xml:"pronounciation"`
	Wordform       string  `xml:"wordform"`
}

type P struct {
	Paradigms []Paradigm `xml:"paradigm"`
}

type Paradigm struct {
	Arch        bool   `xml:"arch"`
	GroupLabel  string `xml:"group_label"`
	Hyphenation string `xml:"hyphenation"`
	POS         string `xml:"part_of_speech"`
	Position    int    `xml:"position"`
	Wordform    string `xml:"wordform"`
	WordformID  int    `xml:"wordform_id"`
}

type PartOfSpeech int

const (
	Adverb       PartOfSpeech = iota // bijwoord
	Adjective                        // bijvoeglijk naamwoord
	Verb                             // werkwoord
	Noun                             // zelfstandig naamwoord
	Preposition                      // voorzetsel
	Conjunction                      // voegwoord
	Pronoun                          // voornaamwoord
	Numeral                          // telwoord
	Interjection                     // tussenwerpsel
	Other                            // overig
)

func (p PartOfSpeech) String() string {
	switch p {
	case Adverb:
		return "ADV"
	case Adjective:
		return "AA"
	case Verb:
		return "VRB"
	case Noun:
		return "NOU"
	case Preposition:
		return "ADP"
	case Conjunction:
		return "CONJ"
	case Pronoun:
		return "PD"
	case Numeral:
		return "NUM"
	case Interjection:
		return "INT"
	case Other:
		return "RES"
	default:
		return ""
	}
}

func (p PartOfSpeech) ToEnglish() string {
	switch p {
	case Adverb:
		return "adverb"
	case Adjective:
		return "adjective"
	case Verb:
		return "verb"
	case Noun:
		return "noun"
	case Preposition:
		return "preposition"
	case Conjunction:
		return "conjunction"
	case Pronoun:
		return "pronoun"
	case Numeral:
		return "numeral"
	case Interjection:
		return "interjection"
	case Other:
		return "other"
	default:
		return ""
	}
}
