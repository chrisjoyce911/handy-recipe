package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Ingredient represents a group of ingredients.
type Ingredient struct {
	Title   string   `json:"title"`
	Content []string `json:"content"` // Each ingredient as a separate item.
}

// Stage represents a section (e.g. instructions) of a recipe.
type Stage struct {
	Title   string   `json:"title"`
	Content []string `json:"content"` // Each step as a separate string.
	Images  []string `json:"images"`  // URLs or paths to images.
}

// Recipe holds the structured data for a recipe.
type Recipe struct {
	Slug        string       `json:"slug"` // Unique identifier (e.g., "beef-and-broccolli")
	Title       string       `json:"title"`
	CoverImage  string       `json:"cover_image"` // Path to a cover image; if empty, we'll use a fallback.
	Description string       `json:"description"` // A short description.
	Overview    string       `json:"overview"`    // A one-line overview.
	PrepTime    string       `json:"prep_time"`   // e.g. "30 min."
	CookTime    string       `json:"cook_time"`   // e.g. "15 min."
	Servings    string       `json:"servings"`    // e.g. "6 servings"
	Tags        []string     `json:"tags"`        // e.g., ["Beef", "broccoli", "garlic"]
	Ingredients []Ingredient `json:"ingredients"` // Groups of ingredients.
	Stages      []Stage      `json:"stages"`      // Instructions and other sections.
	Favorite    bool         `json:"favorite"`    // for favorites.
}

// TagInfo holds tag name and the count of recipes that have that tag.
type TagInfo struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Store defines the methods our data store should implement.
type Store interface {
	GetRecipeBySlug(slug string) (*Recipe, error)
	GetAllTags() []string
	GetRecipesByTag(tag string) ([]Recipe, error)
	// New method: GetTagsWithCount returns a slice of TagInfo.
	GetTagsWithCount() []TagInfo

	// New method to return all recipes
	GetAllRecipes() ([]Recipe, error)

	GetFavoriteRecipes() ([]Recipe, error)
	GetRecipesWithImages(limit int) ([]Recipe, error)
}

// JSONStore is a concrete implementation of Store that loads each recipe from its own JSON file.
type JSONStore struct {
	recipes []Recipe
}

// NewJSONStoreFromDir loads all recipe JSON files from the specified directory and returns a JSONStore.
func NewJSONStoreFromDir(dir string) (*JSONStore, error) {
	var recipes []Recipe

	// Walk the directory and load each JSON file.
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Only process files that have a .json extension.
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".json") {
			file, err := os.Open(path)
			if err != nil {
				// Log the error and skip the file.
				log.Printf("error opening file %s: %v", path, err)
				return nil
			}
			defer file.Close()

			var r Recipe
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&r); err != nil {
				// Log the error with file path and skip this file.
				log.Printf("error decoding file %s: %v", path, err)
				return nil
			}
			recipes = append(recipes, r)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return nil, fmt.Errorf("no valid recipes found in %s", dir)
	}

	return &JSONStore{recipes: recipes}, nil
}

// GetRecipeBySlug returns the recipe with the matching slug.
func (js *JSONStore) GetRecipeBySlug(slug string) (*Recipe, error) {
	slug = strings.ToLower(slug)
	for _, r := range js.recipes {
		if r.Slug == slug {
			return &r, nil
		}
	}
	return nil, errors.New("recipe not found")
}

// GetAllTags returns a list of unique tags across all recipes.
func (js *JSONStore) GetAllTags() []string {
	tagSet := make(map[string]struct{})
	for _, r := range js.recipes {
		for _, tag := range r.Tags {
			tagSet[tag] = struct{}{}
		}
	}
	tags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tags = append(tags, tag)
	}
	return tags
}

// GetRecipesByTag returns all recipes that contain the specified tag.
func (js *JSONStore) GetRecipesByTag(tag string) ([]Recipe, error) {
	var result []Recipe
	tag = strings.ToLower(tag)
	for _, r := range js.recipes {
		for _, t := range r.Tags {
			if strings.ToLower(t) == tag {
				result = append(result, r)
				break
			}
		}
	}
	if len(result) == 0 {
		return nil, errors.New("no recipes found for tag")
	}
	return result, nil
}

// GetTagsWithCount returns a slice of TagInfo with the count of recipes for each tag,
// sorted alphabetically by tag name.
func (js *JSONStore) GetTagsWithCount() []TagInfo {
	tagMap := make(map[string]int)
	for _, r := range js.recipes {
		for _, t := range r.Tags {
			tagMap[t]++
		}
	}
	var tags []TagInfo
	for name, count := range tagMap {
		tags = append(tags, TagInfo{Name: name, Count: count})
	}
	// Sort tags alphabetically by name.
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	})
	return tags
}

// GetAllRecipes returns all recipes stored.
func (js *JSONStore) GetAllRecipes() ([]Recipe, error) {
	if len(js.recipes) == 0 {
		return nil, fmt.Errorf("no recipes found")
	}
	return js.recipes, nil
}

// GetFavoriteRecipes returns only recipes marked as favorite.
func (js *JSONStore) GetFavoriteRecipes() ([]Recipe, error) {
	var favorites []Recipe
	for _, r := range js.recipes {
		if r.Favorite {
			favorites = append(favorites, r)
		}
	}
	if len(favorites) == 0 {
		return nil, fmt.Errorf("no favorite recipes found")
	}
	return favorites, nil
}

// GetRecipesWithImages returns up to 'limit' recipes that have a non-empty CoverImage in random order.
func (js *JSONStore) GetRecipesWithImages(limit int) ([]Recipe, error) {
	var recipesWithImages []Recipe
	for _, r := range js.recipes {
		if r.CoverImage != "" {
			recipesWithImages = append(recipesWithImages, r)
		}
	}
	if len(recipesWithImages) == 0 {
		return nil, fmt.Errorf("no recipes with images found")
	}
	// Seed the random generator (ideally do this once in your main/init code)

	rand.Shuffle(len(recipesWithImages), func(i, j int) {
		recipesWithImages[i], recipesWithImages[j] = recipesWithImages[j], recipesWithImages[i]
	})
	if limit > len(recipesWithImages) {
		limit = len(recipesWithImages)
	}
	return recipesWithImages[:limit], nil
}
