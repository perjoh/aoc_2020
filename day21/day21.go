package main

import (
	"aoc2020/util"
	"fmt"
	"sort"
	"strings"
)

type foodType struct {
	ingredients []string
	allergens   []string
}

func parseFood(input string) ([]string, []string) {
	tmp := strings.Split(input, "(contains ")
	ingredients := strings.Split(strings.TrimSpace(tmp[0]), " ")
	allergens := strings.Split(tmp[1][:len(tmp[1])-1], ", ")
	return ingredients, allergens
}

func parseFoods(input []string) []foodType {
	foods := []foodType{}
	for _, line := range input {
		ingredients, allergens := parseFood(line)
		foods = append(foods, foodType{ingredients: ingredients, allergens: allergens})
	}
	return foods
}

func innerJoin(lstA []string, lstB []string) []string {
	result := []string{}
	for _, a := range lstA {
		for _, b := range lstB {
			if a == b {
				result = append(result, a)
			}
		}
	}
	return result
}

func inList(lst []string, val string) bool {
	for _, a := range lst {
		if a == val {
			return true
		}
	}
	return false
}

func leftJoin(left []string, right []string) []string {
	result := []string{}
	for _, a := range left {
		if !inList(right, a) {
			result = append(result, a)
		}
	}
	return result
}

type allergenMap map[string][]string
type ingredientMap map[string]int

func processFoodsInternal(foods []foodType, aMap *allergenMap) bool {
	for _, food := range foods {
		for _, allergen := range food.allergens {
			oldIngredients, has := (*aMap)[allergen]
			if !has {
				(*aMap)[allergen] = food.ingredients
			} else {
				newIngredients := innerJoin(food.ingredients, oldIngredients)
				(*aMap)[allergen] = newIngredients
				if len(newIngredients) == 1 {
					// remove ingredient from other allergens
					removed := false
					for key, val := range *aMap {
						if key != allergen {
							newVal := leftJoin(val, newIngredients)
							if len(newVal) < len(val) {
								(*aMap)[key] = leftJoin(val, newIngredients)
								removed = true
							}
						}
					}

					if removed {
						return true
					}
				}
			}
		}
	}
	return false
}

func processFoods(foods []foodType) allergenMap {
	aMap := allergenMap{}
	for processFoodsInternal(foods, &aMap) {
	}
	return aMap
}

func countIngredients(foods []foodType) ingredientMap {
	ing := ingredientMap{}
	for _, food := range foods {
		for _, i := range food.ingredients {
			ing[i]++
		}
	}
	return ing
}

func removeAllergenic(ingredients ingredientMap, allergens allergenMap) ingredientMap {
	for _, val := range allergens {
		if len(val) == 1 {
			delete(ingredients, val[0])
		}
	}
	return ingredients
}

func sumIngredientCount(ingredients ingredientMap) int {
	sum := 0
	for _, count := range ingredients {
		sum += count
	}
	return sum
}

func part1(input []string) int {
	foods := parseFoods(input)
	allergens := processFoods(foods)
	ingredientCount := countIngredients(foods)
	noAllergenicIngredients := removeAllergenic(ingredientCount, allergens)
	return sumIngredientCount(noAllergenicIngredients)
}

func report(allergens allergenMap) string {
	tmp := []string{}
	for key, val := range allergens {
		tmp = append(tmp, key+":"+val[0])
	}
	sort.Strings(tmp)

	result := ""
	for _, line := range tmp {
		result += strings.Split(line, ":")[1] + ","
	}
	return result[:len(result)-1]
}

func part2(input []string) string {
	foods := parseFoods(input)
	allergens := processFoods(foods)
	return report(allergens)
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
