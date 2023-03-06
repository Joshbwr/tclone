package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/otiai10/copy"
)

var TEMPLATE_PATH = "/Users/joshuabrewer/Documents/dev/starter-templates"

func getFoldersInDirectory() ([]fs.DirEntry, error) {

	folders, err := os.ReadDir(TEMPLATE_PATH)
	if err != nil {
		return nil, fmt.Errorf("Error reading directory: %w", err)
	}

	return folders, nil

}

type TemplateOption struct {
	Name string
}


func buildProjectSelection(folders []fs.DirEntry) (string, error){
	options := []TemplateOption{}

	for _, folder := range folders {
		options = append(options,  TemplateOption{Name: folder.Name()})
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 {{ .Name | cyan | bold }}",
		Inactive: "   {{ .Name | white }}",
		Selected: "{{.Name | green | bold }}: {{ . | cyan }}",
		Details:  "",
	}

	prompt := promptui.Select{
		Label:     "Select Directory to Clone",
		Items:     options,
		Templates: templates,
		Size:      4,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return "", errors.New("Config Prompt Failure")
	}

	return options[i].Name, nil

}

func namePrompt() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter a name for your new project",
		Validate: func(input string) error {
			if len(input) < 3 {
				return errors.New("Name must be at least 3 characters")
			}
			return nil
		},
	}

	result, err := prompt.Run()

	if err != nil {
		return "", errors.New("Name Prompt Failure")
	}

	return result, nil
}

func copyDirectory(newProjectName string, selectedTemplate string) {
	// get the current directory path that the user is in
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: %w", err)
	}

	// copy the selected template to the new project name and current directory
	err = copy.Copy(TEMPLATE_PATH + "/" + selectedTemplate, wd + "/" + newProjectName)
	if err != nil {
		fmt.Println("Error copying directory: %w", err)
	}
}


func Init(){
	// get all folders in the templates directory
	folders, err := getFoldersInDirectory()
	if err != nil {
		fmt.Println("Error reading directory: %w", err)
		os.Exit(500)
	}

	// build prompt options and prompt user to select a template
	selectedTemplate, err := buildProjectSelection(folders)
	if err != nil {
		fmt.Println("Error selecting project template: %w", err)
		os.Exit(500)
	}

	// prompt user to enter a name for the new project
	newProjectName, err := namePrompt()
	if err != nil {
		fmt.Println("Error getting project name: %w", err)
		os.Exit(500)
	}
	
	// clone the selected template to the new project name and current directory
	copyDirectory(newProjectName, selectedTemplate)
}

