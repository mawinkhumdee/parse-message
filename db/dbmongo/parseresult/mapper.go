package parseresult

import (
	"parse-message/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func fromModel(pr *model.ParseResult) (*parseResultSchema, error) {
	var id primitive.ObjectID
	var err error

	if pr.ID != "" {
		id, err = primitive.ObjectIDFromHex(pr.ID)
		if err != nil {
			return nil, err
		}
	}

	// Convert fields
	fields := make([]fieldValueSchema, len(pr.Fields))
	for i, f := range pr.Fields {
		fields[i] = fieldValueSchema{
			Key:   f.Key,
			Type:  f.Type,
			Value: f.Value,
		}
	}

	return &parseResultSchema{
		ID:                    id,
		UserID:                pr.UserID,
		MessageID:             pr.MessageID,
		Intent:                pr.Intent,
		ExpenseAmount:         pr.ExpenseAmount,
		ExpenseCurrency:       pr.ExpenseCurrency,
		ExpenseCategoryMain:   pr.ExpenseCategoryMain,
		ExpenseCategoryDet:    pr.ExpenseCategoryDet,
		ExpenseDate:           timeToDateTime(pr.ExpenseDate),
		ExpenseDatetime:       timeToDateTime(pr.ExpenseDatetime),
		ExpenseNote:           pr.ExpenseNote,
		IncomeAmount:          pr.IncomeAmount,
		IncomeCurrency:        pr.IncomeCurrency,
		IncomeSource:          pr.IncomeSource,
		IncomeDate:            timeToDateTime(pr.IncomeDate),
		IncomeNote:            pr.IncomeNote,
		ScheduleDatetime:      timeToDateTime(pr.ScheduleDatetime),
		ScheduleDate:          timeToDateTime(pr.ScheduleDate),
		ScheduleTime:          pr.ScheduleTime,
		ScheduleTitle:         pr.ScheduleTitle,
		ScheduleLocation:      pr.ScheduleLocation,
		SchedulePeople:        pr.SchedulePeople,
		ScheduleNote:          pr.ScheduleNote,
		TodoTitle:             pr.TodoTitle,
		TodoDescription:       pr.TodoDescription,
		TodoDueDate:           timeToDateTime(pr.TodoDueDate),
		TodoDueDatetime:       timeToDateTime(pr.TodoDueDatetime),
		TodoPriority:          pr.TodoPriority,
		TodoStatus:            pr.TodoStatus,
		NoteText:              pr.NoteText,
		NoteTags:              pr.NoteTags,
		NoteTopic:             pr.NoteTopic,
		EmotionPrimary:        pr.EmotionPrimary,
		EmotionIntensity:      pr.EmotionIntensity,
		EmotionSentiment:      pr.EmotionSentiment,
		EmotionStress:         pr.EmotionStress,
		EmotionEnergy:         pr.EmotionEnergy,
		EmotionReason:         pr.EmotionReason,
		JournalSummary:        pr.JournalSummary,
		JournalTopics:         pr.JournalTopics,
		JournalTags:           pr.JournalTags,
		JournalDate:           timeToDateTime(pr.JournalDate),
		JournalDatetime:       timeToDateTime(pr.JournalDatetime),
		HealthSleepHours:      pr.HealthSleepHours,
		HealthWaterIntakeML:   pr.HealthWaterIntakeML,
		HealthExerciseType:    pr.HealthExerciseType,
		HealthExerciseMinutes: pr.HealthExerciseMinutes,
		ItemName:              pr.ItemName,
		ItemQuantity:          pr.ItemQuantity,
		ItemPrice:             pr.ItemPrice,
		ItemCategory:          pr.ItemCategory,
		ItemWarrantyEnd:       timeToDateTime(pr.ItemWarrantyEnd),
		Fields:                fields,
		RawText:               pr.RawText,
		Confidence:            pr.Confidence,
		CreatedAt:             primitive.NewDateTimeFromTime(pr.CreatedAt),
	}, nil
}

