package twitter_test

import (
	"github.com/Zereker/scraper/pkg/twitter"
	"testing"
)

func TestGetGuestToken(t *testing.T) {
	scraper := twitter.New()
	if err := scraper.GetGuestToken(); err != nil {
		t.Errorf("getGuestToken() error = %v", err)
	}
	if !scraper.IsGuestToken() {
		t.Error("Expected non-empty guestToken")
	}
}
