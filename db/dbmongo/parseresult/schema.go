package parseresult

import "go.mongodb.org/mongo-driver/bson/primitive"

type parseResultSchema struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty"`
	UserID                string              `bson:"user_id,omitempty"`
	MessageID             string              `bson:"message_id,omitempty"`
	Intent                string              `bson:"intent"`
	ExpenseAmount         *float64            `bson:"expense_amount,omitempty"`
	ExpenseCurrency       *string             `bson:"expense_currency,omitempty"`
	ExpenseCategoryMain   *string             `bson:"expense_category_main,omitempty"`
	ExpenseCategoryDet    *string             `bson:"expense_category_det,omitempty"`
	ExpenseDate           *primitive.DateTime `bson:"expense_date,omitempty"`
	ExpenseDatetime       *primitive.DateTime `bson:"expense_datetime,omitempty"`
	ExpenseNote           *string             `bson:"expense_note,omitempty"`
	IncomeAmount          *float64            `bson:"income_amount,omitempty"`
	IncomeCurrency        *string             `bson:"income_currency,omitempty"`
	IncomeSource          *string             `bson:"income_source,omitempty"`
	IncomeDate            *primitive.DateTime `bson:"income_date,omitempty"`
	IncomeNote            *string             `bson:"income_note,omitempty"`
	ScheduleDatetime      *primitive.DateTime `bson:"schedule_datetime,omitempty"`
	ScheduleDate          *primitive.DateTime `bson:"schedule_date,omitempty"`
	ScheduleTime          *string             `bson:"schedule_time,omitempty"`
	ScheduleTitle         *string             `bson:"schedule_title,omitempty"`
	ScheduleLocation      *string             `bson:"schedule_location,omitempty"`
	SchedulePeople        []string            `bson:"schedule_people,omitempty"`
	ScheduleNote          *string             `bson:"schedule_note,omitempty"`
	TodoTitle             *string             `bson:"todo_title,omitempty"`
	TodoDescription       *string             `bson:"todo_description,omitempty"`
	TodoDueDate           *primitive.DateTime `bson:"todo_due_date,omitempty"`
	TodoDueDatetime       *primitive.DateTime `bson:"todo_due_datetime,omitempty"`
	TodoPriority          *string             `bson:"todo_priority,omitempty"`
	TodoStatus            *string             `bson:"todo_status,omitempty"`
	NoteText              *string             `bson:"note_text,omitempty"`
	NoteTags              []string            `bson:"note_tags,omitempty"`
	NoteTopic             []string            `bson:"note_topic,omitempty"`
	EmotionPrimary        *string             `bson:"emotion_primary,omitempty"`
	EmotionIntensity      *int                `bson:"emotion_intensity,omitempty"`
	EmotionSentiment      *string             `bson:"emotion_sentiment,omitempty"`
	EmotionStress         *int                `bson:"emotion_stress,omitempty"`
	EmotionEnergy         *int                `bson:"emotion_energy,omitempty"`
	EmotionReason         *string             `bson:"emotion_reason,omitempty"`
	JournalSummary        *string             `bson:"journal_summary,omitempty"`
	JournalTopics         []string            `bson:"journal_topics,omitempty"`
	JournalTags           []string            `bson:"journal_tags,omitempty"`
	JournalDate           *primitive.DateTime `bson:"journal_date,omitempty"`
	JournalDatetime       *primitive.DateTime `bson:"journal_datetime,omitempty"`
	HealthSleepHours      *float64            `bson:"health_sleep_hours,omitempty"`
	HealthWaterIntakeML   *float64            `bson:"health_water_intake_ml,omitempty"`
	HealthExerciseType    *string             `bson:"health_exercise_type,omitempty"`
	HealthExerciseMinutes *int                `bson:"health_exercise_minutes,omitempty"`
	ItemName              *string             `bson:"item_name,omitempty"`
	ItemQuantity          *int                `bson:"item_quantity,omitempty"`
	ItemPrice             *float64            `bson:"item_price,omitempty"`
	ItemCategory          *string             `bson:"item_category,omitempty"`
	ItemWarrantyEnd       *primitive.DateTime `bson:"item_warranty_end,omitempty"`
	Fields                []fieldValueSchema  `bson:"fields"`
	RawText               string              `bson:"raw_text"`
	Confidence            float64             `bson:"confidence"`
	CreatedAt             primitive.DateTime  `bson:"created_at"`
}

type fieldValueSchema struct {
	Key   string      `bson:"key"`
	Type  string      `bson:"type"`
	Value interface{} `bson:"value"`
}
