package v1

import (
	"fmt"
	"github.com/amiosamu/adv-backend-trainee-assignment/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"

	"os"
)

func NewRouter(services *service.Services) *gin.Engine {
	router := gin.New()

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf(`{"time":"%s", "method":"%s","uri":"%s", "status":%d,"error":"%s"}`,
				param.TimeStamp.Format(time.RFC3339Nano),
				param.Method,
				param.Path,
				param.StatusCode,
				param.ErrorMessage,
			)
		},
		Output: setLogsFile(),
	}))
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := router.Group("/api/v1")
	{
		newAdvertisementRoutes(v1.Group("/advertisements"), service.AdvertisementService))
	}

	return router
}

func setLogsFile() *os.File {
	file, err := os.OpenFile("/logs/requests.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
