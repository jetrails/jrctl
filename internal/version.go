package internal

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

const (
	CLEAR_SCREEN  string = "\033[H\033[2J"
	GREEN_RUNE           = "\033[32m%c\033[0m"
	YELLOW_RUNE          = "\033[33m%c\033[0m"
	GREEN_STRING         = "\033[32m%s\033[0m"
	YELLOW_STRING        = "\033[33m%s\033[0m"
)

type Guess []Letter

type Letter struct {
	Value      rune
	InWord     bool
	InPosition bool
}

type Wordle struct {
	Word       string
	Guesses    []Guess
	DateString string
}

type Article struct {
	Content string `json:"content"`
}

func NewWordle(word string) *Wordle {
	year, month, day := time.Now().Date()
	return &Wordle{
		Word:       strings.ToUpper(word),
		Guesses:    []Guess{},
		DateString: fmt.Sprintf("%02d/%02d/%04d", month, day, year),
	}
}

func (g Guess) Equals(target string) bool {
	value := ""
	for _, letter := range g {
		value += string(letter.Value)
	}
	return value == target
}

func (w *Wordle) GetGuess(index int) Guess {
	if index < len(w.Guesses) {
		return w.Guesses[index]
	}
	return Guess{
		Letter{Value: ' '},
		Letter{Value: ' '},
		Letter{Value: ' '},
		Letter{Value: ' '},
		Letter{Value: ' '},
	}
}

func (w *Wordle) AddGuess(guess string) {
	guess = strings.ToUpper(guess)
	entry := Guess{}
	leftover := w.Word
	for i, value := range guess {
		inPosition := rune(w.Word[i]) == value
		letter := Letter{
			Value:      value,
			InWord:     false,
			InPosition: inPosition,
		}
		if inPosition {
			leftover = leftover[:i] + " " + leftover[i+1:]
		}
		entry = append(entry, letter)
	}
	for i, letter := range entry {
		if strings.Index(leftover, string(letter.Value)) > -1 {
			entry[i].InWord = true
			leftover = strings.Replace(leftover, string(letter.Value), "*", 1)
		}
	}
	w.Guesses = append(w.Guesses, entry)
}

func (w *Wordle) HasWon() bool {
	if len(w.Guesses) > 0 {
		return w.Guesses[len(w.Guesses)-1].Equals(w.Word)
	}
	return false
}

func (w *Wordle) IsDone() bool {
	if len(w.Guesses) >= 6 {
		return true
	}
	return w.HasWon()
}

func (w *Wordle) Prompt() string {
	re := regexp.MustCompile(`[^a-z]+`)
	input := ""
	for len(input) != 5 {
		fmt.Printf("Guess: ")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		input = re.ReplaceAllString(input, "")
		if len(input) != 5 {
			fmt.Println("Error: word must be 5 alphabetic chars\n")
		}
	}
	return input
}
func (w *Wordle) Print() {
	fmt.Println()
	fmt.Println("JetRails Wordle")
	fmt.Println(w.DateString)
	fmt.Println()
	fmt.Printf(GREEN_STRING+"  = in word & correct position\n", "GREEN")
	fmt.Printf(YELLOW_STRING+" = in word & incorrect position\n", "YELLOW")
	fmt.Println()
	fmt.Println("╔═══╦═══╦═══╦═══╦═══╗")
	for i := 0; i < 6; i++ {
		guess := w.GetGuess(i)
		fmt.Printf("║")
		for _, letter := range guess {
			switch true {
			case letter.InPosition:
				fmt.Printf(" "+GREEN_RUNE+" ║", letter.Value)
			case letter.InWord:
				fmt.Printf(" "+YELLOW_RUNE+" ║", letter.Value)
			default:
				fmt.Printf(" %c ║", letter.Value)
			}
		}
		fmt.Printf("  <= Guess #%d", i+1)
		fmt.Println()
		if i < 5 {
			fmt.Println("╠═══╬═══╬═══╬═══╬═══╣")
		}
	}
	fmt.Println("╚═══╩═══╩═══╩═══╩═══╝")
	fmt.Println()
}

func (w *Wordle) Run() {
	for !w.IsDone() {
		fmt.Printf(CLEAR_SCREEN)
		w.Print()
		w.AddGuess(w.Prompt())
	}
	fmt.Printf(CLEAR_SCREEN)
	w.Print()
	if w.HasWon() {
		fmt.Println("🎉 Congratulations! You solved the word of the day!")
	} else {
		fmt.Printf("😞 Sorry, the word was %q!\n", w.Word)
	}
	fmt.Println()
}

func GetWordOfTheDay() *string {
	var request = gorequest.New()
	response, body, errs := request.
		Timeout(10 * time.Second).
		Get("https://learn.jetrails.com/api/search.json").
		End()
	if errs == nil && response != nil {
		var articles []Article
		json.Unmarshal([]byte(body), &articles)
		re := regexp.MustCompile(`\b[a-z]{5}\b`)
		words := []string{}
		seen := map[string]bool{}
		for _, article := range articles {
			matches := re.FindAllString(strings.ToLower(article.Content), -1)
			for _, word := range matches {
				if _, ok := seen[word]; !ok {
					seen[word] = true
					words = append(words, word)
				}
			}
		}
		if len(words) < 1 {
			return nil
		}
		year, month, day := time.Now().Date()
		seed := day + (int(month) * 100) + (year * 10000)
		rand.Seed(int64(seed))
		return &words[rand.Intn(len(words))]
	}
	return nil
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display current version",
	Example: text.Examples([]string{
		"jrctl version",
		"jrctl version -q",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		if !quiet {
			fmt.Printf("jrctl version %s\n", version.VersionString)
			return
		}
		if word := GetWordOfTheDay(); word != nil {
			game := NewWordle(*word)
			game.Run()
		}
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.Flags().SortFlags = true
	versionCmd.Flags().BoolP("quiet", "q", false, "be quiet, very quiet")
}
