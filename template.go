package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// previewResumeTemplate generates a resume preview using the selected template
func previewResumeTemplate(data *UserData, templateType string, w http.ResponseWriter) error {
    templatePath := filepath.Join("./static/templates/", templateType)

    // Define a custom function map for template functions
    funcMap := template.FuncMap{
        "join": strings.Join,
    }

    // Parse the main template with custom functions
    tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).ParseFiles(templatePath)
    if err != nil {
        log.Printf("Error parsing template: %v", err)
        return err
    }

    w.Header().Set("Content-Type", "text/html")
    return tmpl.Execute(w, data) // Render directly to the response writer
}

// Helper function to check if a template is available
func isTemplateAvailable(templateType string) bool {
    for _, tmpl := range AvailableTemplates {
        if tmpl == templateType {
            return true
        }
    }
    return false
}

// LoadAvailableTemplates scans a directory for available HTML templates
func LoadAvailableTemplates(dirPath string) ([]string, error) {
    // Open the directory
    files, err := os.ReadDir(dirPath)
    if err != nil {
        return nil, err
    }

    var templates []string
    for _, file := range files {
        // Only consider files with .html extension
        if !file.IsDir() && filepath.Ext(file.Name()) == ".html" {
            // Append the base name (without extension) or full name as needed
            templates = append(templates, file.Name())
        }
    }

    return templates, nil
}