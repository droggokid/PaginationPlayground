package internal

import (
	"PaginationPlayground/internal/models"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpClient struct {
	client http.Client
}

func NewHttpClient() HttpClient {
	return HttpClient{client: http.Client{Timeout: time.Second * 10}}
}

func (h HttpClient) GetOSRSData(c *gin.Context) error {
	category := c.DefaultQuery("category", "1")
	alpha := c.DefaultQuery("alpha", "c")
	page := c.DefaultQuery("page", "1")

	u, _ := url.Parse("https://secure.runescape.com/m=itemdb_oldschool/api/catalogue/items.json")
	q := u.Query()
	q.Set("category", category)
	q.Set("alpha", alpha)
	q.Set("page", page)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("User-Agent", "PaginationPlayground/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{
			"upstream_status": resp.Status,
			"content_type":    resp.Header.Get("Content-Type"),
		})
		return nil
	}

	var out models.SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "decode json failed",
			"details": err.Error(),
		})
		return err
	}

	c.JSON(http.StatusOK, out)
	return nil
}
