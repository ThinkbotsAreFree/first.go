package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"flag"
)

func main() {

	var (
		dataStack []string
		progStack []string		
	)
	variable := make(map[string]string)
	function := make(map[string]string)

	tron := flag.Bool("tron", false, "trace on")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	white := color.New(color.FgWhite)
	magenta := color.New(color.FgMagenta).Add(color.Underline)

	ask := func(prompt string) string {

		white.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", 1)
		return (text)
	}

	push := func(stack *[]string, item string) {

		*stack = append(*stack, item)
	}

	pop := func(stack *[]string) string {

		if (len(*stack) == 0) { return("") }

		result := (*stack)[len(*stack)-1]
		if len(result) > 1 {
			result = strings.Replace(result, "_", " ", -1)
		}
		*stack = (*stack)[:len(*stack)-1]
		return (result)
	}

	popProg := func() string {

		return (pop(&progStack))
	}

	popData := func() string {

		return (strings.Replace(pop(&dataStack), "'", "", 1))
	}

	execute := func(token string) {

		do := func() {
			source := strings.Split(pop(&dataStack), " ")
			for s := 0; s < len(source); s++ {
				push(&progStack, source[s])
			}
		}

		switch token {

		case "pop":
			_ = popData()
			_ = popProg()

		case "dup":
			item := popData()
			push(&dataStack, item)
			push(&dataStack, item)
			_ = popProg()

		case "swap":
			item1 := popData()
			item2 := popData()
			push(&dataStack, item1)
			push(&dataStack, item2)
			_ = popProg()

		case "popd":
			item1 := popData()
			_ = popData()
			push(&dataStack, item1)
			_ = popProg()

		case "popop":
			_ = popData()
			_ = popData()
			_ = popProg()

		case "dupd":
			item1 := popData()
			item2 := popData()
			push(&dataStack, item2)
			push(&dataStack, item2)
			push(&dataStack, item1)
			_ = popProg()

		case "swapd":
			item1 := popData()
			item2 := popData()
			item3 := popData()
			push(&dataStack, item2)
			push(&dataStack, item3)
			push(&dataStack, item1)
			_ = popProg()

		case "rolldown":
			item1 := popData()
			item2 := popData()
			item3 := popData()
			push(&dataStack, item2)
			push(&dataStack, item1)
			push(&dataStack, item3)
			_ = popProg()

		case "rollup":
			item1 := popData()
			item2 := popData()
			item3 := popData()
			push(&dataStack, item1)
			push(&dataStack, item3)
			push(&dataStack, item2)
			_ = popProg()

		case "choice":
			condition := popData()
			item2 := popData()
			item3 := popData()
			if condition == "1" {
				push(&dataStack, item2)
			} else {
				push(&dataStack, item3)
			}
			_ = popProg()

		case "output":
			white.Println(popData())
			_ = popProg()

		case "input":
			prompt := popData()
			push(&dataStack, ask(prompt+" "))
			_ = popProg()

		case "cons":
			item1 := popData()
			item2 := popData()
			push(&dataStack, item2+" "+item1)
			_ = popProg()

		case "uncons":
			aggregate := strings.Split(popData(), " ")
			for a := 0; a < len(aggregate); a++ {
				push(&dataStack, aggregate[a])
			}
			_ = popProg()

		case "append":
			push(&dataStack, popData()+popData())
			_ = popProg()

		case "remove":
			push(&dataStack, strings.Replace(popData(), popData(), "", 1))
			_ = popProg()

		case "replace":
			push(&dataStack, strings.Replace(popData(), popData(), popData(), 1))
			_ = popProg()

		case "removeall":
			push(&dataStack, strings.Replace(popData(), popData(), "", -1))
			_ = popProg()

		case "replaceall":
			push(&dataStack, strings.Replace(popData(), popData(), popData(), -1))
			_ = popProg()

		case "do":
			_ = popProg()
			do()

		case "get":
			_ = popProg()
			push(&dataStack, variable[pop(&dataStack)])

		case "set":
			variable[popData()] = popData()
			_ = popProg()

		case "def":
			function[popData()] = popData()
			_ = popProg()

		case "pick":
			where := len(dataStack) - len(popData()) - 1
			if where >= 0 && where < len(dataStack) {
				push(&dataStack, dataStack[where])
			}
			_ = popProg()

		case "if":
			condition := popData()
			_ = popProg()
			if condition == "1" {
				do()
			} else {
				_ = popData()
			}

		case "ife":
			condition := popData()
			_ = popProg()
			if condition == "1" {
				do()
				_ = popData()
			} else {
				_ = popData()
				do()
			}

		case "and":
			boolean1 := popData()
			boolean2 := popData()
			if boolean1 == "1" && boolean2 == "1" {
				push(&dataStack, "1")
			} else {
				push(&dataStack, "0")
			}
			_ = popProg()

		case "or":
			boolean1 := popData()
			boolean2 := popData()
			if boolean1 == "1" || boolean2 == "1" {
				push(&dataStack, "1")
			} else {
				push(&dataStack, "0")
			}
			_ = popProg()

		case "not":
			boolean1 := popData()
			if boolean1 != "1" {
				push(&dataStack, "1")
			} else {
				push(&dataStack, "0")
			}
			_ = popProg()

		case "equals":
			item1 := popData()
			item2 := popData()
			if item1 == item2 {
				push(&dataStack, "1")
			} else {
				push(&dataStack, "0")
			}
			_ = popProg()

		case "contains":
			item1 := popData()
			item2 := popData()
			if strings.Contains(item1, item2) {
				push(&dataStack, "1")
			} else {
				push(&dataStack, "0")
			}
			_ = popProg()

		case "prefix":
			item1 := popData()
			item2 := popData()
			if strings.HasPrefix(item1, item2) {
				push(&dataStack, "1")
			} else {
				push(&dataStack, "0")
			}
			_ = popProg()

		case "suffix":
			item1 := popData()
			item2 := popData()
			if strings.HasSuffix(item1, item2) {
				push(&dataStack, "1")
			} else {
				push(&dataStack, "0")
			}
			_ = popProg()

		case "quote":
			_ = popProg()
			item := pop(&dataStack)
			push(&progStack, "'"+item)

		case "nothing":
			_ = popProg()
			push(&dataStack, "")

		case "space":
			_ = popProg()
			push(&dataStack, " ")

		default:
			item := popProg()
			value, ok := function[item]
			if ok {
				push(&dataStack, value)
				do()
			} else {
				push(&dataStack, item)
			}
		}

	}


	interpret := func(command string) {

		tokens := strings.Split(command, " ")

		for _, token := range tokens {
			push(&progStack, token)
		}

		for len(progStack) > 0 {
			if (*tron) {
				magenta.Print("  dataStack", dataStack)
				fmt.Println(" " + progStack[len(progStack)-1])
			}
			execute(progStack[len(progStack)-1])
		}
	}

	text := ""
	for text != "end" {

		if (*tron) {
			magenta.Println("  dataStack", dataStack)
			magenta.Println("  functions", function)
			magenta.Println("  variables", variable)
		}
		text = ask("\n> ")
		interpret(text)

	}

}
