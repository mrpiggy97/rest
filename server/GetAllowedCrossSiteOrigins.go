package server

import (
	"os"
	"strings"
)

func GetAllowedCrossSiteOrigin() []string {
	var allowedOriginsString string = os.Getenv("ALLOWED_CROSS_SITE_ORIGINS")
	var allowedOrigins []string = strings.Split(allowedOriginsString, ",")
	return allowedOrigins
}
