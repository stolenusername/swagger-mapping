package main

  

import (

"encoding/json"

"fmt"

"io/ioutil"

"os"

"strings"

)

  

// Swagger struct represents the structure of Swagger 2.0 JSON file

type Swagger struct {

Paths map[string]map[string]struct {

Tags []string

Summary string

Description string

}

}

  

func main() {

if len(os.Args) != 2 {

fmt.Println("Usage: ./swagger_parser <swagger.json>")

os.Exit(1)

}

  

filename := os.Args[1]

  

// Read the Swagger JSON file

content, err := ioutil.ReadFile(filename)

if err != nil {

fmt.Printf("Error reading file: %v\n", err)

os.Exit(1)

}

  

// Parse JSON content into Swagger struct

var swagger Swagger

err = json.Unmarshal(content, &swagger)

if err != nil {

fmt.Printf("Error parsing JSON: %v\n", err)

os.Exit(1)

}

  

// Initialize map to store API endpoints

endpoints := make(map[string]string)

  

// Iterate over paths and methods

for path, methods := range swagger.Paths {

for method, info := range methods {

// Construct endpoint URL

endpoint := fmt.Sprintf("%s %s", strings.ToUpper(method), path)

  

// Construct endpoint description

description := info.Summary

if info.Description != "" {

description += " - " + info.Description

}

  

// Store endpoint and description in map

endpoints[endpoint] = description

}

}

  

// Print API endpoint map

fmt.Println("API Endpoint Map:")

for endpoint, description := range endpoints {

fmt.Printf("%s: %s\n", endpoint, description)

}

}
