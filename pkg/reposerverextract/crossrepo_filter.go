package reposerverextract

import "github.com/dag-andersen/argocd-diff-preview/pkg/repository"

// ShouldRenderInCrossRepoMode decides whether a child application should be
// rendered when running in cross-repo mode (--local-repo differs from --repo).
//
// In cross-repo mode:
//   - patchRepo is the app repo being diffed (e.g. sidenio/oc-go/cert-manager-webhook-siden)
//   - localRepo is the infra repo checked out locally (e.g. sidenio/infrastructure)
//
// Only child apps whose source repoURL matches patchRepo should be rendered.
// Infra apps and unrelated apps are traversal targets only — rendering them
// would fail (wrong branch) or produce noise.
//
// Returns true if the app should be rendered.
func ShouldRenderInCrossRepoMode(sourceRepoURL string, localSelector, patchSelector repository.Selector) bool {
	// Not cross-repo mode — render everything
	if localSelector.String() == patchSelector.String() {
		return true
	}

	// No source URL — can't filter, render it
	if sourceRepoURL == "" {
		return true
	}

	// Keep apps matching the app repo being diffed
	return patchSelector.Matches(sourceRepoURL)
}
