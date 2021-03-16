module github.com/zerobugdebug/kaart-api

go 1.16

replace github.com/zerobugdebug/kaart-api/models => ./models

replace github.com/zerobugdebug/kaart-api/controllers => ./controllers

require (
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)
