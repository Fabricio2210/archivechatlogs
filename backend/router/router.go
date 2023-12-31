package router

import (
	"encoding/json"
	"github.com/Fabricio2210/gofiber/elastic"
	"github.com/Fabricio2210/gofiber/pagination"
	"github.com/Fabricio2210/gofiber/rawInfo"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"math"
)

type Request struct {
	UserName string `json:"userName"`
	Message  string `json:"message"`
	Hour     string `json:"hour"`
	DateFrom string `json:"dateFrom"`
	DateEnd  string `json:"dateEnd"`
}

// DefaultRouter handles the default router for the specified subject
func DefaultRouter(app *fiber.App, subject string) {
	app.Post("/"+subject, func(c *fiber.Ctx) error {
		// Parse the "page" query parameter
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Println("Error parsing page:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "Invalid page number"})
		}
		
		// Parse the "limit" query parameter
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			log.Println("Error parsing limit:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "Invalid limit value"})
		}
		
		// Get pagination information
		paginationInfo := pagination.PaginationInfo(page)
		nextPage := paginationInfo["nextPage"]
		previousPage :=  paginationInfo["previousPage"]
		
		// Build the Elasticsearch query
		query := elastic.Query(c, subject)
		
		// Parse the request body
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Check if at least one field is filled in the request
		if req.UserName == "" && req.Message == "" && req.Hour == "" && req.DateFrom == "" && req.DateEnd == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Fill at least one field",
			})
		} else {
			// Perform the search and retrieve raw log info
			rawLogInfo, hit, _ := rawInfo.Search(query, page, limit)

			totalPages := int(math.Ceil(hit / float64(limit)))

			// Define the response data structure
			type DataLog struct {
				Data         []rawInfo.LogInfo `json:"data"`
				Page         int               `json:"page"`
				NextPage     int               `json:"nextPage"`
				PreviousPage int               `json:"previousPage"`
				TotalPages   int               `json:"totalPages"`
				TotalResults int               `json:"totalResults"`
			}

			// Populate the response data
			data := DataLog{
				Data:         rawLogInfo,
				Page:		  page,
				PreviousPage: previousPage,
				NextPage:     nextPage,
				TotalPages:   totalPages,
				TotalResults: int(hit),
			}

			// Convert the data to JSON
			jsonData, err := json.Marshal(data)
			if err != nil {
				return err
			}

			// Set the appropriate headers
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			// Send the JSON data in the response body
			return c.Send(jsonData)
		}
	})
}
