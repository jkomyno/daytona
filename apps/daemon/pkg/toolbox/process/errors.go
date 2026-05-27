// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package process

import (
	"errors"
	"net/http"

	"github.com/daytonaio/daemon/pkg/common"
	common_errors "github.com/daytonaio/common-go/pkg/errors"
	"github.com/gin-gonic/gin"
)

// ErrProcessExecutionTimeout signals a process exceeded its configured execution timeout.
var ErrProcessExecutionTimeout = errors.New("command execution timeout")

// ErrProcessInvalidCommand signals the command argument failed semantic validation
// (empty, reserved session ID, env-key conflict, etc.).
var ErrProcessInvalidCommand = errors.New("invalid command")

// ErrProcessNotFound signals the requested process artifact (session, command,
// interpreter context, PTY session) was not located.
var ErrProcessNotFound = errors.New("process not found")

func classifyProcessError(err error) error {
	var code common.DaemonErrorCode
	var statusCode int

	switch {
	case errors.Is(err, ErrProcessExecutionTimeout):
		code, statusCode = common.CodeProcessExecutionTimeout, http.StatusRequestTimeout

	case errors.Is(err, ErrProcessInvalidCommand):
		code, statusCode = common.CodeProcessInvalidCommand, http.StatusBadRequest

	case errors.Is(err, ErrProcessNotFound):
		code, statusCode = common.CodeProcessNotFound, http.StatusNotFound

	default:
		code, statusCode = "", http.StatusInternalServerError
	}

	return common_errors.NewCustomError(statusCode, err.Error(), string(code))
}

func abortWithProcessError(c *gin.Context, err error) {
	_ = c.Error(classifyProcessError(err))
	c.Abort()
}
