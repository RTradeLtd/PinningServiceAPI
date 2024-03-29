/*
 * IPFS Pinning Service
 *
 * This is the IPFS Pinning Service API. It attempts to be an implementation-agnostic API.
 *
 * API version: v0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PinsCidDelete - Remove a Pin
func PinsCidDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PinsCidGet - Get status of a Pin
func PinsCidGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PinsCidPost - Add a new Pin
func PinsCidPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PinsGet - Get all pins
func PinsGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PinsPost - Add an array of pins
func PinsPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
