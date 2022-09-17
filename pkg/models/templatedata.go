package models

//TemplateData holds data sents from handler to templates
type TemplateData struct {
	StringMap map[string]string
	InMap     map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //cross-site request token
	Flash     string //message sent to the end-user
	Warning   string
	Error     string
}
