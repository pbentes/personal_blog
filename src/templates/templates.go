// Package templates implements utility functions for use with templates.
package templates

import (
	"fmt"

	"github.com/a-h/templ"
)

var templates map[string][]templ.Component = make(map[string][]templ.Component)

func SetTemplate(key string, template templ.Component, templateWithLayout ...templ.Component) {
	templates[key] = append(templates[key], template)
	if len(templateWithLayout) > 0 {
		templates[key] = append(templates[key], templateWithLayout[0])
	} else {
		templates[key] = append(templates[key], template)
	}
}

func GetTemplate(key string, layout bool) (templ.Component, error) {
	val, ok := templates[key]
	if !ok {
		return nil, fmt.Errorf("the template you tried to access doesn't exist")
	}

	if layout {
		// Return template with the layout
		return val[1], nil
	}

	// Return template without the layout
	return val[0], nil
}

// Get a value from a map as a string
func GetValue(Map map[string]interface{}, key string) string {
	v, ok := Map[key].(string)

	if !ok {
		v = fmt.Sprintf("%v", Map[key])
		if v == "<nil>" {
			v = ""
		}
		return v
	}

	return v
}
