package wordy

import (
	"strings"
	"errors"
	"regexp"
	"strconv"
)

func calculate(a,b int, operator string) (int, error) {
	switch operator {
	case "plus":
		return a + b, nil
	case "minus":
		return a - b, nil
	case "multiplied by":
		return a * b, nil
	case "divided by":
		if b == 0 {
			return 0, errors.New("Cannot divide by 0")
		} else {
		return a / b, nil 
		}
	default:
		return 0, errors.New("Unknown operator")
	}
}

func Answer(question string) (int, bool) {
	numbers := []int{}
	operators := []string{}
	numberRegexp := regexp.MustCompile(`^-?\d+`)
	operationRegexp := regexp.MustCompile(`plus|minus|multiplied by|divided by`)
	endRegexp := regexp.MustCompile(`^\s*\?$`)
	prefix := "What is "
	var positions []int
	if strings.HasPrefix(question, prefix) {
		question = strings.Replace(question,prefix,"", -1)
	} else {
		return 0, false
	}
	for {
		question = strings.TrimSpace(question)
		positions = numberRegexp.FindStringIndex(question)
		if positions != nil {
			number, _ := strconv.Atoi(question[positions[0]:positions[1]])
			question = question[positions[1]:]
			numbers = append(numbers, number)
			if len(numbers) == 2 {
				if len(operators) == 1 {
				result, err := calculate(numbers[0],numbers[1],operators[0])
				if err != nil {
					return 0,false
				}
          // reset numbers and operators stacks
					numbers = []int{result}
					operators = []string{}
				} else {
          // we have 2 numbers but no operators
					return 0, false
			}
		}
			continue
		}	
		positions = operationRegexp.FindStringIndex(question)
		if positions != nil {
			if len(operators) != 0 {
				// duplicated operator
				return 0, false
			}
			if len(numbers) != 1 {
				// invalid structure
				return 0, false
			}
			operator := question[positions[0]:positions[1]]
			operators = append(operators, operator)
			question = question[positions[1]:]
			continue
		}
		if endRegexp.MatchString(question) {
			break
		}
		// if execution gets here is because there's some invalid token
		return 0,false
	}
	if len(operators) != 0 {
    // at the end we should only have 1 number and no operators
		return 0, false
		}
	if len(numbers) == 1 {
	return numbers[0], true
	} else {
		return 0, false
	}
}
