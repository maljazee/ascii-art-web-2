ASCII Art Generator

Description
------------
The ASCII Art Generator is a web application that converts user-inputted text into ASCII art. It allows users to choose from different ASCII art styles to render their text. This Go application serves an HTML form where users can input their text and select a banner style. The server processes the input, generates the ASCII art, and displays it back to the user.

Authors
--------
- Mohamed Saeed - mohasaeed
- Mohammed Aljazeeri - maljazee

Usage
------
To run the ASCII Art Generator, follow these steps:

1. Clone the repository:
   git clone https://github.com/yourusername/ascii-art-generator.git

2. Navigate to the project directory:
   cd ascii-art-generator

3. Prepare your environment:
   - Ensure you have Go installed (version 1.18 or later is recommended).

4. Run the server:
   go run main.go

5. Open your web browser and navigate to http://localhost:8080 to access the application.

6. Submit text and select a banner style on the web interface to see the ASCII art generated.

Implementation Details
-----------------------
Algorithm
1. Handling Requests:
   - The server listens for HTTP requests on port 8080. It serves an HTML form for input via the / route.
   - For POST requests, the server processes form data which includes the text to be converted and the selected ASCII art style (banner).

2. Generating ASCII Art:
   - The generateAsciiArt function takes the user input text and chosen banner style.
   - It reads the ASCII art style file corresponding to the selected banner.
   - Text is split into lines based on \n, and each line is converted to ASCII art using the specified banner style.

3. Reading ASCII Art Style Files:
   - The printLine function reads lines from the banner style file and extracts the appropriate lines for each character of the input text.
   - It constructs the ASCII art by iterating through each character and appending the appropriate lines from the style file.

4. Displaying Results:
   - The result is rendered back to the user using an HTML template (templates/index.html), showing the generated ASCII art.

Dependencies
-------------
- net/http: Standard library for HTTP server and client implementations.
- html/template: Standard library for HTML template parsing and execution.
- bufio, os, strings, and log for file handling, string manipulation, and logging.

Directory Structure
--------------------
- main.go: Main application code.
- templates/index.html: HTML template for rendering the form and ASCII art.
- ArtStyles/: Directory containing ASCII art style files.

Feel free to reach out if you have any questions or need further assistance!
