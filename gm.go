package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getFileList(path string) (map[int]string, error) {
	path = getCurrentDirectory(path)
	//pathSep := string(os.PathSeparator)
	pathSep := "/"
	files := make(map[int]string)
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	i := 1
	for _, file := range fs {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), "_test.go") {
			files[i] = path + pathSep + file.Name()
			i++
		}
	}
	return files, nil
}
func getCurrentDirectory(path string) string {
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
	//return dir
}
func main() {
	// path 是路径chdir,testType是测试文件1还是方法2
	var path string
	var no int
	argsCount := len(os.Args)
	if argsCount >= 2 {
		path = os.Args[1]
	} else {
		fmt.Println("please input like 'gm chdir' test again!")
		return
	}
	files, err := getFileList(path)
	if err != nil {
		panic(err)
	}
	// 测试方法
	// 提取文件中的所有方法
	fileMethodMap := getFileMethodMap(files)
	sortedKeys := getSortedKeys(fileMethodMap)
	for _,key := range sortedKeys {
		value := fileMethodMap[key]
		fmt.Println(strconv.Itoa(key) + " ---- " + value[strings.LastIndex(value, "/")+1:])
	}
	fmt.Println("请选择要执行的编号：")
	fmt.Scanln(&no)
	if fileMethodMap[no] == ""{
		fmt.Println("please enter right No.!")
		fmt.Scanln(&no)
	}
	selectFileMethod := fileMethodMap[no]
	fmt.Println("_----------------" + selectFileMethod)
	fm := strings.Split(selectFileMethod, "----")
	args := fmt.Sprintf("test -v %s -test.run %s", fm[0], fm[1])
	execCmd("go", args)

}

func getSortedKeys(files map[int]string) []int {
	var sortedKeys []int
	for key := range files {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Ints(sortedKeys)
	return sortedKeys
}

func getFileMethodMap(filesMap map[int]string) map[int]string {
	fileMethodMap := make(map[int]string)
	regexp1 := regexp.MustCompile(`\b(Test\w*)\(`)
	i := 1
	for _, file := range filesMap {
		fileByte, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		fileContent := string(fileByte)
		// 匹配提取
		matches := regexp1.FindAllStringSubmatch(fileContent, -1)
		for _, match := range matches {
			value := fmt.Sprintf("%s----%s", file, strings.TrimSpace(match[1]))
			fileMethodMap[i] = value
			i++
		}
	}
	return fileMethodMap
}

// golang command使用go，args=test -v file.go
func execCmd(command, args string) {
	//command := "go"
	//args := "test -v E:/go3/go-test-tool/logfactory_test.go"
	argArray := strings.Split(args, " ")
	cmd := exec.Command(command, argArray...)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, command, args)
		return
	}
	fmt.Fprintf(os.Stdout, "%s", buf)
}
