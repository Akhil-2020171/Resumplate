package main

// Experience is a struct that holds the data for the experience section
type Experience struct {
	Company    string
	Position   string
	Website    string
	StartDate  string
	EndDate    string
	Highlights []string
}

// Education is a struct that holds the data for the education section
type Education struct {
	Institution string
	Area        string
	StudyType   string
	StartDate   string
	EndDate     string
	GPA         string
}

// Skill is a struct that holds the data for the skills section
type Skill struct {
	Name     string
	Level    string
	Keywords []string
}

// Project is a struct that holds the data for the projects section
type Project struct {
	Name        string
	Description string
	Link        string
	StartDate   string
	EndDate     string
}

// Award is a struct that holds the data for the awards section
type Award struct {
	Title   string
	Date    string
	Awarder string
	Summary string
}

// Certificate is a struct that holds the data for the certificates section
type Certificate struct {
	Title   string
	Date    string
	License string
}

// Language is a struct that holds the data for the languages section
type Language struct {
	Language string
	Fluency  string
}

// UserData is a struct that holds the data for the resume
type UserData struct {
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Website      string
	Location     string
	Profile      string
	Experience   []Experience
	Education    []Education
	Skills       []Skill
	Projects     []Project
	Awards       []Award
	Certificates []Certificate
	Languages    []Language
}