package studentslist_query

import (
	"fmt"
	"os/exec"

	"github.com/graphql-go/graphql"
)

// type File struct {
// 	Input string `json:"input"`
// }

type FileResponse struct {
	Message    string `json:"message"`
	Statuscode int    `json:"statuscode"`
	Data       string `json:"data"`
}

func Fetchpythonfile(params graphql.ResolveParams) (interface{}, error) {
	//pythonPath := `C:\Users\niharika.m.s\AppData\Local\Programs\Python\Python312\python.exe`
	var output string
	output, err := runPythonScript("python3", "helloworld.py")
	if err != nil {
		fmt.Printf("Error executing Python script: %s\n", err)
		return "", nil
	}
	fmt.Println("Output from python file - ", output)

	return FileResponse{Message: "Successfully fetched file", Statuscode: 200, Data: output}, nil
	//return FileResponse, nil
}

func runPythonScript(pythonPath, scriptPath string) (string, error) {
	cmd := exec.Command(pythonPath, scriptPath)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing Python script: %s", err)
	}
	return string(output), nil
}
