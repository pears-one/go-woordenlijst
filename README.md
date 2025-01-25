# Woordenlijst

The `woordenlijst` Go package allows you to search and retrieve information about Dutch words from woordenlijst.org. It currently exposes functionality which I need but this can be expanded.

## Installation

To install the package, use the following command:

```bash
go get github.com/pears-one/go-woordenlijst
```

## Usage

### Importing the Package

In your Go program, import the package:

```go
import "github.com/pears-one/go-woordenlijst"
```

### Searching for Nouns

To search for nouns, use the `SearchNouns` function:

```go
package main

import (
    "fmt"
    "log"
    "github.com/pears-one/go-woordenlijst"
)

func main() {
    nouns, err := woordenlijst.SearchNouns("kantoor")
    if err != nil {
        log.Fatal(err)
    }

    for _, noun := range nouns {
        article, _ := noun.Article()
        fmt.Printf("Word with article: %s %s\n", article, noun.Word())
    }
}
```

### Searching for Verbs

To search for verbs, use the `SearchVerbs` function:

```go
package main

import (
    "fmt"
    "log"
    "github.com/pears-one/go-woordenlijst"
)

func main() {
    verbs, err := woordenlijst.SearchVerbs("voetbal")
    if err != nil {
        log.Fatal(err)
    }

    for _, verb := range verbs {
        fmt.Printf("Verb: %s\n", verb.Word())
    }
}
```

### Custom Search

For a custom search, you can use the `SearchWithParams` function. First, create a `SearchParameters` struct:

```go
params := &woordenlijst.SearchParameters{
    WordForm:      "huis",
    PartsOfSpeech: []woordenlijst.PartOfSpeech{woordenlijst.Noun},
    Paradigm:      true,
    Diminutive:    true,
}
```

Then, call the `SearchWithParams` function:

```go
results, err := woordenlijst.SearchWithParams(params)
if err != nil {
    log.Fatal(err)
}

for _, lemma := range results {
    fmt.Printf("Word: %s\n", lemma.Word())
}
```

### Lemma Methods

Once you have a `Lemma` object, you can use various methods to get more information. The functions available depend of the word's part of speech. For example the following functions are available for nouns:


- `SingularForm() (string, error)`: Returns the singular form.
- `PluralForm() (string, error)`: Returns the plural form.
- `Gender() (string, error)`: Returns the gender.
- `Article() (string, error)`: Returns the article (de/het).

Example:

```go
for _, lemma := range results {
    singular, _ := lemma.SingularForm()
    plural, _ := lemma.PluralForm()
    gender, _ := lemma.Gender()
    article, _ := lemma.Article()

    fmt.Printf("Lemma: %s, Singular: %s, Plural: %s, Gender: %s, Article: %s\n",
               lemma.Lemma, singular, plural, gender, article)
}
```

## Functions

### Search Functions

- `func Search(word string, partsOfSpeech []PartOfSpeech) ([]Lemma, error)`
- `func SearchNouns(word string) ([]Lemma, error)`
- `func SearchVerbs(word string) ([]Lemma, error)`
- `func SearchWithParams(params *SearchParameters) ([]Lemma, error)`

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## Contact

For any questions or issues, please open an issue on GitHub.