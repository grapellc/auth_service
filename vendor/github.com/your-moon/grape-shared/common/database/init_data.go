package database

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sync"

	models2 "github.com/your-moon/grape-shared/entities"

	"github.com/sirupsen/logrus"
)

// ModuleData represents the structure from modules.json
type ModuleData struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	IconLink     string `json:"icon_link"`
	URL          string `json:"url"`
	Module       string `json:"module"`
	Status       string `json:"status"`
	MobileStatus string `json:"mobile_status"`
	Index        int    `json:"index"`
}

// AdvertisementWordData represents the structure from advertisement_words.json
type AdvertisementWordData struct {
	ID   int    `json:"id"`
	Word string `json:"word"`
}

// InitData loads initial data from JSON files into the database
func InitData() {
	logrus.Info("Starting data initialization...")

	// First, initialize independent data in parallel
	var wg sync.WaitGroup
	wg.Add(8)

	go func() {
		defer wg.Done()
		initModules()
	}()

	go func() {
		defer wg.Done()
		initSections()
	}()

	go func() {
		defer wg.Done()
		initAdvertisementWords()
	}()

	go func() {
		defer wg.Done()
		initLocationAreas()
	}()

	go func() {
		defer wg.Done()
		initCategories()
	}()

	go func() {
		defer wg.Done()
		initTags()
	}()

	go func() {
		defer wg.Done()
		initUsers()
	}()

	go func() {
		defer wg.Done()
		initMarketTypes()
	}()

	// Wait for independent data to complete
	wg.Wait()

	// Then initialize locations (needed for products)
	initLocations()

	// Finally initialize products (depends on locations)
	initProducts()

	initJobs()

	logrus.Info("Data initialization completed")
}

// initMarketTypes loads market types from market_types.json
func initMarketTypes() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "market_types.json"))
	if err != nil {
		logrus.Warnf("Could not read market_types.json: %v", err)
		return
	}

	var marketTypesData []models2.MarketType
	if err := json.Unmarshal(jsonFile, &marketTypesData); err != nil {
		logrus.Errorf("Error parsing market_types.json: %v", err)
		return
	}

	// Check if any market types exist
	var count int64
	DBClient.Model(&models2.MarketType{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Market types already exist (%d records), skipping", count)
		return
	}

	// Bulk insert all market types
	if err := DBClient.CreateInBatches(marketTypesData, 100).Error; err != nil {
		logrus.Errorf("Error creating market types: %v", err)
	} else {
		logrus.Infof("Created %d market types", len(marketTypesData))
	}
}

func initJobs() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "jobs.json"))
	if err != nil {
		logrus.Warnf("Could not read jobs.json: %v", err)
		return
	}

	var jobsData []models2.Job
	if err := json.Unmarshal(jsonFile, &jobsData); err != nil {
		logrus.Errorf("Error parsing jobs.json: %v", err)
		return
	}

	var count int64
	DBClient.Model(&models2.Job{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Jobs already exist (%d records), skipping", count)
		return
	}

	if err := DBClient.CreateInBatches(jobsData, 100).Error; err != nil {
		logrus.Errorf("Error creating jobs: %v", err)
	} else {
		logrus.Infof("Created %d jobs", len(jobsData))
	}
}

// initSections loads sections from sections.json
func initSections() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "sections.json"))
	if err != nil {
		logrus.Warnf("Could not read sections.json: %v", err)
		return
	}

	var sectionsData []ModuleData
	if err := json.Unmarshal(jsonFile, &sectionsData); err != nil {
		logrus.Errorf("Error parsing sections.json: %v", err)
		return
	}

	// Check if any sections exist
	var count int64
	DBClient.Model(&models2.Section{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Sections already exist (%d records), skipping", count)
		return
	}

	// Convert to models and bulk insert
	var sections []models2.Section
	for _, sectionData := range sectionsData {
		section := models2.Section{
			Name:        sectionData.Name,
			Slug:        &sectionData.Slug,
			Description: &sectionData.Description,
			IconLink:    &sectionData.IconLink,
			URL:         &sectionData.URL,
			Status:      sectionData.Status,
			Index:       sectionData.Index,
		}
		sections = append(sections, section)
	}

	// Bulk insert all sections
	if err := DBClient.CreateInBatches(sections, 100).Error; err != nil {
		logrus.Errorf("Error creating sections: %v", err)
	} else {
		logrus.Infof("Created %d sections", len(sections))
	}
}

