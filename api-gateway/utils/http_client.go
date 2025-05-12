package utils

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ForwardRequest(c *gin.Context, target string) {
	method := c.Request.Method
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	url := target + path
	if query != "" {
		url += "?" + query
	}

	body, _ := io.ReadAll(c.Request.Body)
	req, err := http.NewRequest(method, url, io.NopCloser(strings.NewReader(string(body))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Service unreachable"})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}
