package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func readLines(filepath string) ([]string, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}
	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func coveringJournals(cwd string) (files []string) {
	files = []string{}

	log.Printf("gathering journals from %q\n", cwd)
	dirs := parentPaths(cwd, []string{})[1:]
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			filename := entry.Name()
			ok := filename == "ledgerrc"
			ok = ok || filename == ".ledgerrc"
			ok = ok || path.Ext(filename) == ".ldg"
			if ok {
				includeFile := path.Join(dir, filename)
				log.Printf("auto including %q\n", includeFile)
				files = append(files, includeFile)
			}
		}
	}
	return files
}

func listJournals(cwd string) ([]string, error) {
	var files []string
	items, err := os.ReadDir(cwd)
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}
	for _, item := range items {
		if path.Ext(item.Name()) == ".ldg" {
			files = append(files, filepath.Join(cwd, item.Name()))
		}
	}
	return files, nil
}

func findJournals(cwd string) ([]string, error) {
	var files []string
	_ = filepath.Walk(
		cwd,
		func(pathDir string, info os.FileInfo, err error) error {
			if path.Ext(info.Name()) == ".ldg" {
				files = append(files, filepath.Join(pathDir, info.Name()))
			}
			return nil
		},
	)
	return files, nil
}

func parentPaths(dirPath string, acc []string) (dirs []string) {
	dir, _ := path.Split(dirPath)
	if dir != "" {
		acc = append(acc, dirPath)
		dir = strings.TrimRight(dir, string(os.PathSeparator))
		return parentPaths(dir, acc)
	}
	return acc
}