// initModules loads modules from modules.json
func initModules() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "modules.json"))
	if err != nil {
		logrus.Warnf("Could not read modules.json: %v", err)
		return
	}

	var modulesData []ModuleData
	if err := json.Unmarshal(jsonFile, &modulesData); err != nil {
		logrus.Errorf("Error parsing modules.json: %v", err)
		return
	}

	// Check if any modules exist
	var count int64
	DBClient.Model(&models2.Module{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Modules already exist (%d records), skipping", count)
		return
	}

	// Convert to models and bulk insert
	var modules []models2.Module
	for _, moduleData := range modulesData {
		module := models2.Module{
			Name:         moduleData.Name,
			Slug:         &moduleData.Slug,
			Description:  &moduleData.Description,
			IconLink:     &moduleData.IconLink,
			URL:          &moduleData.URL,
			Status:       moduleData.Status,
			MobileStatus: moduleData.MobileStatus,
			Index:        moduleData.Index,
		}
		modules = append(modules, module)
	}

	// Bulk insert all modules
	if err := DBClient.CreateInBatches(modules, 100).Error; err != nil {
		logrus.Errorf("Error creating modules: %v", err)
	} else {
		logrus.Infof("Created %d modules", len(modules))
	}
}

// initAdvertisementWords loads advertisement words from advertisement_words.json
func initAdvertisementWords() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "advertisement_words.json"))
	if err != nil {
		logrus.Warnf("Could not read advertisement_words.json: %v", err)
		return
	}

	var wordsData []AdvertisementWordData
	if err := json.Unmarshal(jsonFile, &wordsData); err != nil {
		logrus.Errorf("Error parsing advertisement_words.json: %v", err)
		return
	}

	// Check if any advertisement words exist
	var count int64
	DBClient.Model(&models2.AdvertisementWords{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Advertisement words already exist (%d records), skipping", count)
		return
	}

	// Convert to models and bulk insert
	var words []models2.AdvertisementWords
	for _, wordData := range wordsData {
		word := models2.AdvertisementWords{
			Word: &wordData.Word,
		}
		words = append(words, word)
	}

	// Bulk insert all advertisement words
	if err := DBClient.CreateInBatches(words, 100).Error; err != nil {
		logrus.Errorf("Error creating advertisement words: %v", err)
	} else {
		logrus.Infof("Created %d advertisement words", len(words))
	}
}

// initLocationAreas loads location areas from location_areas.json
func initLocationAreas() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "location_areas.json"))
	if err != nil {
		logrus.Warnf("Could not read location_areas.json: %v", err)
		return
	}

	var locationAreasData []models2.LocationArea
	if err := json.Unmarshal(jsonFile, &locationAreasData); err != nil {
		logrus.Errorf("Error parsing location_areas.json: %v", err)
		return
	}

	// Check if any location areas exist
	var count int64
	DBClient.Model(&models2.LocationArea{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Location areas already exist (%d records), skipping", count)
		return
	}

	// Bulk insert all location areas
	if err := DBClient.CreateInBatches(locationAreasData, 100).Error; err != nil {
		logrus.Errorf("Error creating location areas: %v", err)
	} else {
		logrus.Infof("Created %d location areas", len(locationAreasData))
	}
}

// initLocations loads locations from locations.json
func initLocations() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "locations.json"))
	if err != nil {
		logrus.Warnf("Could not read locations.json: %v", err)
		return
	}

	var locationsData []models2.Location
	if err := json.Unmarshal(jsonFile, &locationsData); err != nil {
		logrus.Errorf("Error parsing locations.json: %v", err)
		return
	}

	// Check if any locations exist
	var count int64
	DBClient.Model(&models2.Location{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Locations already exist (%d records), skipping", count)
		return
	}

	// Bulk insert all locations
	if err := DBClient.CreateInBatches(locationsData, 100).Error; err != nil {
		logrus.Errorf("Error creating locations: %v", err)
	} else {
		logrus.Infof("Created %d locations", len(locationsData))
	}
}

