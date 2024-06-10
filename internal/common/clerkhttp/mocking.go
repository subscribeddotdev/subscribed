package clerkhttp

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
)

func SetupClerkForTestingMode(serverAddress string) {
	clerk.SetBackend(clerk.NewBackend(&clerk.BackendConfig{
		HTTPClient: &http.Client{},
		URL:        &serverAddress,
	}))
}
