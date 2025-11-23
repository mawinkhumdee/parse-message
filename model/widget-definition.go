package model

import "time"

type WidgetDefinition struct {
	ID          string
	WidgetKey   string
	Name        string
	Description string
	Type        string // "list" | "summary_card" | "chart" | "calendar"
	Intent      string
	DataSource  DataSource
	Aggregation *Aggregation
	Chart       *ChartConfig
	Mapping     map[string]string
	UISchema    map[string]any
	IsActive    bool
	Version     int
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SortField struct {
	Field string
	Order string // "asc" | "desc"
}

// TimeRange → preset หรือ custom ช่วงเวลา
type TimeRange struct {
	Field  string     // eg. "expense_date"
	Preset string     // "this_month" | "last_month" | "this_week"
	From   *time.Time // optional
	To     *time.Time // optional
}

// DataSource → ควบคุมว่า widget จะดึงข้อมูลจากไหน
type DataSource struct {
	Collection string
	Filter     map[string]any
	TimeRange  *TimeRange
	Sort       []SortField
	Limit      int
}

// Metric → sum/avg/count/max/min
type Metric struct {
	Name   string
	Field  string
	Agg    string // "sum" | "avg" | "count" | "max" | "min"
	Label  string
	Format string // optional UI hint
}

// Aggregation → สำหรับ summary widget หรือ chart
type Aggregation struct {
	Metrics []Metric
	GroupBy string // eg. "expense_category_main"
}

type ChartConfig struct {
	ChartType string // "line" | "bar" | "pie" | "area"
	XField    string // eg. "expense_date"
	XBucket   string // "day" | "week" | "month" | "year"
}
