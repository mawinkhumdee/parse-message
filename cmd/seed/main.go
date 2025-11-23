package main

import (
	"context"
	"fmt"
	"log"
	"parse-message/config"
	"parse-message/db"
	"parse-message/model"
)

func main() {
	cfg := config.LoadConfig()
	database := db.New(cfg.DB)
	ctx := context.Background()

	defaultFields := []model.StandardField{
		// Expense
		{Key: "expense.amount", Type: "number", Category: "expense", Description: "Amount of expense"},
		{Key: "expense.currency", Type: "string", Category: "expense", Description: "Currency of expense"},
		{Key: "expense.category_main", Type: "string", Category: "expense", Description: "Main category of expense"},
		{Key: "expense.category_det", Type: "string", Category: "expense", Description: "Detailed category of expense"},
		{Key: "expense.date", Type: "date", Category: "expense", Description: "Date of expense"},
		{Key: "expense.datetime", Type: "datetime", Category: "expense", Description: "Datetime of expense"},
		{Key: "expense.note", Type: "string", Category: "expense", Description: "Note for expense"},

		// Income
		{Key: "income.amount", Type: "number", Category: "income", Description: "Amount of income"},
		{Key: "income.currency", Type: "string", Category: "income", Description: "Currency of income"},
		{Key: "income.source", Type: "string", Category: "income", Description: "Source of income"},
		{Key: "income.date", Type: "date", Category: "income", Description: "Date of income"},
		{Key: "income.note", Type: "string", Category: "income", Description: "Note for income"},

		// Schedule
		{Key: "schedule.datetime", Type: "datetime", Category: "schedule", Description: "Datetime of event"},
		{Key: "schedule.date", Type: "date", Category: "schedule", Description: "Date of event"},
		{Key: "schedule.time", Type: "string", Category: "schedule", Description: "Time of event"},
		{Key: "schedule.title", Type: "string", Category: "schedule", Description: "Title of event"},
		{Key: "schedule.location", Type: "string", Category: "schedule", Description: "Location of event"},
		{Key: "schedule.people", Type: "array", Category: "schedule", Description: "People involved in event"},
		{Key: "schedule.note", Type: "string", Category: "schedule", Description: "Note for event"},

		// Todo
		{Key: "todo.title", Type: "string", Category: "todo", Description: "Title of todo task"},
		{Key: "todo.description", Type: "string", Category: "todo", Description: "Description of todo task"},
		{Key: "todo.due_date", Type: "date", Category: "todo", Description: "Due date of todo task"},
		{Key: "todo.due_datetime", Type: "datetime", Category: "todo", Description: "Due datetime of todo task"},
		{Key: "todo.priority", Type: "string", Category: "todo", Description: "Priority of todo task"},
		{Key: "todo.status", Type: "string", Category: "todo", Description: "Status of todo task"},

		// Note
		{Key: "note.text", Type: "string", Category: "note", Description: "Content of note"},
		{Key: "note.tags", Type: "array", Category: "note", Description: "Tags for note"},
		{Key: "note.topic", Type: "array", Category: "note", Description: "Topics of note"},

		// Emotion
		{Key: "emotion.primary", Type: "string", Category: "emotion", Description: "Primary emotion"},
		{Key: "emotion.intensity", Type: "number", Category: "emotion", Description: "Intensity of emotion (1-10)"},
		{Key: "emotion.sentiment", Type: "string", Category: "emotion", Description: "Sentiment (positive/neutral/negative)"},
		{Key: "emotion.stress", Type: "number", Category: "emotion", Description: "Stress level (1-10)"},
		{Key: "emotion.energy", Type: "number", Category: "emotion", Description: "Energy level (1-10)"},
		{Key: "emotion.reason", Type: "string", Category: "emotion", Description: "Reason for emotion"},

		// Journal
		{Key: "journal.summary", Type: "string", Category: "journal", Description: "Summary of journal entry"},
		{Key: "journal.topics", Type: "array", Category: "journal", Description: "Topics of journal entry"},
		{Key: "journal.tags", Type: "array", Category: "journal", Description: "Tags for journal entry"},
		{Key: "journal.date", Type: "date", Category: "journal", Description: "Date of journal entry"},
		{Key: "journal.datetime", Type: "datetime", Category: "journal", Description: "Datetime of journal entry"},

		// Health
		{Key: "health.sleep_hours", Type: "number", Category: "health", Description: "Hours of sleep"},
		{Key: "health.water_intake_ml", Type: "number", Category: "health", Description: "Water intake in ml"},
		{Key: "health.exercise_type", Type: "string", Category: "health", Description: "Type of exercise"},
		{Key: "health.exercise_minutes", Type: "number", Category: "health", Description: "Minutes of exercise"},

		// Item
		{Key: "item.name", Type: "string", Category: "item", Description: "Name of item"},
		{Key: "item.quantity", Type: "number", Category: "item", Description: "Quantity of item"},
		{Key: "item.price", Type: "number", Category: "item", Description: "Price of item"},
		{Key: "item.category", Type: "string", Category: "item", Description: "Category of item"},
		{Key: "item.warranty_end", Type: "date", Category: "item", Description: "Warranty end date"},
	}

	for _, field := range defaultFields {
		field.Version = 1
		field.Intents = []string{field.Category}
		if err := database.Standardfield.Upsert(ctx, field); err != nil {
			log.Printf("Failed to upsert %s: %v", field.Key, err)
		} else {
			fmt.Printf("Upserted %s\n", field.Key)
		}
	}
	fmt.Println("Seeding completed.")
}
