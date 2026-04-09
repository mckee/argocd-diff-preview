package reposerverextract

import "testing"

func TestShouldRenderInCrossRepoMode(t *testing.T) {
	tests := []struct {
		name         string
		sourceURL    string
		localRepo    string
		patchRepo    string
		shouldRender bool
	}{
		{
			name:         "same-repo mode renders everything",
			sourceURL:    "https://gitlab.com/sidenio/infrastructure.git",
			localRepo:    "sidenio/infrastructure",
			patchRepo:    "sidenio/infrastructure",
			shouldRender: true,
		},
		{
			name:         "cross-repo: app repo matches patchRepo — render",
			sourceURL:    "https://gitlab.com/sidenio/oc-go/cert-manager-webhook-siden.git",
			localRepo:    "sidenio/infrastructure",
			patchRepo:    "sidenio/oc-go/cert-manager-webhook-siden",
			shouldRender: true,
		},
		{
			name:         "cross-repo: infra repo matches localRepo — skip",
			sourceURL:    "https://gitlab.com/sidenio/infrastructure.git",
			localRepo:    "sidenio/infrastructure",
			patchRepo:    "sidenio/oc-go/cert-manager-webhook-siden",
			shouldRender: false,
		},
		{
			name:         "cross-repo: unrelated repo — skip",
			sourceURL:    "https://gitlab.com/sidenio/oc-go/switchboard.git",
			localRepo:    "sidenio/infrastructure",
			patchRepo:    "sidenio/oc-go/cert-manager-webhook-siden",
			shouldRender: false,
		},
		{
			name:         "cross-repo: empty source URL — render (can't filter)",
			sourceURL:    "",
			localRepo:    "sidenio/infrastructure",
			patchRepo:    "sidenio/oc-go/cert-manager-webhook-siden",
			shouldRender: true,
		},
		{
			name:         "cross-repo: case insensitive match",
			sourceURL:    "https://gitlab.com/Sidenio/OC-Go/Cert-Manager-Webhook-Siden.git",
			localRepo:    "sidenio/infrastructure",
			patchRepo:    "sidenio/oc-go/cert-manager-webhook-siden",
			shouldRender: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShouldRenderInCrossRepoMode(tt.sourceURL, tt.localRepo, tt.patchRepo)
			if got != tt.shouldRender {
				t.Errorf("ShouldRenderInCrossRepoMode(%q, %q, %q) = %v, want %v",
					tt.sourceURL, tt.localRepo, tt.patchRepo, got, tt.shouldRender)
			}
		})
	}
}
