package main

import (
	"bufio"
	"fmt"
	"os"
)

// find all env files
// add option to config new files
func main() {
	// fmt.Println("starting..")
	// input, err := ioutil.ReadFile("env.sh")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// lines := strings.Split(string(input), "\n")

	// for i, line := range lines {
	// 	tmp := strings.Split(line, "=")
	// 	if len(tmp[0]) > 0 && tmp != nil {
	// 		fmt.Println(tmp[0])
	// 		lines[i] = tmp[0] + "=\"\""
	// 	}
	// }
	// output := strings.Join(lines, "\n")
	// fmt.Println(output)
	//cmd.Execute()

	// err = ioutil.WriteFile("env2.sh", []byte(output), 0644)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your name is: %s\n", text)
}
