package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"flag"
	"fmt"
)

func fromBoolToInt(commands *[]int, value *bool) []int {
	if *value {
		*commands = append(*commands, 1)
	} else {
		*commands = append(*commands, 0)
	}
	return *commands
}

func parseArguments() ([]int, string, string) {
	paramC := flag.Bool("c", false, "number of meetings")
	paramD := flag.Bool("d", false, "")
	paramU := flag.Bool("u", false, "")
	paramF := flag.Int("f", 0, "")
	paramS := flag.Int("s", 0, "")
	paramI := flag.Bool("i", false, "")

	flag.Parse()

	arguments := make([]int, 0)

	fromBoolToInt(&arguments, paramC)
	fromBoolToInt(&arguments, paramD)
	fromBoolToInt(&arguments, paramU)
	arguments = append(arguments, *paramF)
	arguments = append(arguments, *paramS)
	fromBoolToInt(&arguments, paramI)

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	return arguments, inputFile, outputFile
}

func checkExtraOptions(options *[]int) bool {
	countOptions := 0

	for _, arg := range *options {
		if arg == 1 {
			countOptions++
		}
	}
	if countOptions > 1 {
		return true
	}
	return false
}

func optionDorU(data *[]string, repeated bool) []string {
	keys := make(map[string]int)
	repeatedOrNot := make([]string, 0)
	for idx, arg := range *data {
		if repeated {
			if idx == 0 || arg == (*data)[idx-1] && len(repeatedOrNot) != 0 && keys[arg]+1 == idx {
				continue
			}
			if arg == (*data)[idx-1] {
				keys[arg] = idx
				repeatedOrNot = append(repeatedOrNot, arg)
			}
		} else {
			if idx != len(*data)-1 && arg != (*data)[idx+1] && idx != 0 && arg != (*data)[idx-1] {
				repeatedOrNot = append(repeatedOrNot, arg)
			}
		}
	}
	return repeatedOrNot
}

type line struct {
	str   string
	times int
}

func optionC(data *[]string) []line {
	result := make([]line, 0)
	keys := make(map[string]int)
	for idx, arg := range *data {
		if _, ok := keys[arg]; !ok {
			keys[arg] = 1
		} else {
			keys[arg]++
		}
		if idx == len(*data)-1 || (*data)[idx+1] != arg {
			result = append(result, line{arg, keys[arg]})
			keys = make(map[string]int)
		}
	}
	return result
}

func optionIorNan(data *[]string) []string {
	uniqStrings := make([]string, 0)
	keys := make(map[string]int)
	for idx, arg := range *data {
		lower_arg := strings.ToLower(arg)
		if idx == 0 {
			keys[lower_arg] = idx
			uniqStrings = append(uniqStrings, arg)
			continue
		}
		before := strings.ToLower((*data)[idx-1])
		if lower_arg == before && keys[lower_arg]+1 == idx {
			keys[lower_arg] = idx
			continue
		}
		if lower_arg != strings.ToLower((*data)[idx-1]) {
			keys[lower_arg] = idx
			uniqStrings = append(uniqStrings, arg)
		}
	}
	return uniqStrings
}

func fromMapToArray(mapData []line) []string {
	result := make([]string, 0)
	for _, value := range mapData {
		result = append(result, strconv.Itoa(value.times)+" "+value.str)
	}
	return result
}

func deleteNumWords(word string, num int) string {
	splitSpace := strings.Split(word, " ")
	if len(splitSpace) < num {
		return word
	}
	newWord := strings.Join(splitSpace[num:], " ")
	return newWord
}

func optionF(data *[]string, num int) []string {
	uniqStrings := make([]string, 0)
	uniqStrings = append(uniqStrings, (*data)[0])
	currentUniq := deleteNumWords((*data)[0], num)
	for _, arg := range *data {
		deleteWords := deleteNumWords(arg, num)
		if deleteWords != currentUniq {
			uniqStrings = append(uniqStrings, arg)
			currentUniq = deleteWords
		}
	}
	return uniqStrings
}

