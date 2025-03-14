// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path/filepath"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Create a new Gin router
// 	r := gin.Default()

// 	// Enable CORS
// 	r.Use(func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
// 		c.Next()
// 	})

// 	// Ensure directory exists
// 	staticPath := filepath.Join("public", ".well-known")
// 	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
// 		fmt.Println("Creating missing directory:", staticPath)
// 		os.MkdirAll(staticPath, os.ModePerm)
// 	}

// 	// Serve static files from public/.well-known directory
// 	r.StaticFS("/.well-known", http.Dir(staticPath))

// 	// Root route
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Deep Link API is Running!")
// 	})

// 	// Handle dynamic routes safely
// 	r.GET("/:path", func(c *gin.Context) {
// 		fullPath := c.Param("path") // Extract dynamic path
// 		query := c.Request.URL.RawQuery
// 		deepLink := fmt.Sprintf("fldeeplink://%s%s", fullPath, func() string {
// 			if query != "" {
// 				return "?" + query
// 			}
// 			return ""
// 		}())

// 		fmt.Println("Redirecting to deep link:", deepLink)
// 		c.Redirect(http.StatusFound, deepLink)
// 	})

// 	// Print all registered routes (for debugging)
// 	for _, route := range r.Routes() {
// 		fmt.Println(route.Method, route.Path)
// 	}

// 	// Start the server on port 3000 (or from environment variable)
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "3000"
// 	}

//		fmt.Println("Server running at http://localhost:" + port)
//		r.Run(":" + port)
//	}
package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Setup Gin Router
func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Ensure .well-known directory exists
	staticPath := filepath.Join("public", ".well-known")
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		fmt.Println("Creating missing directory:", staticPath)
		os.MkdirAll(staticPath, os.ModePerm)
	}

	// Serve static files from public/.well-known directory
	router.StaticFS("/.well-known", http.Dir(staticPath))

	// Root route
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Deep Link API is Running!")
	})

	// Handle dynamic routes safely
	router.GET("/:path", func(c *gin.Context) {
		fullPath := c.Param("path") // Extract dynamic path
		query := c.Request.URL.RawQuery
		deepLink := fmt.Sprintf("fldeeplink://%s%s", fullPath, func() string {
			if query != "" {
				return "?" + query
			}
			return ""
		}())

		fmt.Println("Redirecting to deep link:", deepLink)
		c.Redirect(http.StatusFound, deepLink)
	})

	return router
}

// ✅ This function is required for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router := setupRouter()
	router.ServeHTTP(w, r)
}

// ✅ Runs Locally Only
// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "3000"
// 	}
// 	fmt.Println("Server running at http://localhost:" + port)
// 	setupRouter().Run(":" + port)
// }
