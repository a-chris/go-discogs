package discogs

import (
	"net/url"
	"strconv"
)

const (
	listingURI          = "/listings/"
	priceSuggestionsURI = "/price_suggestions/"
	releaseStatsURI     = "/stats/"
)

type marketPlaceService struct {
	url      string
	currency string
}

type MarketPlaceService interface {
	// The best price suggestions according to grading
	// Authentication is required.
	PriceSuggestions(releaseID int) (*PriceListing, error)
	// Short summary of marketplace listings
	// Authentication is optional.
	ReleaseStatistics(releaseID int) (*Stats, error)
	// Listing returns the listing by listing's ID.
	Listing(listingID int) (*Listing, error)
}

func newMarketPlaceService(url string, currency string) MarketPlaceService {
	return &marketPlaceService{
		url:      url,
		currency: currency,
	}
}

// Listing is a marketplace listing with the user's currency and a price value
type Price struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

// PriceListings are Listings per grading quality
type PriceListing struct {
	VeryGood     *Listing `json:"Very Good (VG),omitempty"`
	GoodPlus     *Listing `json:"Good Plus (G+),omitempty"`
	NearMint     *Listing `json:"Near Mint (NM or M-)"`
	Good         *Listing `json:"Good (G),omitempty"`
	VeryGoodPlus *Listing `json:"Very Good Plus (VG+),omitempty"`
	Mint         *Listing `json:"Mint (M),omitempty"`
	Fair         *Listing `json:"Fair (F),omitempty"`
	Poor         *Listing `json:"Poor (P),omitempty"`
}

// Stats returns the marketplace stats summary for a release containing
type Stats struct {
	LowestPrice *Listing `json:"lowest_price"`
	ForSale     int      `json:"num_for_sale"`
	Blocked     bool     `json:"blocked_from_sale"`
}

type Listing struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Status          string `json:"status"`
	Price           Price  `json:"price"`
	Condition       string `json:"condition"`
	SleeveCondition string `json:"sleeve_condition"`
	ShipsFrom       string `json:"ships_from"`
	Comments        string `json:"comments"`
	Location        string `json:"location"`
}

func (s *marketPlaceService) Listing(listingID int) (*Listing, error) {
	params := url.Values{}
	params.Set("curr_abbr", s.currency)

	var listing *Listing
	err := request(s.url+listingURI+strconv.Itoa(listingID), params, &listing)
	return listing, err
}

func (s *marketPlaceService) ReleaseStatistics(releaseID int) (*Stats, error) {
	params := url.Values{}
	params.Set("curr_abbr", s.currency)

	var stats *Stats
	err := request(s.url+releaseStatsURI+strconv.Itoa(releaseID), params, &stats)
	return stats, err
}

func (s *marketPlaceService) PriceSuggestions(releaseID int) (*PriceListing, error) {
	var listings *PriceListing
	err := request(s.url+priceSuggestionsURI+strconv.Itoa(releaseID), nil, &listings)
	return listings, err
}
