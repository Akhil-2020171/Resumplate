package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"github.com/gorilla/csrf"
)

// Global CSRF protection key
var csrfKey = []byte("AXEFQ3J45TOP6912WMHHFKDV63DPIN7S") // 32 bytes 

// Template rendering function
var tmpl *template.Template

var (
    AvailableTemplates []string
    mu                 sync.Mutex
)

func init() {
    // Parse your resume.html template once and reuse it.
    tmpl = template.Must(template.ParseFiles("./static/resume.html"))
}


func main() {
	// Set up the CSRF protection middleware with the secret key
    CSRF := csrf.Protect(csrfKey,
        csrf.Secure(false),                             // Set to true in production (for HTTPS)
        csrf.HttpOnly(true),                            // Makes CSRF token cookie HTTP-only
        csrf.Path("/"),                                 // Cookie path
        csrf.MaxAge(3600),                              // Set CSRF cookie expiry to 1 hour
        csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            http.Error(w, "Forbidden - CSRF token invalid or missing", http.StatusForbidden)
        })),
    )
        
    // Define API routes first
    http.HandleFunc("/resume", renderResumeForm)
    
    // Serve static files last
    http.HandleFunc("/", serveFile)    

    http.HandleFunc("/submitResume", handleFormSubmission)
    http.HandleFunc("/availableTemplates", fetchAvailableTemplates) 

    log.Println("Server started on :8080")
    errs := http.ListenAndServe(":8080", CSRF(http.DefaultServeMux))
    if errs != nil {
        log.Fatal("Listen-And-Serve: ", errs)
    }   

    log.Println("For Application, go to http://localhost:8080/resume")   
}

// Render the resume form with the CSRF token embedded
func renderResumeForm(w http.ResponseWriter, r *http.Request) {
    // Pass CSRF token to template
    data := map[string]interface{}{
        "csrfField": csrf.TemplateField(r),
    }

    err := tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
    }
}

// Custom file server to handle requests
func serveFile(w http.ResponseWriter, r *http.Request) {
    log.Printf("Request Path: %s", r.URL.Path) // Log the request path

    // Check if the request is for an HTML file
    if strings.HasSuffix(r.URL.Path, ".html") {
        // Remove the .html extension and redirect
        openQuery := r.URL.Query().Get("open") // Get the 'open' query parameter
        if openQuery == "true" && r.Header.Get("X-Requested-With") == "XMLHttpRequest"{
            if strings.HasPrefix(r.URL.Path, "/layout/"){
                filePath := path.Join("./static", r.URL.Path)
                log.Printf("Checking file path: %s", filePath) // Log the resolved file path
    
                // Check if the requested layout file exists
                if _, err := os.Stat(filePath); os.IsNotExist(err) {
                    log.Printf("File not found: %s", filePath)
                    http.ServeFile(w, r, "./static/404.html")
                    return
                }
                // Serve the file if it exists
                http.ServeFile(w, r, filePath)
                return
            }
        }

        newPath := strings.TrimSuffix(r.URL.Path, ".html")
        http.Redirect(w, r, newPath, http.StatusMovedPermanently)
        return
    }  

    // Prevent directory listing
    if r.URL.Path == "/" {
        http.ServeFile(w, r, "./static/NotAllowed.html")
        return
    }

    // Serve files without revealing the directory structure
    filePath := path.Join("./static", r.URL.Path)

    // Use os.Stat to check if the requested file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        log.Printf("File not found: %s", filePath) // Log if file is not found
        http.ServeFile(w, r, "./static/404.html") // Ensure the path matches where your 404.html is located
        return
    }

    http.ServeFile(w, r, filePath)  
}

// Fetch available templates from a text file
func fetchAvailableTemplates(w http.ResponseWriter, r *http.Request) {
    // Check if the request is likely coming from a browser
    if strings.Contains(r.Header.Get("Accept"), "text/html") {
        http.ServeFile(w, r, "./static/404.html") // Serve your custom 404 page
        return
    }

    w.Header().Set("Content-Type", "application/json")

    // Locking to avoid race conditions during concurrent access
    mu.Lock()
    defer mu.Unlock()

    // Debug: Log incoming request
    log.Printf("Received request for available templates")

    if len(AvailableTemplates) == 0 {
        var err error
        AvailableTemplates, err = LoadAvailableTemplates("./static/templates")
        if err != nil {
            log.Printf("Error loading templates: %v", err)
            http.Error(w, "Failed to load templates", http.StatusInternalServerError)
            return
        }
    }

    // Send the templates as JSON response
    if err := json.NewEncoder(w).Encode(AvailableTemplates); err != nil {
        log.Printf("Error encoding templates: %v", err)
        http.Error(w, "Failed to encode templates", http.StatusInternalServerError)
    }
}