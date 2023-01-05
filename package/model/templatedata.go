package model

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	warning   string
	Error     string
}
