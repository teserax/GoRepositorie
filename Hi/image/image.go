package image

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"path/filepath"
	"strings"
)

const JPGExt = ".jpg"

var imagesExtensionBucket = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "scraper",
		Name:      "ebay_images_extensions",
		Help:      "Variations of ebay images extensions"},
	[]string{"extension"},
)

// EbayImagesToJPG check if is ebay image URL and replace extention with jpg. If image not is ebay image, then skip it.
func EbayImagesToJPG(images []string) ([]string, error) {
	for i, img := range images {
		// skip if not is ebay image
		if !strings.Contains(img, "i.ebayimg") {
			if strings.Contains(img, "ebay") {
				return nil, fmt.Errorf("for url %q incorrect detect ebay image", img)
			}
			continue
		}
		ext := filepath.Ext(img)
		if ext == "" {
			return nil, fmt.Errorf("for url %q got empty extension", img)
		}
		// append metrics
		if ext == JPGExt {
			continue
		}
		// replace
		images[i] = strings.TrimSuffix(img, ext) + JPGExt
	}
	return images, nil
}
