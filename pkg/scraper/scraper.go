package scraper

import (
	"github.com/scottshotgg/zeigarnik/pkg/storage"
)

type (
	Scraper interface {
		Scrape(s storage.Storage)
	}
)
