// Code generated by "stringer -type=Exp"; DO NOT EDIT.

package symbols

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Undefined-0]
	_ = x[Unexpected-2]
	_ = x[InvalidHyphen-3]
	_ = x[SubExpressionEnd-4]
	_ = x[ObjectEnd-5]
	_ = x[ArrayEnd-6]
	_ = x[DataValues-7]
	_ = x[Bareword-8]
	_ = x[SubExpressionBegin-9]
	_ = x[ObjectBegin-10]
	_ = x[ArrayBegin-11]
	_ = x[QuoteSingle-12]
	_ = x[QuoteDouble-13]
	_ = x[Number-14]
	_ = x[Boolean-15]
	_ = x[Null-16]
	_ = x[Operations-17]
	_ = x[Assign-18]
	_ = x[AssignAndAdd-19]
	_ = x[AssignAndSubtract-20]
	_ = x[AssignAndDivide-21]
	_ = x[AssignAndMultiply-22]
	_ = x[EqualTo-23]
	_ = x[NotEqualTo-24]
	_ = x[Like-25]
	_ = x[NotLike-26]
	_ = x[Regexp-27]
	_ = x[NotRegexp-28]
	_ = x[GreaterThan-29]
	_ = x[GreaterThanOrEqual-30]
	_ = x[LessThan-31]
	_ = x[LessThanOrEqual-32]
	_ = x[Add-33]
	_ = x[Subtract-34]
	_ = x[Multiply-35]
	_ = x[Divide-36]
}

const (
	_Exp_name_0 = "Undefined"
	_Exp_name_1 = "UnexpectedInvalidHyphenSubExpressionEndObjectEndArrayEndDataValuesBarewordSubExpressionBeginObjectBeginArrayBeginQuoteSingleQuoteDoubleNumberBooleanNullOperationsAssignAssignAndAddAssignAndSubtractAssignAndDivideAssignAndMultiplyEqualToNotEqualToLikeNotLikeRegexpNotRegexpGreaterThanGreaterThanOrEqualLessThanLessThanOrEqualAddSubtractMultiplyDivide"
)

var (
	_Exp_index_1 = [...]uint16{0, 10, 23, 39, 48, 56, 66, 74, 92, 103, 113, 124, 135, 141, 148, 152, 162, 168, 180, 197, 212, 229, 236, 246, 250, 257, 263, 272, 283, 301, 309, 324, 327, 335, 343, 349}
)

func (i Exp) String() string {
	switch {
	case i == 0:
		return _Exp_name_0
	case 2 <= i && i <= 36:
		i -= 2
		return _Exp_name_1[_Exp_index_1[i]:_Exp_index_1[i+1]]
	default:
		return "Exp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
