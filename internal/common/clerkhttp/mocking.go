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

type MockWebHookVerifier struct{}

func (m MockWebHookVerifier) Verify(payload []byte, headers http.Header) error {
	return nil
}
