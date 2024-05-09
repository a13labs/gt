/*
Copyright © 2024 Alexandre Pires

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd_test

import (
	"testing"
	"time"

	"github.com/a13labs/gt/cmd"
)

func TestRenderTemplate(t *testing.T) {

	template := `{{.name}} is {{.age}} years old`
	data := []byte(`{"name": "John", "age": 30}`)
	output, err := cmd.RenderTemplate([]byte(template), data)
	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "John is 30 years old" {
		t.Errorf("Expected 'John is 30 years old', got %s", output)
	}
	println("output: ", output)
}

func TestRenderTemplateList(t *testing.T) {

	template := `{{range .}}{{.name}} is {{.age}} years old{{end}}`
	data := []byte(`[{"name": "John", "age": 30}, {"name": "Jane", "age": 25}]`)
	output, err := cmd.RenderTemplate([]byte(template), data)
	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "John is 30 years oldJane is 25 years old" {
		t.Errorf("Expected 'John is 30 years oldJane is 25 years old', got %s", output)
	}
	println("output: ", output)
}

func TestRenderTemplateMap(t *testing.T) {

	template := `{{range $key, $value := .}}{{$key}} is {{$value}} years old{{end}}`
	data := []byte(`{"John": 30, "Jane": 25}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "Jane is 25 years oldJohn is 30 years old" {
		t.Errorf("Expected 'Jane is 25 years oldJohn is 30 years old', got %s", output)
	}
}

func TestRenderTemplateEnv(t *testing.T) {

	template := `{{ env "PATH" }}`
	output, err := cmd.RenderTemplate([]byte(template), nil)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output == "" {
		t.Errorf("Expected non-empty string, got %s", output)
	}
}

func TestRenderTemplateNow(t *testing.T) {

	template := `{{ now }}`
	output, err := cmd.RenderTemplate([]byte(template), nil)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != time.Now().Format(time.RFC3339) {
		t.Errorf("Expected '%s' string, got '%s'", time.Now().Format(time.RFC3339), output)
	}
}

func TestRenderTemplateDate(t *testing.T) {

	template := `{{ now | date "2006-01-02" }}`
	output, err := cmd.RenderTemplate([]byte(template), nil)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != time.Now().Format("2006-01-02") {
		t.Errorf("Expected '%s' string, got '%s'", time.Now().Format("2006-01-02"), output)
	}
}

func TestRenderTemplateAdd(t *testing.T) {

	template := `{{.age | add 5}}`
	data := []byte(`{"age": 30}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "35" {
		t.Errorf("Expected '35', got %s", output)
	}
}

func TestRenderTemplateSub(t *testing.T) {

	template := `{{.age | sub 5}}`
	data := []byte(`{"age": 30}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "25" {
		t.Errorf("Expected '25', got %s", output)
	}
}

func TestRenderTemplateMul(t *testing.T) {

	template := `{{.age | mul 5}}`
	data := []byte(`{"age": 30}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "150" {
		t.Errorf("Expected '150', got %s", output)
	}
}

func TestRenderTemplateDiv(t *testing.T) {

	template := `{{.age | div 5}}`
	data := []byte(`{"age": 30}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "6" {
		t.Errorf("Expected '6', got %s", output)
	}
}

func TestRenderTemplateUpper(t *testing.T) {

	template := `{{.name | upper}}`
	data := []byte(`{"name": "John"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "JOHN" {
		t.Errorf("Expected 'JOHN', got %s", output)
	}
}

func TestRenderTemplateLower(t *testing.T) {

	template := `{{.name | lower}}`
	data := []byte(`{"name": "John"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "john" {
		t.Errorf("Expected 'john', got %s", output)
	}
}

func TestRenderTemplateTrim(t *testing.T) {

	template := `{{.name | trim}}`
	data := []byte(`{"name": " John "}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "John" {
		t.Errorf("Expected 'John', got %s", output)
	}
}

func TestRenderTemplateTrimleft(t *testing.T) {

	template := `{{.name | trimleft}}`
	data := []byte(`{"name": " John "}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "John " {
		t.Errorf("Expected 'John ', got %s", output)
	}
}

func TestRenderTemplateTrimright(t *testing.T) {

	template := `{{.name | trimright}}`
	data := []byte(`{"name": " John "}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != " John" {
		t.Errorf("Expected ' John', got %s", output)
	}
}

func TestRenderTemplateReplace(t *testing.T) {

	template := `{{.name | replace "John" "Jane"}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "Jane Doe" {
		t.Errorf("Expected 'Jane Doe', got %s", output)
	}
}

func TestRenderTemplateContains(t *testing.T) {

	template := `{{if .name | contains "John"}}true{{else}}false{{end}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "true" {
		t.Errorf("Expected 'true', got %s", output)
	}
}

func TestRenderTemplateHasprefix(t *testing.T) {

	template := `{{if .name | hasprefix "John"}}true{{else}}false{{end}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "true" {
		t.Errorf("Expected 'true', got %s", output)
	}
}

func TestRenderTemplateHassuffix(t *testing.T) {

	template := `{{if .name | hassuffix "Doe"}}true{{else}}false{{end}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "true" {
		t.Errorf("Expected 'true', got %s", output)
	}
}

func TestRenderTemplateIndexof(t *testing.T) {

	template := `{{.name | indexof "Doe"}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "5" {
		t.Errorf("Expected '5', got %s", output)
	}
}

func TestRenderTemplateLastindexof(t *testing.T) {

	template := `{{.name | lastindexof "Doe"}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "5" {
		t.Errorf("Expected '5', got %s", output)
	}
}

func TestRenderTemplateReverse(t *testing.T) {

	template := `{{.name | reverse}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "eoD nhoJ" {
		t.Errorf("Expected 'eoD nhoJ', got %s", output)
	}
}

func TestRenderTemplateSubstr(t *testing.T) {

	template := `{{.name | substr 5 3}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "Doe" {
		t.Errorf("Expected 'Doe', got %s", output)
	}
}

func TestRenderTemplateEscapeString(t *testing.T) {

	template := `{{.name | escapeString}}`
	data := []byte(`{"name": "<script>alert('XSS')</script>"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;" {
		t.Errorf("Expected '&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;', got %s", output)
	}
}

func TestRenderTemplateLen(t *testing.T) {

	template := `{{.name | len}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	println("output: ", output)
}

func TestRenderTemplateRegexFind(t *testing.T) {

	template := `{{.name | regexFind ".{3}$" }}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}
	if output != "Doe" {
		t.Errorf("Expected 'Doe', got '%s'", output)
	}
}

func TestRenderTemplateEmpty(t *testing.T) {

	template := `{{if not (empty .name) }}true{{else}}false{{end}}`
	data := []byte(`{"name": "John Doe"}`)
	output, err := cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}

	if output != "true" {
		t.Errorf("Expected 'false', got '%s'", output)
	}

	template = `{{if not (empty .name) }}true{{else}}false{{end}}`
	data = []byte(`{"name": ""}`)
	output, err = cmd.RenderTemplate([]byte(template), data)

	if err != nil {
		t.Errorf("Error rendering template: %v", err)
	}

	if output != "false" {
		t.Errorf("Expected 'true', got '%s'", output)
	}
}
