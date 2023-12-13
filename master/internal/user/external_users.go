package user

import (
	"context"

	"github.com/determined-ai/determined/master/pkg/model"
)

// ByExternalToken returns a user session derived from an external authentication token.
func ByExternalToken(ctx context.Context, tokenText string,
	ext *model.ExternalSessions,
) (*model.User, *model.UserSession, error) {
	return nil, nil, nil
}
