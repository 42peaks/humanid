// WARNING: AUTO-GENERATED FILE
// This file has been automatically generated and should not be modified manually.
// Any changes made to this file may be overwritten the next time it is generated.
// For any modifications or issues, please refer to the source code or the generator script.

package humanid

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type AdjectivesType []string

func (t AdjectivesType) GetRandomWord() (string, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(t)-1)))
	if err != nil {
		return "", fmt.Errorf("humanid: Adjectives: error: %s", err.Error())
	}

	return t[index.Int64()], nil
}

func (t AdjectivesType) MustGetRandomWord() string {
	word, err := t.GetRandomWord()
	if err != nil {
		panic(err)
	}

	return word
}

var (
	Adjectives = AdjectivesType{
		"attractive",
		"bald",
		"beautiful",
		"chubby",
		"clean",
		"dazzling",
		"drab",
		"elegant",
		"fancy",
		"fit",
		"flabby",
		"glamorous",
		"gorgeous",
		"handsome",
		"long",
		"magnificent",
		"muscular",
		"plain",
		"plump",
		"quaint",
		"scruffy",
		"shapely",
		"short",
		"skinny",
		"stocky",
		"ugly",
		"unkempt",
		"unsightly",
		"ashy",
		"black",
		"blue",
		"gray",
		"green",
		"icy",
		"lemon",
		"mango",
		"orange",
		"purple",
		"red",
		"salmon",
		"white",
		"yellow",
		"alive",
		"better",
		"careful",
		"clever",
		"dead",
		"easy",
		"famous",
		"gifted",
		"hallowed",
		"helpful",
		"important",
		"inexpensive",
		"mealy",
		"mushy",
		"odd",
		"poor",
		"powerful",
		"rich",
		"shy",
		"tender",
		"unimportant",
		"uninterested",
		"vast",
		"wrong",
		"aggressive",
		"agreeable",
		"ambitious",
		"brave",
		"calm",
		"delightful",
		"eager",
		"faithful",
		"gentle",
		"happy",
		"jolly",
		"kind",
		"lively",
		"nice",
		"obedient",
		"polite",
		"proud",
		"silly",
		"thankful",
		"victorious",
		"witty",
		"wonderful",
		"zealous",
		"angry",
		"bewildered",
		"clumsy",
		"defeated",
		"embarrassed",
		"fierce",
		"grumpy",
		"helpless",
		"itchy",
		"jealous",
		"lazy",
		"mysterious",
		"nervous",
		"obnoxious",
		"panicky",
		"pitiful",
		"repulsive",
		"scary",
		"thoughtless",
		"uptight",
		"worried",
		"broad",
		"chubby",
		"crooked",
		"curved",
		"deep",
		"flat",
		"high",
		"hollow",
		"low",
		"narrow",
		"refined",
		"round",
		"shallow",
		"skinny",
		"square",
		"steep",
		"straight",
		"wide",
		"big",
		"colossal",
		"fat",
		"gigantic",
		"great",
		"huge",
		"immense",
		"large",
		"little",
		"mammoth",
		"massive",
		"microscopic",
		"miniature",
		"petite",
		"puny",
		"scrawny",
		"short",
		"small",
		"tall",
		"teeny",
		"tiny",
		"crashing",
		"deafening",
		"echoing",
		"faint",
		"harsh",
		"hissing",
		"howling",
		"loud",
		"melodic",
		"noisy",
		"purring",
		"quiet",
		"rapping",
		"raspy",
		"rhythmic",
		"screeching",
		"shrilling",
		"squeaking",
		"thundering",
		"tinkling",
		"wailing",
		"whining",
		"whispering",
		"ancient",
		"brief",
		"early",
		"fast",
		"future",
		"late",
		"long",
		"modern",
		"old",
		"old-fashioned",
		"prehistoric",
		"quick",
		"rapid",
		"short",
		"slow",
		"swift",
		"young",
		"acidic",
		"bitter",
		"cool",
		"creamy",
		"delicious",
		"disgusting",
		"fresh",
		"greasy",
		"juicy",
		"hot",
		"moldy",
		"nutritious",
		"nutty",
		"putrid",
		"rancid",
		"ripe",
		"rotten",
		"salty",
		"savory",
		"sour",
		"spicy",
		"spoiled",
		"stale",
		"sweet",
		"tangy",
		"tart",
		"tasteless",
		"tasty",
		"yummy",
		"breezy",
		"bumpy",
		"chilly",
		"cold",
		"cool",
		"cuddly",
		"damaged",
		"damp",
		"dirty",
		"dry",
		"flaky",
		"fluffy",
		"freezing",
		"greasy",
		"hot",
		"icy",
		"loose",
		"melted",
		"prickly",
		"rough",
		"shaggy",
		"sharp",
		"slimy",
		"sticky",
		"strong",
		"tight",
		"uneven",
		"warm",
		"weak",
		"wet",
		"wooden",
		"abundant",
		"billions",
		"enough",
		"few",
		"full",
		"hundreds",
		"incalculable",
		"limited",
		"little",
		"many",
		"most",
		"millions",
		"numerous",
		"scarce",
		"some",
		"sparse",
		"substantial",
		"thousands",
		
	}
)