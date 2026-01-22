package helpers

import (
	"net/http"
	"strconv"
	"strings"
)

func GetIDFromURL(r *http.Request, prefix string) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	return strconv.Atoi(idStr)
}
