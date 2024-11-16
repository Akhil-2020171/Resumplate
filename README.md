
# Resumplate

**Resumplate** is a fast, simple, and responsive resume template builder developed using HTML, CSS, and JavaScript. It allows users to create, customize, and download professional resumes with ease.

## 🚀 Features
- **Fast & Responsive**: Optimized for all devices, ensuring a smooth experience across desktops, tablets, and smartphones.
- **User-Friendly**: Simple interface that makes resume building intuitive for everyone.
- **Clean Design**: Professional and minimal layout to make your resume stand out.
- **Download as PDF**: Easily export your resume as a PDF for sharing or printing.

## 🛠️ How to Use
1. **Clone the Repository**:  
   ```bash
   git clone https://github.com/Akhil-2020171/resumplate.git
   ```

2. **Create Your Resume Template**:  
   - Design your resume content in HTML format. Make sure to include only the `<div>` elements and exclude `<head>` and `<body>` tags.
   
3. **Provide a CSS File**:  
   - Name the CSS file the same as the HTML file. For example, if your HTML template is `resume.html`, the CSS file should be `resume.css`.

4. **Organize Files**:  
   - Place the `.html` file in the `./static/templates/` folder.
   - Place the corresponding `.css` file in the `./static/css/` folder.

## ▶️ How to Run
1. **Run Using Go**:  
   Simply execute the command:  
   ```bash
   go run .
   ```  
   Then open your browser and navigate to:  
   `http://localhost:8080`

   **OR**

2. **Build and Run**:  
   - Build the project:  
     ```bash
     go build .
     ```  
   - Run the executable:  
     ```bash
     ./resumplate
     ```  
   - Open your browser and go to:  
     `http://localhost:8080`

## 💡 Tips
- Ensure your HTML and CSS files follow the correct naming convention.
- You can customize the templates further by editing the CSS to match your style preferences.
