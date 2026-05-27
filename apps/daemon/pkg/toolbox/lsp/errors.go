// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package lsp

import (
	"errors"
	"net/http"

	"github.com/daytonaio/daemon/pkg/common"
	common_errors "github.com/daytonaio/common-go/pkg/errors"
	"github.com/gin-gonic/gin"
)

// ErrLspServerNotInitialized signals the LSP server must be started via /lsp/start first.
var ErrLspServerNotInitialized = errors.New("server not initialized")

func classifyLspError(err error) error {
	var code common.DaemonErrorCode
	var statusCode int

	switch {
	case errors.Is(err, ErrLspServerNotInitialized):
		code, statusCode = common.CodeLspServerNotInitialized, http.StatusBadRequest

	default:
		code, statusCode = common.CodeLspInvalidRequest, http.StatusBadRequest
	}

	return common_errors.NewCustomError(statusCode, err.Error(), string(code))
}

func abortWithLspError(c *gin.Context, err error) {
	_ = c.Error(classifyLspError(err))
	c.Abort()
}
