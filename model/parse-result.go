package model

import "time"

type ParseResult struct {
	ID                    string
	UserID                string
	MessageID             string
	Intent                string // "expense" | "schedule" | "emotion" | ...
	ExpenseAmount         *float64
	ExpenseCurrency       *string
	ExpenseCategoryMain   *string
	ExpenseCategoryDet    *string
	ExpenseDate           *time.Time
	ExpenseDatetime       *time.Time
	ExpenseNote           *string
	IncomeAmount          *float64
	IncomeCurrency        *string
	IncomeSource          *string
	IncomeDate            *time.Time
	IncomeNote            *string
	ScheduleDatetime      *time.Time
	ScheduleDate          *time.Time
	ScheduleTime          *string
	ScheduleTitle         *string
	ScheduleLocation      *string
	SchedulePeople        []string
	ScheduleNote          *string
	TodoTitle             *string
	TodoDescription       *string
	TodoDueDate           *time.Time
	TodoDueDatetime       *time.Time
	TodoPriority          *string
	TodoStatus            *string
	NoteText              *string
	NoteTags              []string
	NoteTopic             []string
	EmotionPrimary        *string // tired / happy / stressed
	EmotionIntensity      *int    // 1–10
	EmotionSentiment      *string // positive / neutral / negative
	EmotionStress         *int    // 1–10
	EmotionEnergy         *int    // 1–10
	EmotionReason         *string
	JournalSummary        *string
	JournalTopics         []string
	JournalTags           []string
	JournalDate           *time.Time
	JournalDatetime       *time.Time
	HealthSleepHours      *float64
	HealthWaterIntakeML   *float64
	HealthExerciseType    *string
	HealthExerciseMinutes *int
	ItemName              *string
	ItemQuantity          *int
	ItemPrice             *float64
	ItemCategory          *string
	ItemWarrantyEnd       *time.Time
	Fields                []FieldValue
	RawText               string
	Confidence            float64
	CreatedAt             time.Time
}

type FieldValue struct {
	Key   string
	Type  string
	Value any
}
