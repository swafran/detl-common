package parsers

//Parser handles various incoming data formats
type Parser interface {
	Parse(input string) interface{}
}
