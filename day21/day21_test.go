package main

import "testing"

func TestParseIngredientMap(t *testing.T) {
	input := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}

	foods := parseFoods(input)
	allergens := processFoods(foods)
	ingredientCount := countIngredients(foods)
	noAllergenicIngredients := removeAllergenic(ingredientCount, allergens)

	if count := sumIngredientCount(noAllergenicIngredients); count != 5 {
		t.Fatalf("countNonAllergenic: %v(%v)", count, 5)
	}
}
