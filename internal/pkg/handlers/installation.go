package handlers

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/installer/instl/internal/pkg/config"
	"github.com/installer/instl/internal/pkg/global"
	"github.com/installer/instl/internal/pkg/platforms"
	"github.com/installer/instl/scripts"
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
