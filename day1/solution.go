package main

import (
  "fmt"
	"os"
  "sort"
  "math"
  "strconv"
	"strings"
)

func parseFile(filename string) [][]int {
  lists, _ := os.ReadFile(filename)
  splitLists := strings.Split(string(lists), "\n");

  parsedLists := make([][]int, 2)

  for _, line := range splitLists {
    nums := strings.Fields(line)
    if len(nums) == 0 {
      continue
    }
    list1Item, _ := strconv.Atoi(nums[0])
    list2Item, _ := strconv.Atoi(nums[1])

    parsedLists[0] = append(parsedLists[0], list1Item)
    parsedLists[1] = append(parsedLists[1], list2Item)
  }

  return parsedLists
}

/*
Distance is the sum of all absolute value difference between the 
two lists essentially after they've been sorted
*/
func calculateDistance(parsedLists [][]int) int {
  var distance float64 = 0

  for i := range parsedLists[0] {
    list1Num := parsedLists[0][i]
    list2Num := parsedLists[1][i]

    distance += math.Abs(float64(list1Num) - float64(list2Num))
  }
  return int(distance)
}


func _countOccurences(list []int, testNum int) int {
  count := 0
  for _, num := range list{
    if num == testNum {
      count+=1
    }
  }
  return count
}
/*
Similarity score is the sum of all the products of the numbers
in the left list with their occurences in the right list
*/
func calculateSimilarityScore(parsedLists [][]int) int {
  similarityScore := 0

  for i := range parsedLists[0]{
    list1Num := parsedLists[0][i]
    
    similarityScore += list1Num * _countOccurences(parsedLists[1], list1Num)
  }

  return similarityScore
}

func main() {
  parsedLists := parseFile("input.txt")
  sort.Ints(parsedLists[0])
  sort.Ints(parsedLists[1])

  distance := calculateDistance(parsedLists)

  similarityScore := calculateSimilarityScore(parsedLists)

  fmt.Printf("Distance between lists: %d\nSimilarity score: %d", distance, similarityScore)


}


