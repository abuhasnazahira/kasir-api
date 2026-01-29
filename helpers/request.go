package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func GetIDFromURL(r *http.Request, prefix string) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	return strconv.Atoi(idStr)
}

// GetQueryInt mengambil nilai query parameter sebagai integer
func GetQueryInt(r *http.Request, key string, defaultValue int) (int, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue, nil
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue, errors.New("invalid " + key + " parameter")
	}
	return intVal, nil
}

// GetQueryString mengambil nilai query parameter sebagai string
func GetQueryString(r *http.Request, key string, defaultValue string) string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}
	return val
}

// ParseRequestBody mengubah request body menjadi struct
func ParseRequestBody(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.New("invalid request body format")
	}

	return nil
}
