This Go application parses a Swagger 2.0 JSON file and outputs each API call's verb, endpoint, and parameters. It's a valuable tool for threat modeling, testing checklists, and generating screenshots for walkthroughs.

Usage:

1. Ensure Go is installed on your local machine.
2. Save the code to a file, e.g., `api-map.go`.
3. Compile the application: `go build api-map.go`.
4. Run the compiled executable with the Swagger JSON file as an argument: `./api-map <file.json>`.
5. The results will be printed to the standard output (STDOUT).
