package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func basicMainFunc() string {
	template := "package main \n \nimport 'fmt' \n \nfunc main(){ \n \n} "
	return template
}

func main() {
	menuItems := []string{"Basic main function", "Hello world", "Basic server"}

	fmt.Println("\nHola, choose the boilerplate you want to start off with. Also if you have an editor of preference then specify that and we will open the code in it. Or else you can leave it there, in which case a file will be created in the current working directory with a name you specify. Happy Coding!\n")
	for i, item := range menuItems {
		fmt.Printf("%d) %s\n", i+1, item)
	}
	fmt.Println("\nYour choice: ")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	opt := input.Text()

	switch opt {
	case "1":
		fmt.Println("I will build you a template for ", menuItems[0], "\n")
		basicMainFunc()
	case "2":
		fmt.Println("I will build you a template for ", menuItems[1], "\n")
	case "3":
		fmt.Println("I will build you a template for ", menuItems[2], "\n")
	}

	fmt.Println("Do you want to")
	fmt.Println("1) Simply save the file")
	fmt.Println("2) Open it in my favorite code editor")
	input.Scan()
	opt = input.Text()

	switch opt {
	case "1":
		cmd := exec.Command("echo", "hey", "| basic.go")
		fmt.Println(cmd.Run())
	case "2":
		fmt.Println("Please enter the command to open your favourite editor ")
		input.Scan()
		editor := input.Text()
		cmd := exec.Command("sh", "-c", "echo %s >> basic.go", basicMainFunc())
		fmt.Println(cmd.Run())

		cmd = exec.Command(editor, "basic.go")
		fmt.Println(cmd.Run())

	}

}
