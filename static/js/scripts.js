function addExperience() {
    const experienceFields = document.getElementById("my-experiences");
    const newEntry = document.createElement("div");
    newEntry.className = "experience-entry";
    newEntry.innerHTML = `
        <label for="Company">Company:</label>
        <input type="text" name="Company" required>
        <label for="Position">Position:</label>
        <input type="text" name="Position" required>
        <label for="Website">Website:</label>
        <input type="url" name="Website">
        <label for="Experience-StartDate">Start Date:</label>
        <input type="date" name="Experience-StartDate" required>
        <label for="Experience-EndDate">End Date:</label>
        <input type="date" name="Experience-EndDate" required>
        <label for="Highlights">Highlights:</label>
        <textarea name="Highlights" required></textarea>
        <button type="button" onclick="this.parentElement.remove()">Remove Experience</button>
    `;
    experienceFields.appendChild(newEntry);
}

function addEducation() {
    const educationFields = document.getElementById("my-educations");
    const newEntry = document.createElement("div");
    newEntry.className = "education-entry";
    newEntry.innerHTML = `
        <label for="Institution">Institution:</label>
        <input type="text" name="Institution" required>
        <label for="Area">Area:</label>
        <input type="text" name="Area" required>
        <label for="StudyType">Study Type:</label>
        <select name="StudyType" required>
            <option value="Bachelor">Bachelor</option>
            <option value="Master">Master</option>
            <option value="Doctorate">Doctorate</option>
        </select>
        <label for="Education-StartDate">Start Date:</label>
        <input type="date" name="Education-StartDate" required>
        <label for="Education-EndDate">End Date:</label>
        <input type="date" name="Education-EndDate" required>
        <label for="GPA">GPA:</label>
        <input type="number" name="GPA" required>
        <button type="button" onclick="this.parentElement.remove()">Remove Education</button>
    `;
    educationFields.appendChild(newEntry);
}

function addSkill() {
    const skillFields = document.getElementById("my-skills");
    const newEntry = document.createElement("div");
    newEntry.className = "skill-entry";
    newEntry.innerHTML = `
        <label for="Skill-Name">Skill:</label>
        <input type="text" name="Skill" required>
        <label for="Level">Level:</label>
        <select name="Level" required>
            <option value="Beginner">Beginner</option>
            <option value="Intermediate">Intermediate</option>
            <option value="Advanced">Advanced</option>
        </select>
        <label for="Keywords">Keywords:</label>
        <input type="text" name="Keywords" required>
        <button type="button" onclick="this.parentElement.remove()">Remove Skill</button>
    `;
    skillFields.appendChild(newEntry);
}

function addLanguage() {
    const languageFields = document.getElementById("my-languages");
    const newEntry = document.createElement("div");
    newEntry.className = "language-entry";
    newEntry.innerHTML = `
        <label for="Language">Language:</label>
        <input type="text" name="Language" required>
        <label for="Fluency">Fluency:</label>
        <select name="Fluency" required>
            <option value="Elementary">Elementary</option>
            <option value="Limited Working">Limited Working</option>
            <option value="Professional Working">Professional Working</option>
            <option value="Full Professional">Full Professional</option>
            <option value="Native Proficiency">Native Proficiency</option>
        </select>
        <button type="button" onclick="this.parentElement.remove()">Remove Language</button>
    `;
    languageFields.appendChild(newEntry);
}

function addProject() {
    const projectFields = document.getElementById("my-projects");
    const newEntry = document.createElement("div");
    newEntry.className = "project-entry";
    newEntry.innerHTML = `
        <label for="projectName">Project Name:</label>
        <input type="text" name="projectName" required>
        <label for="projectDescription">Project Description:</label>
        <textarea name="projectDescription" required></textarea>
        <label for="projectLink">Project Link:</label>
        <input type="url" name="projectLink">
        <label for="Project-StartDate">Start Date:</label>
        <input type="date" name="Project-StartDate" required>
        <label for="Project-EndDate">End Date:</label>
        <input type="date" name="Project-EndDate" required>
        <button type="button" onclick="this.parentElement.remove()">Remove Project</button>
    `;
    projectFields.appendChild(newEntry);
}

function addCertification() {
    const certificationFields = document.getElementById("my-certifications");
    const newEntry = document.createElement("div");
    newEntry.className = "certification-entry";
    newEntry.innerHTML = `
        <label for="certificationName">Certification Name:</label>
        <input type="text" name="certificationName" required>
        <label for="certificationDate">Certification Date:</label>
        <input type="date" name="certificationDate" required>
        <label for="certificationLink">Certification Link:</label>
        <input type="url" name="certificationLink">
        <button type="button" onclick="this.parentElement.remove()">Remove Certification</button>
    `;
    certificationFields.appendChild(newEntry);
}

function addAward() {
    const awardFields = document.getElementById("my-awrads");
    const newEntry = document.createElement("div");
    newEntry.className = "award-entry";
    newEntry.innerHTML = `
        <label for="awardName">Award Name:</label>
        <input type="text" name="awardName" required>
        <label for="awardDate">Award Date:</label>
        <input type="date" name="awardDate" required>
        <label for="awardLink">Award Link:</label>
        <input type="url" name="awardLink">
        <label for="summary">Summary:</label>
        <textarea name="summary" required></textarea>
        <button type="button" onclick="this.parentElement.remove()">Remove Award</button>
    `;
    awardFields.appendChild(newEntry);
}


