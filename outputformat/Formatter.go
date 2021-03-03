package outputformat

type FormatItem struct {
	name      string
	itemType  string
	value     string
	modifiers []string
	isLast    bool
}

type Formatter interface {
	Format(*FormatItem) string
	Close(*FormatItem) string
}
