package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Project struct {
	Name    string   `json:"name"`
	Streaks []Streak `json:"streaks"`
}
type Streak struct {
	Start       time.Time     `json:"start"`
	End         time.Time     `json:"end"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
}
type Work struct {
	Projects           []Project `json:"projects"`
	CurrentProjectName string    `json:"currentProjectName"`
}

var work Work

func readProjects() {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		//file not exists, create directory and file
		color.Yellow(directoryPath + " not exists, creating directory")
		_ = os.Mkdir(directoryPath, os.ModePerm)
		saveWork()
	}
	content := string(file)
	json.Unmarshal([]byte(content), &work)
}

func saveWork() {

	jsonProjects, err := json.Marshal(work)
	check(err)
	err = ioutil.WriteFile(filePath, jsonProjects, 0644)
	check(err)
}
func getProjectIByName(name string) int {
	for i, project := range work.Projects {
		if project.Name == name {
			return i
		}
	}
	return -1
}
func projectExist(name string) bool {
	for _, project := range work.Projects {
		if project.Name == name {
			return true
		}
	}
	return false
}
func newProject(name string) error {
	//first, check if the project name is not taken yet
	if projectExist(name) {
		color.Red("Project name: " + name + ", already exists")
		return errors.New("project name already exist")
	}
	var newProject Project
	newProject.Name = name
	work.Projects = append(work.Projects, newProject)
	return nil
}
func listProjects() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Listing projects")
	fmt.Println("")
	for k, project := range work.Projects {
		fmt.Println("project " + strconv.Itoa(k))
		fmt.Print("name: ")
		color.Blue(project.Name)
		for k2, streak := range project.Streaks {
			fmt.Println("	streak: " + strconv.Itoa(k2))
			fmt.Print("	Start:")
			fmt.Println(streak.Start)
			fmt.Print("	End:")
			fmt.Println(streak.End)
			fmt.Print("	Duration:")
			fmt.Println(streak.Duration)
		}
		fmt.Println("")
	}
}