$(document).ready(function () {
    (async () => {
        try {
            // load the resume data from local storage
            const resumeData = localStorage.getItem('resumeData');
            if (resumeData) {
                console.log('Resume data found in local storage:', resumeData);
                const data = JSON.parse(resumeData);
                console.log('Data:', data);
                Object.entries(data).forEach(([key, value]) => {
                    const element = document.querySelector(`[name="${key}"]`);
                    if (element) {
                        element.value = value;
                    }
                });
            } else {
                console.log('No resume data found in local storage.');
            }

            // Fetch available templates
            console.log('Fetching available templates...');
            const response = await fetch('/availableTemplates', {
                headers: {
                    'Accept': 'application/json' // Ensure the Accept header is set to JSON
                }
            });
            console.log('Response:', response);

            // Check if the response is not OK (e.g., 404, 500)
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            console.log('Response is OK');

            const templates = await response.json();
            console.log('Templates:', templates);

            if (templates.length === 0) {
                Swal.fire('No Templates Available', 'Please contact support.', 'error');
                return;
            }

            const templateButtons = templates.reduce((acc, template) => {
                acc[template] = template;
                return acc;
            }, {});

            Swal.fire({
                title: 'Select Template',
                text: 'Choose a template for your resume preview.',
                input: 'select',
                inputOptions: templateButtons,
                showCancelButton: true,
                inputPlaceholder: 'Select a template',
            }).then((result) => {
                if (result.isConfirmed && result.value) {
                    document.getElementById('templateType').value = result.value;
                    // document.getElementById('resumeForm').submit();
                } else if (!result.value) {
                    Swal.fire('Invalid Selection', 'Please choose a valid template.', 'error');
                }
            });
        } catch (error) {
            Swal.fire('Error', 'Failed to load templates. Please try again.', 'error');
            console.error('Error fetching templates:', error);
        }
    })(); // Immediately invoke the async function
});

function submitForm(event) {
    event.preventDefault();

    const resumeForm = document.getElementById('resumeForm');
    const formData = new FormData(resumeForm);
    const data = Object.fromEntries(formData.entries());

    try {
        // Save to local storage as before
        localStorage.setItem('resumeData', JSON.stringify(data));
        //Swal.fire('Data Saved', 'Your resume data has been saved.', 'success');
    } catch (error) {
        console.error('Error saving resume data to local storage:', error);
        Swal.fire('Error', 'Failed to save resume data. Please try again.', 'error');
        return;
    }

    fetch('/submitResume', {
        method: 'POST',
        body: formData
    })
        .then(response => response.text())
        .then(html => {
            // Open a new popup window with specific size
            const printWindow = window.open('', '_blank', 'width=850,height=1100');  // Popup with A4 size in pixels

            const csrfTokenElement = document.querySelector('input[name="gorilla.csrf.Token"]');
            const csrfToken = csrfTokenElement ? csrfTokenElement.value : null;
            console.log('CSRF Token:', csrfToken);

            // Fetch iframe_layout.html and inject it into the popup
            fetch('/layout/printLayout.html?open=true', {
                method: 'GET',
                headers: {
                    'Content-Type': 'text/html',
                    'X-CSRF-Token': csrfToken, // Include the CSRF token in the headers
                    'X-Requested-With': 'XMLHttpRequest'
                }
            })
                .then(res => res.text())
                .then(layoutHtml => {
                    //alert('Layout HTML:', layoutHtml);
                    // Write the layout HTML into the popup window
                    printWindow.document.write(layoutHtml); // Write the layout HTML into the popup
                    printWindow.document.close();  // Close the document for rendering

                    // Dynamically get the templateType value from the hidden input element
                    const templateTypeElement = document.getElementById('templateType').value;
                    
                    // remove the .html extension from the templateType
                    const templateType = templateTypeElement.replace('.html', '');

                    // Construct the CSS file path based on templateType
                    const cssFilePath = `../css/${templateType}.css`;

                    // Create a <link> element for the CSS file
                    const linkElement = printWindow.document.createElement('link');
                    linkElement.rel = 'stylesheet';
                    linkElement.href = cssFilePath;
                    linkElement.type = 'text/css';

                    // Append the <link> element to the <head> of the popup window
                    const head = printWindow.document.head || printWindow.document.getElementsByTagName('head')[0];
                    if (head) {
                        head.appendChild(linkElement);
                        console.log(`CSS dynamically added to print window: ${cssFilePath}`);
                    } else {
                        console.error('Unable to find <head> in print window.');
                    }                    

                    // Inject the resume HTML into the div inside the popup window
                    const resumeContentDiv = printWindow.document.getElementById('resumeContent');
                    if (resumeContentDiv) {
                        resumeContentDiv.innerHTML = html; // Inject the resume HTML directly
                        console.log("Resume content loaded into div.");
                    } else {
                        console.error('Resume content div not found in popup window.');
                    }
                })
                .catch(error => console.error("Error loading iframe_layout.html:", error));
        })
        .catch(error => console.error("Error:", error));
}