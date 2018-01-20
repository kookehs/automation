package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

var (
	Afinn     map[string]int
	Corpus    map[string][]string
	Resources map[string]string
)

func init() {
	Corpus = make(map[string][]string)
	LoadJson("./data/commands.json", &Resources)

	for k, v := range Resources {
		data := make([]string, 0)
		LoadJson("./data/"+v, &data)
		Corpus[k] = data
	}

	file, err := os.Open("./data/afinn.txt")

	if err != nil {
		log.Println("Error opening afinn: ", err)
	}

	defer file.Close()
	Afinn = make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "\t")
		sentiment, err := strconv.Atoi(tokens[1])

		if err != nil {
			log.Println("Error converting sentiment: ", err)
			continue
		}

		Afinn[tokens[0]] = sentiment
	}
}

func Classify(s string) string {
	var match string
	highestSimilarity := math.SmallestNonzeroFloat64

	for k, v := range Corpus {
		similarity := Compare(s, v)

		if similarity > highestSimilarity {
			highestSimilarity = similarity
			match = k
		}
	}

	return match
}

func Compare(s string, c []string) float64 {
	// TODO: Store processed models
	reducer := nlp.NewTruncatedSVD(4)
	transformer := nlp.NewTfidfTransformer()
	vectoriser := nlp.NewCountVectoriser(true)
	lsiPipeline := nlp.NewPipeline(vectoriser, transformer, reducer)
	lsi, err := lsiPipeline.FitTransform(c...)

	var highestSimilarity float64

	if err != nil {
		log.Println("Failed to process documents: ", err)
		return highestSimilarity
	}

	queryVector, err := lsiPipeline.Transform(s)

	if err != nil {
		log.Println("failed to process documents: ", err)
		return highestSimilarity
	}

	_, docs := lsi.Dims()

	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))

		if similarity > highestSimilarity {
			highestSimilarity = similarity
		}
	}

	return highestSimilarity
}

func Sentiment(s string) int {
	var score int

	tokens := strings.Split(s, " ")

	for _, token := range tokens {
		if v, ok := Afinn[token]; ok {
			score += v
		}
	}

	return score
}
