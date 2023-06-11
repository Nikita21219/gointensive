package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func parseArgs() ([]string, string) {
	dirFlag := flag.String("a", ".", "archive")
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		log.Fatalln("Error args")
	}
	return paths, *dirFlag
}

func getOutName(path, distDir string) string {
	pathSplit := strings.Split(path, "/")
	fileName := strings.Split(pathSplit[len(pathSplit)-1], ".")[0]

	if distDir[len(distDir)-1] == '/' {
		return fmt.Sprintf("%s%s_%d.tar.gz", distDir, fileName, time.Now().Unix())
	}
	return fmt.Sprintf("%s/%s_%d.tar.gz", distDir, fileName, time.Now().Unix())
}

func createTar(path, distDir string) error {
	// Init filename
	outName := getOutName(path, distDir)
	// Create file
	file, err := os.Create(outName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create gzip writer
	gw := gzip.NewWriter(file)
	defer gw.Close()
	// Create tar writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	inFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inFile.Close()

	info, err := inFile.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = path

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, inFile)
	if err != nil {
		return err
	}

	fmt.Printf("Archive %s created successfully\n", outName)
	return nil
}

func main() {
	paths, distDir := parseArgs()

	var wg sync.WaitGroup

	for _, path := range paths {
		path := path
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := createTar(path, distDir)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}()
	}

	wg.Wait()
}
