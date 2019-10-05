package routes

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
}

func getVars(r *http.Request) (string, string) {
	vars := mux.Vars(r)

	return vars["resource"], vars["uid"]
}

func FetchResources(w http.ResponseWriter, r *http.Request) {
	w = setHeaders(w)
	resource, _ := getVars(r)

	var result string

	switch strings.ToLower(resource) {
	case "pod", "pods":
		// TODO: Fetch pods
	case "deployment", "deployments":
		// TODO: Fetch deployments
	case "daemonset", "daemonsets":
		// TODO: Fetch daemonsets
	default:
		w.WriteHeader(http.StatusBadRequest)
		result = "Unexpected resource type"
	}

	w.Write([]byte(result))
}

func KillResource(w http.ResponseWriter, r *http.Request) {
	w = setHeaders(w)
	resource, _ := getVars(r)

	var result string

	switch strings.ToLower(resource) {
	case "pod", "pods":
		// TODO: Fetch pods
	case "deployment", "deployments":
		// TODO: Fetch deployments
	case "daemonset", "daemonsets":
		// TODO: Fetch daemonsets
	default:
		w.WriteHeader(http.StatusBadRequest)
		result = "Unexpected resource type"
	}

	w.Write([]byte(result))

}
