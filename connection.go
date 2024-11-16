package main

import (
	"log"
	"net/http"
)

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

    templateType := r.FormValue("templateType")
    if !isTemplateAvailable(templateType) {
        http.Error(w, "Selected template is not available. Please choose another.", http.StatusBadRequest)
        return
    }

	FirstName := r.FormValue("FirstName")
	LastName := r.FormValue("LastName")
	Email := r.FormValue("Email")
	Phone := r.FormValue("Phone")
	Website := r.FormValue("PersonalWebsite")
	Address := r.FormValue("Address")
	Profile := r.FormValue("Profile")

    userData := UserData{
		FirstName: FirstName,
		LastName: LastName,
		Email: Email,
		Phone: Phone,
		Website: Website,
		Location: Address,
		Profile: Profile,
    }

    // Handle multiple experience entries
    companies := r.Form["Company"]
    positions := r.Form["Position"]
    startDates := r.Form["Experience-StartDate"]
    endDates := r.Form["Experience-EndDate"]
    highlights := r.Form["Highlights"]

    for i := range companies {
        exp := Experience{
            Company:   companies[i],
            Position:  positions[i],
            StartDate: startDates[i],
            EndDate:   endDates[i],
            Highlights: []string{highlights[i]}, // Handle multiple highlights similarly
        }
        userData.Experience = append(userData.Experience, exp)
    }

	// Handle multiple education entries similarly
	Institutions := r.Form["Institution"]
	Degree := r.Form["Area"]
	StudyType := r.Form["StudyType"]
	StartDates := r.Form["Education-StartDate"]
	EndDates := r.Form["Education-EndDate"]
	GPA := r.Form["GPA"]

	for i := range Institutions {
		edu := Education{
			Institution: Institutions[i],
			Area: Degree[i],
			StudyType: StudyType[i],
			StartDate: StartDates[i],
			EndDate: EndDates[i],
			GPA: GPA[i],
		}
		userData.Education = append(userData.Education, edu)
	}

	// Handle multiple skill entries similarly
	Skills := r.Form["Skill"]
	Levels := r.Form["Level"]
	Keywords := r.Form["Keywords"]

	for i := range Skills {
		skill := Skill{
			Name: Skills[i],
			Level: Levels[i],
			Keywords: []string{Keywords[i]}, // Handle multiple keywords similarly
		}
		userData.Skills = append(userData.Skills, skill)
	}

	// Handle multiple project entries similarly	
	Projects := r.Form["projectName"]
	Descriptions := r.Form["projectDescription"]
	Links := r.Form["projectLink"]
	ProjectStartDates := r.Form["Project-StartDate"]
	ProjectEndDates := r.Form["Project-EndDate"]

	for i := range Projects {
		project := Project{
			Name: Projects[i],
			Description: Descriptions[i],
			Link: Links[i],
			StartDate: ProjectStartDates[i],
			EndDate: ProjectEndDates[i],
		}
		userData.Projects = append(userData.Projects, project)
	}

	// Handle multiple award entries similarly
	Awards := r.Form["awardName"]
	AwardDates := r.Form["awardDate"]
	AwardLinks := r.Form["awardLink"]
	AwardSummaries := r.Form["summary"]

	for i := range Awards {
		award := Award{
			Title: Awards[i],
			Date: AwardDates[i],
			Awarder: AwardLinks[i],
			Summary: AwardSummaries[i],
		}
		userData.Awards = append(userData.Awards, award)
	}

	// Handle multiple certificate entries similarly
	Certificates := r.Form["certificationName"]
	CertificateDates := r.Form["certificationDate"]
	CertificateLicenses := r.Form["certificationLink"]

	for i := range Certificates {
		certificate := Certificate{
			Title: Certificates[i],
			Date: CertificateDates[i],
			License: CertificateLicenses[i],
		}
		userData.Certificates = append(userData.Certificates, certificate)
	}

	// Handle multiple language entries similarly
	Languages := r.Form["Language"]
	Fluencies := r.Form["Fluency"]

	for i := range Languages {
		language := Language{
			Language: Languages[i],
			Fluency: Fluencies[i],
		}
		userData.Languages = append(userData.Languages, language)
	}

	log.Printf("Received user data: %+v", userData)

    // Call previewResumeTemplate to generate HTML preview
    err = previewResumeTemplate(&userData, templateType, w)
    if err != nil {
        log.Printf("Error generating template: %v", err)
        http.Error(w, "Error generating template", http.StatusInternalServerError)
        return
    }
}