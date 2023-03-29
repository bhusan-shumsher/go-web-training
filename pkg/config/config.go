package config

import (
	"html/template"
	"log"
)

// this package shouldnt import anything aside standard libraries

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
