package main

import (
	"os"
	"fmt"
	"github.com/harhogefoo/go_training2017/ch04/ex11/github"
	"flag"
	"io/ioutil"
	"os/exec"
)

var issueNo = flag.Int("issue", 0, "issue number")
var title = flag.String("title", "", "Issueのタイトル")
var body = flag.String("body", "", "Issueの内容")

var createFlag = flag.Bool("create", false, "Issueの作成")
var closeFlag = flag.Bool("close", false, "Issueのクローズ")
var editFlag = flag.Bool("edit", false, "Issueの編集")
var printFlag = flag.Bool("print", false, "Issueの表示")

var repository string

func main() {
	flag.Parse()
	fmt.Println(flag.Arg(0))
	validateOperationFlags()
	repository = validateAndGetRepository()

	var user github.Credentials
	user.Query()
	fmt.Println()

	switch true {
	case *createFlag:
		b := *body
		if !isFlagSpecified("body") {
			b = invokeEditor()
		}
		issue, err := github.Create(repository, *title, b, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%+v", issue)
		saveIssueNo(issue.Number)
	case *closeFlag:
		if !isFlagSpecified("issue") {
			fmt.Print("Please specify -issue <issueNo>")
			os.Exit(1)
		}
		issue, err := github.Close(repository, *issueNo, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%+v\n", issue)
	case *printFlag:
		if !isFlagSpecified("issue") {
			fmt.Print("Please specify -issue <issueNo>")
			os.Exit(1)
		}
		issue, err := github.Get(repository, *issueNo, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%+v\n", issue)
	case *editFlag:
		if !isFlagSpecified("issue") {
			fmt.Print("Please specify -issue <issueNo>")
			os.Exit(1)
		}
		b := *body
		if !isFlagSpecified("body") {
			b = invokeEditor()
		}
		issue, err := github.Edit(repository, *title, b, *issueNo, &user)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", issue)
	}

}

func saveIssueNo(issueNo int) {
	f, err := os.Create("issue_no.txt")
	if err != nil {
		return
	}
	fmt.Fprintf(f, "%d", issueNo)
	f.Close()
}

func invokeEditor() string {
	f, err := ioutil.TempFile("", "body.")
	if err != nil {
		panic(fmt.Errorf("Cannot crete a temp file: %v", err))
	}

	name := f.Name()
	f.Close() // ignore Error

	cmd := exec.Command("vim", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(fmt.Errorf("Cannot invoke vim: %v", err))
	}

	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Errorf("Cannot read a temp file: %v", err))
	}
	return string(bytes)
}

func isFlagSpecified(name string) (specified bool) {
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			specified = true
		}
	})
	return
}

func validateOperationFlags() {
	flags := []bool{*createFlag, *closeFlag, *editFlag, *printFlag}

	trueCount := 0
	for _, flag := range flags {
		if flag {
			trueCount++
		}
	}

	if trueCount == 0 {
		fmt.Println("Operation flag(-create, -close, -edit, -print) is not specified")
		os.Exit(1)
	}
	if trueCount >= 2 {
		fmt.Printf("Too many operation flags are specified: %#v\n", flags)
	}
}

func validateAndGetRepository() string {
	if flag.NArg() == 0 {
		fmt.Println("REPOSITORY is not specified")
		os.Exit(1)
	}
	if flag.NArg() != 1 {
		fmt.Println("Too Many argument")
		os.Exit(1)
	}
	return flag.Arg(0)
}

