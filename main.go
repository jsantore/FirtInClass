package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ClassInfo struct {
	name       string
	instructor string
	numCredits int
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	classNames := make([]ClassInfo, 0)
	inputReader := bufio.NewReader(os.Stdin)
	for {
		nextClass := ClassInfo{}
		fmt.Print("What class are you taking this semester?:")
		className, err := inputReader.ReadString('\n')
		println(err)
		nextClass.name = className
		fmt.Print("How many Credits is that class")
		credits, err := inputReader.ReadString('\n')
		credits = strings.TrimSpace(credits)
		creditsNum, err := strconv.Atoi(credits)
		fmt.Print(err)
		nextClass.numCredits = creditsNum
		fmt.Print("Who is the instructor")
		instructor, err := inputReader.ReadString('\n')
		nextClass.instructor = instructor
		classNames = append(classNames, nextClass)
		fmt.Print("Do you have another class? (y/n)")
		answer, err := inputReader.ReadString('\n')
		if strings.TrimSpace(answer) == "n" {
			break
		}
	}
	//fmt.Print("What other class are you taking this semester?:")
	//className, err = inputReader.ReadString('\n')
	//classNames = append(classNames, className)
	//fmt.Print("What other class are you taking this semester?:")
	//className, err = inputReader.ReadString('\n')
	//classNames = append(classNames, className)
	//fmt.Print("What other class are you taking this semester?:")
	//className, err = inputReader.ReadString('\n')
	//classNames = append(classNames, className)
	fmt.Print(classNames)
}
