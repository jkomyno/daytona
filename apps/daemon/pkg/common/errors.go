// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package common

import (
	"time"
)

// DaemonErrorCode identifies the specific error within the daemon.
type DaemonErrorCode string // @name DaemonErrorCode

const (
	// Git
	CodeGitAuthFailed     DaemonErrorCode = "GIT_AUTH_FAILED"
	CodeGitAuthForbidden  DaemonErrorCode = "GIT_AUTH_FORBIDDEN"
	CodeGitRepoNotFound   DaemonErrorCode = "GIT_REPO_NOT_FOUND"
	CodeGitBranchNotFound DaemonErrorCode = "GIT_BRANCH_NOT_FOUND"
	CodeGitRefNotFound    DaemonErrorCode = "GIT_REF_NOT_FOUND"
	CodeGitEmptyRepo      DaemonErrorCode = "GIT_EMPTY_REPO"
	CodeGitPushRejected   DaemonErrorCode = "GIT_PUSH_REJECTED"
	CodeGitDirtyWorktree  DaemonErrorCode = "GIT_DIRTY_WORKTREE"
	CodeGitBranchExists   DaemonErrorCode = "GIT_BRANCH_EXISTS"
	CodeGitMergeConflict  DaemonErrorCode = "GIT_MERGE_CONFLICT"
	CodeGitRepoExists     DaemonErrorCode = "GIT_REPO_EXISTS"

	// File system
	CodeFileNotFound     DaemonErrorCode = "FILE_NOT_FOUND"
	CodeFileAccessDenied DaemonErrorCode = "FILE_ACCESS_DENIED"
	CodeInvalidFilePath  DaemonErrorCode = "INVALID_FILE_PATH"

	// LSP
	CodeLspServerNotInitialized DaemonErrorCode = "LSP_SERVER_NOT_INITIALIZED"
	CodeLspInvalidRequest       DaemonErrorCode = "LSP_INVALID_REQUEST"

	// Process
	CodeProcessExecutionTimeout DaemonErrorCode = "PROCESS_EXECUTION_TIMEOUT"
	CodeProcessInvalidCommand   DaemonErrorCode = "PROCESS_INVALID_COMMAND"
	CodeProcessNotFound         DaemonErrorCode = "PROCESS_NOT_FOUND"

	// Session
	CodeSessionEnded DaemonErrorCode = "SESSION_ENDED"
)

// ErrorResponse is the wire shape of every error response emitted by the daemon.
// Mirrors libs/common-go/pkg/errors.ErrorResponse but declares Code as the typed
// DaemonErrorCode enum so swaggo emits an enum reference in the OpenAPI spec.
//
// MUST be kept in sync with libs/common-go/pkg/errors.ErrorResponse. The two
// declarations are intentionally redundant: clients deserialize against the
// daemon-local one (this), the gin middleware serializes from the common-go
// one. Add a field here whenever you add one there (and vice versa).
//
// IMPORTANT: `code` is optional. Only errors that need to be programmatically
// distinguished beyond their HTTP status carry a code; errors classified by
// status alone (NotFoundError, ConflictError, ...) omit it. Generated SDK
// clients MUST treat `code` and `source` as nullable.
//
//	@Description	Error response
//	@Schema			ErrorResponse
type ErrorResponse struct {
	StatusCode int             `json:"statusCode" example:"400" binding:"required"`
	Message    string          `json:"message" example:"Bad request" binding:"required"`
	Source     string          `json:"source,omitempty" example:"DAYTONA_DAEMON"`
	Code       DaemonErrorCode `json:"code,omitempty" example:"GIT_REPO_NOT_FOUND"`
	Timestamp  time.Time       `json:"timestamp" example:"2023-01-01T12:00:00Z" binding:"required"`
	Path       string          `json:"path" example:"/api/resource" binding:"required"`
	Method     string          `json:"method,omitempty" example:"GET"`
} //	@name	ErrorResponse
