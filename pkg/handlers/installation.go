package handlers

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/divyam234/instl/pkg/config"
	"github.com/divyam234/instl/pkg/global"
	"github.com/divyam234/instl/pkg/platforms"
	"github.com/divyam234/instl/scripts"
	"github.com/gorilla/mux"
)

func Installation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := strings.ToLower(vars["user"])
	repo := strings.ToLower(vars["repo"])
	platform, err := platforms.Parse(vars["os"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	script, err := scripts.ParseTemplateForPlatform(platform,
		config.Config{
			Owner:        owner,
			Repo:         repo,
			CreatedAt:    time.Now(),
			Version:      "latest",
			InstlVersion: global.Version,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	io.Copy(w, strings.NewReader(script))

}
