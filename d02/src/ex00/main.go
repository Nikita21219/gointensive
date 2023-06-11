package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseArgs() (string, map[string]bool, string) {
	if len(os.Args[1:]) < 1 {
		log.Fatalln("Error: Wrong number of arguments")
	}
	args := os.Args
	path := "."
	idx := len(args) - 1
	if args[idx] != "-f" && args[idx] != "-sl" && args[idx] != "-d" {
		path = args[idx]
		os.Args = args[:idx]
	}

	slFlag := flag.Bool("sl", false, "print symlinks")
	dirFlag := flag.Bool("d", false, "print dirs")
	fileFlag := flag.Bool("f", false, "print files")
	fileExtFlag := flag.String("ext", "", "use only with -f flag. print files with specified extension")
	flag.Parse()

	if flag.NArg() > 0 {
		log.Fatalln("Error: Wrong argument")
	}

	if !(*fileFlag) && *fileExtFlag != "" {
		log.Fatalln("-ext flag use only with -f")
	}

	flags := map[string]bool{
		"sl":    *slFlag,
		"dirs":  *dirFlag,
		"files": *fileFlag}

	return path, flags, *fileExtFlag
}

func printSymLinks(path string, info os.FileInfo) {
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		linkPath, err := os.Readlink(path)
		if err != nil {
			fmt.Printf("Error reading link %s\n", path)
		} else if _, err := os.Stat(linkPath); os.IsNotExist(err) {
			fmt.Printf("%s -> [broken]\n", path)
		} else {
			fmt.Printf("%s -> %s\n", path, linkPath)
		}
	}
}

func printDirs(path string, info os.FileInfo) {
	if info.IsDir() && info.Mode()&os.ModeSymlink != os.ModeSymlink {
		dir, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dir.Close()
		files, err := dir.ReadDir(-1)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !(len(files) == 1 && files[0].IsDir()) {
			fmt.Println(path)
		}
	}
}

func printFilesWithExt(path, ext string) {
	arr := strings.Split(path, ".")
	if len(arr) < 2 {
		return
	}
	if arr[len(arr)-1] == ext {
		fmt.Println(path)
	}
}

func printFiles(path string, ext string, info os.FileInfo) {
	if !info.IsDir() && info.Mode()&os.ModeSymlink != os.ModeSymlink {
		if ext == "" {
			fmt.Printf("%s\n", path)
		} else {
			printFilesWithExt(path, ext)
		}
	}
}

func walkFunc(flags map[string]bool, ext, path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if flags["sl"] {
		printSymLinks(path, info)
	}
	if flags["dirs"] {
		printDirs(path, info)
	}
	if flags["files"] {
		printFiles(path, ext, info)
	}
	return nil
}

func readDir(path, ext string, flags map[string]bool) {
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		return walkFunc(flags, ext, path, info, err)
	})
	if err != nil {
		log.Println(err)
	}
}

func main() {
	path, flags, ext := parseArgs()
	if !flags["sl"] && !flags["dirs"] && !flags["files"] {
		flags["sl"] = true
		flags["dirs"] = true
		flags["files"] = true
	}
	readDir(path, ext, flags)
}
