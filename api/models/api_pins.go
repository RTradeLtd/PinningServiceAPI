/*
 * IPFS Pinning Service API
 *
 * This is the IPFS Pinning Service API. It attempts to be an implementation-agnostic API designed to be implemented by Pinning Service Providers.  This specification does **NOT** define the IPFS pinning API.  GitHub discussion around this API can be found in [ipfs/notes](https://github.com/ipfs/notes/issues/378).   **Critical Notice**  To allow backend systems behind the API to associate pins with a particular user, the API requires that a JSON Web Token is passed with requests that contains a `sub` claim (as specified by [OpenIDConnect Standard Claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims)) that represents the user's identity within the backend system, so as to associate the user with their pins. This allows the backend system to use the identity as a form of reference counting on pins, therefore only removing pins once no other users have it pinned.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

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
