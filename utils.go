package woordenlijst

import (
	"net/url"
	"strings"
)

func partsOfSpeechForUrl(partsOfSpeech []PartOfSpeech) string {
	partsOfSpeechForQuery := []string{}
	for _, partOfSpeech := range partsOfSpeech {
		partsOfSpeechForQuery = append(partsOfSpeechForQuery, partOfSpeech.String())
	}
	return url.QueryEscape(strings.Join(partsOfSpeechForQuery, "|"))
}

type PartOfSpeechDetails struct {
	Part string
	Details      map[string]string
}

// parsePartOfSpeech parses a part-of-speech string and returns a PartOfSpeechDetails struct
func parsePartOfSpeech(pos string) PartOfSpeechDetails {
	parts := strings.Split(pos, "(")
	part := parts[0]
	details := make(map[string]string)

	if len(parts) > 1 {
		detailsString := strings.TrimSuffix(parts[1], ")")
		pairs := strings.Split(detailsString, ",")
		for _, pair := range pairs {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				details[kv[0]] = kv[1]
			}
		}
	}

	return PartOfSpeechDetails{
		Part:    part,
		Details: details,
	}
}