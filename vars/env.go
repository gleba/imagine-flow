package vars

import "os"

var PublicImagesFolder = getEnv("PUBLIC_IMAGES_FOLDER", "../images") // "/Users/g.panteleev/Pictures"
var AliveTime = getEnv("ALIVE_TIME", "24h")
var URL_PREFIX = getEnv("URL_PREFIX", "i")
var PORT = getEnv("PORT", "5555")

func getEnv(key string, ifNoneValue string) string {
	var v = os.Getenv(key)
	if v == "" {
		return ifNoneValue
	} else {
		return v
	}
}
