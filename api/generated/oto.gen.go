// Code generated by oto; DO NOT EDIT.

package generated

import (
	"context"
	"log"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"
	
)


type Healthcheck interface {

	Check(context.Context, HealthcheckRequest) (*HealthcheckResponse, error)
}

type PullRequestService interface {

	AverageByWeek(context.Context, RepositoryRequest) (*AverageByWeekResponse, error)
	Oldest(context.Context, RepositoryRequest) (*OldestResponse, error)
}



type healthcheckServer struct {
	server *otohttp.Server
	healthcheck Healthcheck
}

func RegisterHealthcheck(server *otohttp.Server, healthcheck Healthcheck) {
	handler := &healthcheckServer{
		server: server,
		healthcheck: healthcheck,
	}
	server.Register("Healthcheck", "Check", handler.handleCheck)
	}

func (s *healthcheckServer) handleCheck(w http.ResponseWriter, r *http.Request) {
	var request HealthcheckRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.healthcheck.Check(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}


type pullRequestServiceServer struct {
	server *otohttp.Server
	pullRequestService PullRequestService
}

func RegisterPullRequestService(server *otohttp.Server, pullRequestService PullRequestService) {
	handler := &pullRequestServiceServer{
		server: server,
		pullRequestService: pullRequestService,
	}
	server.Register("PullRequestService", "AverageByWeek", handler.handleAverageByWeek)
	server.Register("PullRequestService", "Oldest", handler.handleOldest)
	}

func (s *pullRequestServiceServer) handleAverageByWeek(w http.ResponseWriter, r *http.Request) {
	var request RepositoryRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.pullRequestService.AverageByWeek(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *pullRequestServiceServer) handleOldest(w http.ResponseWriter, r *http.Request) {
	var request RepositoryRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.pullRequestService.Oldest(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}




type PRAverage struct {
	Week string `json:"week"`
MeanTimeToMerge int64 `json:"meanTimeToMerge"`
MedianTimeToMerge int64 `json:"medianTimeToMerge"`

}

type AverageByWeekResponse struct {
	GeneratedAt string `json:"generatedAt"`
Averages []PRAverage `json:"averages"`
Error string `json:"error,omitempty"`

}

type HealthcheckRequest struct {
	
}

type HealthcheckResponse struct {
	Ok string `json:"ok"`
Error string `json:"error,omitempty"`

}

type OldestResponse struct {
	Title string `json:"title"`
URL string `json:"uRL"`
OpenFor string `json:"openFor"`
Error string `json:"error,omitempty"`

}

type RepositoryRequest struct {
	Repository string `json:"repository"`

}