func missNumChars(word string, chars int) string {
	if len(word) < chars {
		return word
	}
	return word[chars:]
}

func optionS(data *[]string, chars int) []string {
	uniqStrings := make([]string, 0)
	uniqStrings = append(uniqStrings, (*data)[0])
	currentUniq := missNumChars((*data)[0], chars)
	for _, arg := range *data {
		deleteWords := missNumChars(arg, chars)
		if deleteWords != currentUniq {
			uniqStrings = append(uniqStrings, arg)
			currentUniq = deleteWords
		}
	}
	return uniqStrings
}

func printArray(data *[]string) {
	for _, arg := range *data {
		fmt.Println(arg)
	}
}

func NoFlags(data *[]string) []string {
	uniqStrings := make([]string, 0)
	keys := make(map[string]int)
	for idx, arg := range *data {
		if idx == 0 {
			keys[arg] = idx
			uniqStrings = append(uniqStrings, arg)
			continue
		}
		if arg == (*data)[idx-1] && keys[arg]+1 == idx {
			keys[arg] = idx
			continue
		}
		if arg != strings.ToLower((*data)[idx-1]) {
			keys[arg] = idx
			uniqStrings = append(uniqStrings, arg)
		}
	}
	return uniqStrings
}

func correctUniqWork(options *[]int, data *[]string) []string {
	result := make([]string, 0)

	if reflect.DeepEqual(*options, []int{0, 0, 0, 0, 0, 0}) {
		result = NoFlags(data)
	}

	if (*options)[3] != 0 {
		result = optionIorNan(data)
	}

	if (*options)[5] != 0 {
		if len(result) == 0 {
			result = optionS(data, (*options)[5])
		} else {
			result = optionS(&result, (*options)[5])
		}
	}

	if (*options)[4] != 0 {
		if len(result) == 0 {
			result = optionF(data, (*options)[4])
		} else {
			result = optionF(&result, (*options)[4])
		}
	}

	if (*options)[1] != 0 {
		if len(result) == 0 {
			result = optionDorU(data, true)
		} else {
			result = optionDorU(&result, true)
		}
	}

	if (*options)[2] != 0 {
		if len(result) == 0 {
			result = optionDorU(data, false)
		} else {
			result = optionDorU(&result, false)
		}
	}

	if (*options)[0] != 0 {
		if len(result) == 0 {
			mapStrings := optionC(data)
			result = fromMapToArray(mapStrings)
		} else {
			mapStrings := optionC(&result)
			result = fromMapToArray(mapStrings)
		}
	}

	return result
}

func checkOutputFile(outputFile *string, optionsStrings *[]string) {
	if *outputFile != "" {
		output, err := os.Create(*outputFile)
		if err != nil {
			log.Fatalf("Error while create stdin #{err}")
			return
		}

		_, err = output.WriteString(strings.Join(*optionsStrings, "\n"))
		if err != nil {
			log.Fatalf("Error while write stdin #{err}")
			return
		}
	} else {
		printArray(optionsStrings)
	}
}

func main() {
	arguments, inputFile, outputFile := parseArguments()

	sliceArg := arguments[0:3]
	if checkExtraOptions(&sliceArg) {
		fmt.Println("Format: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		return
	}

	if inputFile == "" {
		input := bufio.NewScanner(os.Stdin)
		stringsText := make([]string, 0)
		for input.Scan() {
			text := input.Text()
			stringsText = append(stringsText, text)
		}

		if err := input.Err(); err != nil {
			log.Fatalf("Error while reading stdin #{err}")
			return
		}

		optionsStrings := correctUniqWork(&arguments, &stringsText)

		checkOutputFile(&outputFile, &optionsStrings)
	} else {
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Println(err)
		}

		text := string(data)

		stringsText := strings.Split(text, "\n")

		optionsStrings := correctUniqWork(&arguments, &stringsText)

		checkOutputFile(&outputFile, &optionsStrings)
	}
}