func toModel(doc *parseResultSchema) model.ParseResult {
	var id string
	if !doc.ID.IsZero() {
		id = doc.ID.Hex()
	}

	// Convert fields
	fields := make([]model.FieldValue, len(doc.Fields))
	for i, f := range doc.Fields {
		fields[i] = model.FieldValue{
			Key:   f.Key,
			Type:  f.Type,
			Value: f.Value,
		}
	}

	return model.ParseResult{
		ID:                    id,
		UserID:                doc.UserID,
		MessageID:             doc.MessageID,
		Intent:                doc.Intent,
		ExpenseAmount:         doc.ExpenseAmount,
		ExpenseCurrency:       doc.ExpenseCurrency,
		ExpenseCategoryMain:   doc.ExpenseCategoryMain,
		ExpenseCategoryDet:    doc.ExpenseCategoryDet,
		ExpenseDate:           dateTimeToTime(doc.ExpenseDate),
		ExpenseDatetime:       dateTimeToTime(doc.ExpenseDatetime),
		ExpenseNote:           doc.ExpenseNote,
		IncomeAmount:          doc.IncomeAmount,
		IncomeCurrency:        doc.IncomeCurrency,
		IncomeSource:          doc.IncomeSource,
		IncomeDate:            dateTimeToTime(doc.IncomeDate),
		IncomeNote:            doc.IncomeNote,
		ScheduleDatetime:      dateTimeToTime(doc.ScheduleDatetime),
		ScheduleDate:          dateTimeToTime(doc.ScheduleDate),
		ScheduleTime:          doc.ScheduleTime,
		ScheduleTitle:         doc.ScheduleTitle,
		ScheduleLocation:      doc.ScheduleLocation,
		SchedulePeople:        doc.SchedulePeople,
		ScheduleNote:          doc.ScheduleNote,
		TodoTitle:             doc.TodoTitle,
		TodoDescription:       doc.TodoDescription,
		TodoDueDate:           dateTimeToTime(doc.TodoDueDate),
		TodoDueDatetime:       dateTimeToTime(doc.TodoDueDatetime),
		TodoPriority:          doc.TodoPriority,
		TodoStatus:            doc.TodoStatus,
		NoteText:              doc.NoteText,
		NoteTags:              doc.NoteTags,
		NoteTopic:             doc.NoteTopic,
		EmotionPrimary:        doc.EmotionPrimary,
		EmotionIntensity:      doc.EmotionIntensity,
		EmotionSentiment:      doc.EmotionSentiment,
		EmotionStress:         doc.EmotionStress,
		EmotionEnergy:         doc.EmotionEnergy,
		EmotionReason:         doc.EmotionReason,
		JournalSummary:        doc.JournalSummary,
		JournalTopics:         doc.JournalTopics,
		JournalTags:           doc.JournalTags,
		JournalDate:           dateTimeToTime(doc.JournalDate),
		JournalDatetime:       dateTimeToTime(doc.JournalDatetime),
		HealthSleepHours:      doc.HealthSleepHours,
		HealthWaterIntakeML:   doc.HealthWaterIntakeML,
		HealthExerciseType:    doc.HealthExerciseType,
		HealthExerciseMinutes: doc.HealthExerciseMinutes,
		ItemName:              doc.ItemName,
		ItemQuantity:          doc.ItemQuantity,
		ItemPrice:             doc.ItemPrice,
		ItemCategory:          doc.ItemCategory,
		ItemWarrantyEnd:       dateTimeToTime(doc.ItemWarrantyEnd),
		Fields:                fields,
		RawText:               doc.RawText,
		Confidence:            doc.Confidence,
		CreatedAt:             doc.CreatedAt.Time(),
	}
}

func timeToDateTime(t *time.Time) *primitive.DateTime {
	if t == nil {
		return nil
	}
	dt := primitive.NewDateTimeFromTime(*t)
	return &dt
}

func dateTimeToTime(dt *primitive.DateTime) *time.Time {
	if dt == nil {
		return nil
	}
	t := dt.Time()
	return &t
}
