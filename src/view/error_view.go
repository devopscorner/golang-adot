package view

import (
	"net/http"

	"github.com/devopscorner/golang-adot/src/config"
	"github.com/gin-gonic/gin"
)

// ----- Error Response -----

func ErrorBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func ErrorInternalServer(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func ErrorInvalidId(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": config.ERR_INVALID_BOOK_ID})
}

func ErrorInvalidCredentials(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": config.ERR_INVALID_CREDENTIALS})
	ctx.Abort()
}

func ErrorInvalidToken(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": config.ERR_INVALID_TOKEN})
	ctx.Abort()
}

func ErrorInvalidExpiredToken(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": config.ERR_INVALID_EXPIRED_TOKEN})
	ctx.Abort()
}

func ErrorInvalidExpiredRefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": config.ERR_INVALID_EXPIRED_REFRESH_TOKEN})
	ctx.Abort()
}

func ErrorGenerateToken(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": config.ERR_GENERATE_TOKEN})
}

func ErrorGenerateRefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": config.ERR_GENERATE_REFRESH_TOKEN})
}

func ErrorAuthHeader(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": config.ERR_MISSING_AUTH_HEADER})
	ctx.Abort()
}

func ErrorNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"error": config.ERR_BOOK_NOT_FOUND})
}

func ErrorInvalidRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": config.ERR_INVALID_REQUEST_PAYLOAD})
}

func ErrorUpdate(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": config.ERR_UPDATE_BOOK})
}

func ErrorDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": config.ERR_DELETE_BOOK})
}
