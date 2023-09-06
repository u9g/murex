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
	_ = x[QuoteParenthesis-14]
	_ = x[Number-15]
	_ = x[Boolean-16]
	_ = x[Null-17]
	_ = x[Scalar-18]
	_ = x[Calculated-19]
	_ = x[Operations-20]
	_ = x[Assign-21]
	_ = x[AssignUpdate-22]
	_ = x[AssignAndAdd-23]
	_ = x[AssignAndSubtract-24]
	_ = x[AssignAndDivide-25]
	_ = x[AssignAndMultiply-26]
	_ = x[AssignAndMerge-27]
	_ = x[LogicalOr-28]
	_ = x[LogicalAnd-29]
	_ = x[EqualTo-30]
	_ = x[NotEqualTo-31]
	_ = x[Like-32]
	_ = x[NotLike-33]
	_ = x[Regexp-34]
	_ = x[NotRegexp-35]
	_ = x[GreaterThan-36]
	_ = x[GreaterThanOrEqual-37]
	_ = x[LessThan-38]
	_ = x[LessThanOrEqual-39]
	_ = x[Add-40]
	_ = x[Subtract-41]
	_ = x[MergeInto-42]
	_ = x[Multiply-43]
	_ = x[Divide-44]
}

const (
	_Exp_name_0 = "Undefined"
	_Exp_name_1 = "UnexpectedInvalidHyphenSubExpressionEndObjectEndArrayEndDataValuesBarewordSubExpressionBeginObjectBeginArrayBeginQuoteSingleQuoteDoubleQuoteParenthesisNumberBooleanNullScalarCalculatedOperationsAssignAssignUpdateAssignAndAddAssignAndSubtractAssignAndDivideAssignAndMultiplyAssignAndMergeLogicalOrLogicalAndEqualToNotEqualToLikeNotLikeRegexpNotRegexpGreaterThanGreaterThanOrEqualLessThanLessThanOrEqualAddSubtractMergeIntoMultiplyDivide"
)

var (
	_Exp_index_1 = [...]uint16{0, 10, 23, 39, 48, 56, 66, 74, 92, 103, 113, 124, 135, 151, 157, 164, 168, 174, 184, 194, 200, 212, 224, 241, 256, 273, 287, 296, 306, 313, 323, 327, 334, 340, 349, 360, 378, 386, 401, 404, 412, 421, 429, 435}
)

func (i Exp) String() string {
	switch {
	case i == 0:
		return _Exp_name_0
	case 2 <= i && i <= 44:
		i -= 2
		return _Exp_name_1[_Exp_index_1[i]:_Exp_index_1[i+1]]
	default:
		return "Exp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
