// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package computeruse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// computerUseDisabledMiddleware returns a middleware that handles requests when computer-use is disabled
func ComputerUseDisabledMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	_ = c.Error(newComputerUseError(http.StatusServiceUnavailable, "computer-use plugin disabled: missing X11 dependencies", CodeComputerUseUnavailable))
		c.Abort()
	}
}
