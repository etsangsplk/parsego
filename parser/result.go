package parser

// ParseResult defines the values that can result from parsing.
type ParseResult interface {
	Matched() bool
	Result() interface{}
	TextRange() TextRange
	Error() error
}

// SuccessResult is returned by parsers on a successful parser.
type SuccessResult struct {
	textRange TextRange
	result    interface{}
}

// Matched returns whether the parser matched the input (true in this
// case).
func (r *SuccessResult) Matched() bool {
	return true
}

// Result retuns the parsed value (usually a string).
func (r *SuccessResult) Result() interface{} {
	return r.result
}

// TextRange returns the range of text parsed.
func (r *SuccessResult) TextRange() TextRange {
	return r.textRange
}

// Error returns the reason for failing (nil in this case).
func (r *SuccessResult) Error() error {
	return nil
}

// FailedResult is returned by parsers when they fail to parse the
// input.
type FailedResult struct {
	textRange TextRange
	err       error
}

// Matched returns whether the parser matched the input (false in this
// case).
func (r *FailedResult) Matched() bool {
	return false
}

// Result retuns the parsed value (always nil, in this case).
func (r *FailedResult) Result() interface{} {
	return nil
}

// TextRange returns the range of text parsed.
func (r *FailedResult) TextRange() TextRange {
	return r.textRange
}

// Error returns the reason for failing.
func (r *FailedResult) Error() error {
	return r.err
}
