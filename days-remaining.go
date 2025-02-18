package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type DueItem struct {
	DueDate time.Time
	Name    string
}

func exeName() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
	}
	exeName := filepath.Base(exePath)

	return exeName
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// parseDueDates converts a slice of "YYYY-MM-DD==name" strings into structured data
func parseDueDates(pairs []string) ([]DueItem, error) {
	var result []DueItem
	dateFormat := "2006-01-02" // Go's reference time format for YYYY-MM-DD

	for _, pair := range pairs {
		parts := strings.SplitN(pair, "==", 2) // Split into at most 2 parts
		if len(parts) != 2 {
			fmt.Printf("Skipping invalid entry: %s\n", pair)
			continue
		}

		// Parse the due date string into a time.Time object
		dueDate, err := time.Parse(dateFormat, parts[0])
		if err != nil {
			fmt.Printf("Skipping invalid date format: %s\n", parts[0])
			continue
		}

		result = append(result, DueItem{
			DueDate: dueDate,
			Name:    parts[1],
		})
	}

	return result, nil
}

// daysUntil calculates the number of days from today to the given date
func daysUntil(targetDate time.Time) int {
	today := time.Now().Truncate(24 * time.Hour) // Remove time part for accurate comparison
	target := targetDate.Truncate(24 * time.Hour)

	days := int(target.Sub(today).Hours() / 24) // Convert duration to days
	return days
}

func sortDueItems(items []DueItem) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].DueDate.Before(items[j].DueDate) // Ascending order
	})
}

func main() {
	exename := exeName()

	if len(os.Args) < 2 {
		fmt.Println("Usage:", exename, "<filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	lines, err := readLines(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parsing
	parsedData, err := parseDueDates(lines)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}

	sortDueItems(parsedData)

	// Set colours for prettier output
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	// Output results
	for _, item := range parsedData {
		daysRemaining := daysUntil(item.DueDate)

		if daysRemaining >= 10 {
			fmt.Print(green(fmt.Sprintf("%s - Remaining: %d\n", item.Name, daysRemaining)))
		} else if daysRemaining > 0 {
			fmt.Print(yellow(fmt.Sprintf("%s - Remaining: %d\n", item.Name, daysRemaining)))
		} else {
			fmt.Print(red(fmt.Sprintf("%s - Remaining: %d\n", item.Name, daysRemaining)))
		}

		// fmt.Printf("%s - Remaining: %d\n", item.Name, daysRemaining) # use this if colours are disabled
	}
}
