package woordenlijst

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	host            = "https://woordenlijst.org"
	path            = "MolexServe/lexicon/find_wordform"
	DefaultDatabase = "gig_pro_wrdlst"
)

func NewSearchParameters(word string, partsOfSpeech []PartOfSpeech) *SearchParameters {
	return &SearchParameters{
		PartsOfSpeech: partsOfSpeech,
		WordForm:      word,
		Paradigm:      true,
		Diminutive:    true,
	}
}

// ToQueryString returns the query string for this search parameter.
func (p *SearchParameters) ToQueryString() string {
	queryParameters := map[string]string{}
	if len(p.PartsOfSpeech) > 0 {
		queryParameters["part_of_speech"] = partsOfSpeechForUrl(p.PartsOfSpeech)
	}
	if p.Paradigm {
		queryParameters["paradigm"] = "true"
	}
	if p.Diminutive {
		queryParameters["diminutive"] = "true"
	}
	query := strings.Join([]string{
		"database=" + DefaultDatabase,
		"wordform=" + p.WordForm,
	}, "&")
	for k, v := range queryParameters {
		query += "&" + k + "=" + v
	}
	return query
}

func fetchXML(params *SearchParameters) ([]byte, error) {
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}

	c := http.DefaultClient

	url := fmt.Sprintf("%s/%s?%s", host, path, params.ToQueryString())
	resp, err := c.Get(url)
	fmt.Println(url)
	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %w", url, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}
	fmt.Println(string(body))
	return body, nil
}
