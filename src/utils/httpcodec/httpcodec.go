package httpcodec

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/pbentes/80_20/src/templates"
)

type Key string

func Encode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return err
		}
	} else {
		var template string
		if !(r.Header.Get("Template") == "") {
			template = r.Header.Get("Template")
		} else {
			template = fmt.Sprintf("%v", r.Context().Value(Key("template")))
		}

		if template == "<nil>" || template == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprint(w, "could not process the request because no template was defined")
			return nil
		}

		if dataMap, ok := data.(*map[string]interface{}); ok {
			ctx := context.WithValue(r.Context(), Key("data"), *dataMap)

			w.Header().Set("Content-Type", "text/html")
			template, err := templates.GetTemplate(template, r.Header.Get("Hx-Boosted") == "")
			if err != nil {
				return err
			}

			template.Render(ctx, w)
		} else if dataArray, ok := data.(*[]map[string]interface{}); ok {
			ctx := context.WithValue(r.Context(), Key("data"), *dataArray)

			w.Header().Set("Content-Type", "text/html")
			template, err := templates.GetTemplate(template, r.Header.Get("Hx-Boosted") == "")
			if err != nil {
				return err
			}

			template.Render(ctx, w)
		} else {
			fmt.Println("Unknown data type")
		}
	}

	return nil
}

func Decode(r *http.Request) (map[string]interface{}, error) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, errors.New("could not parse request body")
	}

	return data, nil
}
