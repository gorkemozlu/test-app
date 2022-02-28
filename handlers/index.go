package handlers

import (
	"fmt"
	"net/http"

	"github.com/cloudfoundry-samples/test-app/helpers"
)

type Index struct {
	Index string
}

func (_ *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, err := helpers.FetchIndex()

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(fmt.Sprintf("%d", index)))
}
