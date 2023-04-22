package console

import (
	"strings"
	"unicode"
)

const spellParamsLen = 1

type spellCommand struct {
}

func newSpellCommand() *spellCommand {
	return &spellCommand{}
}

func (s *spellCommand) Process(params []string) error {
	if len(params) != spellParamsLen {
		return InvalidInput
	}

	result, err := s.action(params[0])
	if err != nil {
		return err
	}

	return s.show(result)
}

func (s *spellCommand) show(letters []rune) error {
	var stringBuilder strings.Builder
	for _, letter := range letters {
		stringBuilder.WriteRune(letter)
		stringBuilder.WriteRune(' ')
	}
	stringBuilder.WriteRune('\n')
	return nil
}

func (s *spellCommand) action(word string) ([]rune, error) {
	result := make([]rune, len(word))
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			return nil, InvalidInput
		}
		result = append(result, letter)
	}

	return result, nil
}
