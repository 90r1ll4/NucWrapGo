package wrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Data struct {
	CurlCommand      string   `json:"curl-command"`
	ExtractedResults []string `json:"extracted-results"`
	Host             string   `json:"host"`
	Info             struct {
		Author      []string `json:"author"`
		Description string   `json:"description"`
		Name        string   `json:"name"`
		Reference   string   `json:"reference"`
		Severity    string   `json:"severity"`
		Tags        []string `json:"tags"`
	} `json:"info"`
	IP            string `json:"ip"`
	MatchedAt     string `json:"matched-at"`
	MatchedLine   string `json:"matched-line"`
	MatcherStatus bool   `json:"matcher-status"`
	Template      string `json:"template"`
	TemplateID    string `json:"template-id"`
	TemplatePath  string `json:"template-path"`
	TemplateURL   string `json:"template-url"`
	Timestamp     string `json:"timestamp"`
	Type          string `json:"type"`
}

func Nuclei(url string, url_list string, json_value bool, tables bool) {

	//adding timer
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	argstr := []string{}
	if url != "" {
		argstr = append(argstr, "-u", url)
	} else if url_list != "" {
		argstr = append(argstr, "-ul", url_list)
	}

	if json_value || tables {
		argstr = append(argstr, "--json")
	}

	//Running shell command
	cmd := exec.CommandContext(ctx, "nuclei", argstr...)
	out, err := cmd.Output()
	if err == nil {
		fmt.Println("Error running nuclei")
	}

	json_changes := "[" + strings.Replace(string(out), "}\n{", "},\n{", -1) + "]"

	// Converting in json formate with indentation
	var jsonData []interface{}
	errors := json.Unmarshal([]byte(json_changes), &jsonData)
	if errors != nil {
		panic(err)
	}

	jsonForm, err := json.MarshalIndent(jsonData, "", "\t")
	if err != nil {
		panic(err)
	}
	if json_value {
		fmt.Println(string(jsonForm))
	}
	// For genrating the tables
	if tables {
		var items []Data
		err := json.Unmarshal(jsonForm, &items)
		if err != nil {
			panic(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Template-Id", "Name", "Host", "Type", "severity"})
		for i := range items {
			row := []string{items[i].TemplateID, items[i].Info.Name, items[i].Host, items[i].Type, items[i].Info.Severity}
			table.Append(row)
		}
		table.Render()
	}

}
