package handlers

import (
	"net/http"
	"strings"
)

// --- Handler functions ---

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	// Get tag data from the store.
	tagInfos := h.Store.GetTagsWithCount()

	// Retrieve all recipes. (Assuming GetAllRecipes returns ([]Recipe, error))
	recipes, err := h.Store.GetFavoriteRecipes()
	if err != nil {
		http.Error(w, "No recipes found", http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"Title":   "Home Page",
		"Message": "Welcome to our website!",
		"Tags":    tagInfos,
		"Recipes": recipes,
	}
	renderTemplate(w, "home.html", data)
}

func (h *Handler) about(w http.ResponseWriter, r *http.Request) {

	recipes, err := h.Store.GetRecipesWithImages(3)
	if err != nil {
		http.Error(w, "No recipes found", http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"Title":   "About Us",
		"Message": "Learn more about us on this page.",
		"Recipes": recipes,
	}
	renderTemplate(w, "about.html", data)
}

func (h *Handler) tags(w http.ResponseWriter, r *http.Request) {
	// Use the new GetTagsWithCount function.
	tagInfos := h.Store.GetTagsWithCount()
	data := map[string]interface{}{
		"Title": "Tags",
		"Tags":  tagInfos,
	}
	renderTemplate(w, "tags.html", data)
}

func (h *Handler) recipes(w http.ResponseWriter, r *http.Request) {
	// Get tag data from the store.
	tagInfos := h.Store.GetTagsWithCount()

	// Retrieve all recipes. (Assuming GetAllRecipes returns ([]Recipe, error))
	recipes, err := h.Store.GetAllRecipes()
	if err != nil {
		http.Error(w, "No recipes found", http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"Title":   "All Recipes",
		"Tags":    tagInfos,
		"Recipes": recipes,
	}
	renderTemplate(w, "recipes.html", data)
}

func (h *Handler) contact(w http.ResponseWriter, r *http.Request) {

	recipes, err := h.Store.GetRecipesWithImages(3)
	if err != nil {
		http.Error(w, "No recipes found", http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"Title":   "Contact",
		"Message": "Contact us at contact@example.com",
		"Recipes": recipes,
	}
	renderTemplate(w, "contact.html", data)
}

func (h *Handler) recipe(w http.ResponseWriter, r *http.Request) {
	// Extract the slug from the URL using the new API.
	slug := r.PathValue("slug")
	slug = strings.ToLower(slug)
	recipe, err := h.Store.GetRecipeBySlug(slug)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}
	data := map[string]interface{}{
		"Title":  recipe.Title,
		"Recipe": recipe,
	}
	renderTemplate(w, "recipe.html", data)
}

func (h *Handler) tag(w http.ResponseWriter, r *http.Request) {
	// Extract the tag from the URL.
	tag := r.PathValue("tag")
	tag = strings.Title(tag)
	// Fetch recipes by tag (assuming the store's GetRecipesByTag method expects lowercase).
	recipes, err := h.Store.GetRecipesByTag(strings.ToLower(tag))
	if err != nil {
		http.Error(w, "No recipes found for tag", http.StatusNotFound)
		return
	}

	tagInfos := h.Store.GetTagsWithCount()

	data := map[string]interface{}{
		"Title":   "Tag: " + tag,
		"Tag":     tag,
		"Recipes": recipes,
		"Tags":    tagInfos,
	}
	renderTemplate(w, "recipes.html", data)
}