// initCategories loads categories from categories.json
func initCategories() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "categories.json"))
	if err != nil {
		logrus.Warnf("Could not read categories.json: %v", err)
		return
	}

	var categoriesData []models2.Category
	if err := json.Unmarshal(jsonFile, &categoriesData); err != nil {
		logrus.Errorf("Error parsing categories.json: %v", err)
		return
	}

	// Check if any categories exist
	var count int64
	DBClient.Model(&models2.Category{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Categories already exist (%d records), skipping", count)
		return
	}

	// Bulk insert all categories
	if err := DBClient.CreateInBatches(categoriesData, 100).Error; err != nil {
		logrus.Errorf("Error creating categories: %v", err)
	} else {
		logrus.Infof("Created %d categories", len(categoriesData))
	}
}

// initTags loads tags from tags.json
func initTags() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "tags.json"))
	if err != nil {
		logrus.Warnf("Could not read tags.json: %v", err)
		return
	}

	var tagsData []models2.Tag
	if err := json.Unmarshal(jsonFile, &tagsData); err != nil {
		logrus.Errorf("Error parsing tags.json: %v", err)
		return
	}

	// Check if any tags exist
	var count int64
	DBClient.Model(&models2.Tag{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Tags already exist (%d records), skipping", count)
		return
	}

	// Bulk insert all tags
	if err := DBClient.CreateInBatches(tagsData, 100).Error; err != nil {
		logrus.Errorf("Error creating tags: %v", err)
	} else {
		logrus.Infof("Created %d tags", len(tagsData))
	}
}

// initUsers loads users from users.json
func initUsers() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "users.json"))
	if err != nil {
		logrus.Warnf("Could not read users.json: %v", err)
		return
	}

	var usersData []models2.User
	if err := json.Unmarshal(jsonFile, &usersData); err != nil {
		logrus.Errorf("Error parsing users.json: %v", err)
		return
	}

	// Check if any users exist
	var count int64
	DBClient.Model(&models2.User{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Users already exist (%d records), skipping", count)
		return
	}

	// Bulk insert all users
	if err := DBClient.CreateInBatches(usersData, 100).Error; err != nil {
		logrus.Errorf("Error creating users: %v", err)
	} else {
		logrus.Infof("Created %d users", len(usersData))
	}
}

// initProducts loads products from products.json
func initProducts() {
	jsonFile, err := ioutil.ReadFile(filepath.Join("init", "products.json"))
	if err != nil {
		logrus.Warnf("Could not read products.json: %v", err)
		return
	}

	var productsData []models2.Product
	if err := json.Unmarshal(jsonFile, &productsData); err != nil {
		logrus.Errorf("Error parsing products.json: %v", err)
		return
	}

	// Check if any products exist
	var count int64
	DBClient.Model(&models2.Product{}).Count(&count)
	if count > 0 {
		logrus.Debugf("Products already exist (%d records), skipping", count)
		return
	}

	// Check what location IDs actually exist
	var locationIDs []uint
	DBClient.Model(&models2.Location{}).Pluck("id", &locationIDs)
	logrus.Infof("Available location IDs: %v", locationIDs)

	// Create a map for quick lookup
	locationExists := make(map[uint]bool)
	for _, id := range locationIDs {
		locationExists[id] = true
	}

	// Filter products to only include those with valid location IDs
	var validProducts []models2.Product
	for _, product := range productsData {
		if product.LocationID != nil && locationExists[*product.LocationID] {
			validProducts = append(validProducts, product)
		} else if product.LocationID == nil {
			// Allow products with nil location_id
			validProducts = append(validProducts, product)
		} else {
			logrus.Warnf("Skipping product '%s' - location_id %d does not exist", product.Name, *product.LocationID)
		}
	}

	// Insert valid products
	if len(validProducts) > 0 {
		if err := DBClient.CreateInBatches(validProducts, 100).Error; err != nil {
			logrus.Errorf("Error creating products: %v", err)
		} else {
			logrus.Infof("Created %d products (skipped %d with invalid location_ids)", len(validProducts), len(productsData)-len(validProducts))
		}
	} else {
		logrus.Warnf("No valid products to insert - all products have invalid location_ids")
	}
}
