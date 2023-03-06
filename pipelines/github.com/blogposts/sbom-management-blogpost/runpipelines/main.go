package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	var (
		apiKey              string
		authURL             string
		resultsFile         string
		pipelineRunTemplate string
		runsDir             string
	)
	flag.StringVar(&apiKey, "apiKey", "", "dependency track api key")
	flag.StringVar(&authURL, "url", "", "dependency track base url")
	flag.StringVar(&resultsFile, "resultsFile", "", "path to the file where the mappings of repoName to dependency track project ID are stored")
	flag.StringVar(&pipelineRunTemplate, "template", "", "path to the templated pipelinerun.yaml")
	flag.StringVar(&runsDir, "output", "", "path to the dir where we should save the final pipelineruns.yaml")

	flag.Parse()
	template, err := os.ReadFile(pipelineRunTemplate)
	if err != nil {
		log.Fatal(err)
	}
	dat, err := os.Open(resultsFile)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		runForRepo := strings.Clone(string(template))
		line := fileScanner.Text()
		components := strings.Split(line, ":")
		repoURL := "https://github.com/" + components[0] + ".git"
		projectName := components[0]
		projectUUID := components[1]
		runForRepo = strings.ReplaceAll(runForRepo, "$repo-url", strings.TrimSpace(repoURL))
		runForRepo = strings.ReplaceAll(runForRepo, "$project-name", strings.TrimSpace(projectName))
		runForRepo = strings.ReplaceAll(runForRepo, "$project-uuid", strings.TrimSpace(projectUUID))
		runForRepo = strings.ReplaceAll(runForRepo, "$dt-token", apiKey)
		err := os.WriteFile(path.Join(runsDir, strings.ReplaceAll(projectName+".yaml", "/", "-")), []byte(runForRepo), 0644)
		if err != nil {
			panic(err)
		}
	}
	defer dat.Close()
}
