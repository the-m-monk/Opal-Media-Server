package api

import (
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"opal/internal/librarymgmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

const imageNamespace = "9427cfd6-d8cf-47d0-82aa-86c6d17bb390"

// http://localhost:7096/Items/eceec891-3674-5020-ad92-a9cc24d640b2/Images/Primary
// fillHeight=200
// fillWidth=355
// quality=96
// tag=backdrop.jpg
func EndpointItemsByUuidImagesPrimary(w http.ResponseWriter, r *http.Request) {
	//Note: this endpoint does not require authentication (i think)
	query := r.URL.Query()

	itemUuid := r.PathValue("uuid")
	tag := query.Get("tag")

	heightStr := query.Get("fillHeight")
	widthStr := query.Get("fillWidth")
	qualityStr := query.Get("fillWidth")
	height, errH := strconv.Atoi(heightStr)
	width, errW := strconv.Atoi(widthStr)
	quality, errQ := strconv.Atoi(qualityStr)

	if errH != nil || errW != nil || errQ != nil {
		http.Error(w, "could not parse dimensions", http.StatusBadRequest)
		return
	}

	iNS := uuid.MustParse(imageNamespace)

	cacheUuidData := fmt.Sprintf("%s-%s-%s-%s", itemUuid, widthStr, heightStr, tag)
	cacheUuid := uuid.NewSHA1(iNS, []byte(cacheUuidData)).String()

	imageCachePath := filepath.Join(librarymgmt.CacheDir, "images")
	imageFileCachePath := filepath.Join(imageCachePath, cacheUuid+".jpg")

	if _, err := os.Stat(imageFileCachePath); err == nil {
		w.Header().Set("Content-Type", "image/jpeg")
		http.ServeFile(w, r, imageFileCachePath)
		return
	}

	item := librarymgmt.LibTreeMap[itemUuid]
	if item == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	jpegOptions := &jpeg.Options{Quality: quality}

	switch tag {
	case "backdrop.jpg":
		imagePath := filepath.Join(librarymgmt.MetadataDir, item.ImdbId, "backdrop.jpg")
		image, err := imaging.Open(imagePath)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		retImage := imaging.Resize(image, width, height, imaging.Lanczos)

		os.MkdirAll(imageCachePath, 0755)
		out, err := os.Create(imageFileCachePath)
		if err == nil {
			defer out.Close()
			jpeg.Encode(out, retImage, jpegOptions)
		}

		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, retImage, jpegOptions)
	case "library.jpg":
		imageCachePath := filepath.Join(librarymgmt.CacheDir, "images")
		libNameCardPath := filepath.Join(imageCachePath, item.RootUuid+".png")

		image, err := imaging.Open(libNameCardPath)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		retImage := imaging.Resize(image, width, height, imaging.Lanczos)

		os.MkdirAll(imageCachePath, 0755)
		out, err := os.Create(imageFileCachePath)
		if err == nil {
			defer out.Close()
			jpeg.Encode(out, retImage, jpegOptions)
		}

		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, retImage, jpegOptions)

	default:
		log.Printf("[WARN] EndpointItemsByUuidImagesPrimary: image tag \"%s\" not supported\n", tag)
		http.Error(w, "tag type not supported", http.StatusNotImplemented)
		return
	}
}
