package test

import (
	"github.com/gotestyourself/gotestyourself/assert"
	"testing"

	"github.com/aristotll/goutils/stringutils"
)

func TestIsEmpty(t *testing.T) {
	testStr := "999ss999asf2323323asd"

	out := stringutils.IsEmpty(testStr)
	if out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsAnyEmpty(t *testing.T) {

	out := stringutils.IsAnyEmpty("", "2")
	if !out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsNoneEmpty(t *testing.T) {

	out := stringutils.IsNoneEmpty("", "2")
	if out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsBlank(t *testing.T) {

	out := stringutils.IsBlank("")
	if !out {
		t.Error(out)
	}

	out = stringutils.IsBlank(" ")
	if !out {
		t.Error(out)
	}

	t.Log(out)
}

func TestIsNotBlank(t *testing.T) {

	out := stringutils.IsNotBlank("")
	if out {
		t.Error(out)
	}

	out = stringutils.IsNotBlank(" ")
	if out {
		t.Error(out)
	}

	out = stringutils.IsNotBlank(" test")
	if !out {
		t.Error(out)
	}
}

func TestIsAnyBlank(t *testing.T) {

	out := stringutils.IsAnyBlank(" ", "d")
	if !out {
		t.Error(out)
	}

	out = stringutils.IsAnyBlank("dd", "ee")
	if out {
		t.Error(out)
	}

	t.Log(out)
}

func TestTruncate(t *testing.T) {

	out := stringutils.Truncate("012345678", 3)
	assert.Assert(t, out == "012", out)

	out = stringutils.Truncate("012345678", 9)
	if out != "012345678" {
		t.Error(out)
	}

	out = stringutils.Truncate("012345678", 8)
	if out != "01234567" {
		t.Error(out)
	}

	out = stringutils.Truncate("012345678", 0)
	if out != "" {
		t.Error(out)
	}

	t.Log(out)
}

func TestTruncateFromWithMaxWith(t *testing.T) {

	out, _ := stringutils.TruncateFromWithMaxWith("0123456789", 3, 7)
	if out != "3456789" {
		t.Error(out)
	}

	t.Log(out)
}

func TestSubString(t *testing.T) {

	out, _ := stringutils.SubString("0123456789", 3)
	if out != "3456789" {
		t.Error(out)
	}

	t.Log(out)
}

func TestSubStringBetween(t *testing.T) {

	out, err := stringutils.SubStringBetween("0123456789", 0, 9)
	if out != "012345678" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	out, err = stringutils.SubStringBetween("120.00", 0, 2)
	if out != "12" {
		t.Error(out)
	}

	out, err = stringutils.SubStringBetween("120.00", 0, 0)
	if out != "" {
		t.Error(out)
	}

	t.Log(out)
}

func TestCharAt(t *testing.T) {

	out, err := stringutils.CharAt("0123456789", 2)
	if out != "2" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	out, err = stringutils.CharAt("01中国456789", 2)
	if out != "中" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	out, err = stringutils.CharAt("01中国456789", 4)
	if out != "4" {
		t.Error(out)
	}

	if err != nil {
		t.Error(err)
	}

	t.Log(out)
}

func TestIndexOf(t *testing.T) {

	out := stringutils.IndexOf("0123456789", "2")
	if out != 2 {
		t.Error(out)
	}

	out = stringutils.IndexOf("01中国456789", "中")
	if out != 2 {
		t.Error(out)
	}

	out = stringutils.IndexOf("01中国456789", "d中")
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.IndexOf("中国456789", "国")
	if out != 1 {
		t.Error(out)
	}

	t.Log(out)
}

func TestStripStart(t *testing.T) {

	out := stringutils.StripStart("0123456789", "d081d2d348")
	if out != "56789" {
		t.Error(out)
	}

	out = stringutils.StripStart("01中国456789", "01中国456789")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.StripStart("yxabc  ", "xyz")
	if out != "abc  " {
		t.Error(out)
	}

	t.Log(out)
}

func TestStripEnd(t *testing.T) {

	out := stringutils.StripEnd("0123456789", "d081d2d348")
	if out != "0123456789" {
		t.Error(out)
	}

	out = stringutils.StripEnd("01中国456789", "01中国456789")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.StripEnd("120.00", ".0")
	if out != "12" {
		t.Error(out)
	}

	t.Log(out)
}

func TestStripWithChar(t *testing.T) {
	out := stringutils.StripWithChar("  abcyx", "xyz")
	if out != "  abc" {
		t.Error(out)
	}
}

func TestStrip(t *testing.T) {

	out := stringutils.Strip("")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Strip("   ")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Strip("abc")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip("  abc")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip("abc  ")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip(" abc ")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Strip(" ab c ")
	if out != "ab c" {
		t.Error(out)
	}
}

func TestToCharArray(t *testing.T) {

	_1 := []string{" ", "a", "b", " ", "c", " "}
	out := stringutils.ToCharArray(" ab c ")
	if len(out) != len(_1) {
		t.Error(len(out), len(_1))
	}

	for i, v := range _1 {
		if v != _1[i] {
			t.Error(v, _1[i])
		}
	}

	_2 := []string{" ", "中", "国", "a"}
	out = stringutils.ToCharArray(" 中国a")
	if len(out) != len(_2) {
		t.Error(len(out), len(_2))
	}

	for i, v := range _2 {
		if v != _2[i] {
			t.Error(v, _2[i])
		}
	}
}

func TestRegionMatches(t *testing.T) {

	out := stringutils.RegionMatches(" ab c ", true, 0, " aB c ", 0, len(" ab c "))
	if !out {
		t.Error(out)
	}

	out = stringutils.RegionMatches(" ab c ", true, 0, "aB c ", 0, len(" ab c "))
	if out {
		t.Error(out)
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	out := stringutils.EqualsIgnoreCase(" ab c ", " aB c ")
	if !out {
		t.Error(out)
	}

	out = stringutils.EqualsIgnoreCase(" ab c ", " aBc ")
	if out {
		t.Error(out)
	}

}

func TestCompare(t *testing.T) {

	out := stringutils.Compare("a", "b")
	if out > 0 {
		t.Error(out)
	}

	out = stringutils.Compare("a", "a")
	if out != 0 {
		t.Error(out)
	}

	// the rune of char '中' is '20013'
	// and char '国' has rune '22269'
	out = stringutils.Compare("国", "中")
	if out < 0 {
		t.Error(out)
	}

	// the rune of char '中' is '20013'
	// and char '国' has rune '22269'
	out = stringutils.Compare("中国国", "中国中")
	if out < 0 {
		t.Error(out)
	}
}

func TestCompareIgnoreCase(t *testing.T) {

	out := stringutils.CompareIgnoreCase("ABC", "abc")
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.CompareIgnoreCase("ABC", "abcb")
	if out >= 0 {
		t.Error(out)
	}

	out = stringutils.CompareIgnoreCase("ABC国", "abc国")
	if out != 0 {
		t.Error(out)
	}
}

func TestEqualsAny(t *testing.T) {

	out := stringutils.EqualsAny("ABC", "abc", "abcd")
	if out {
		t.Error(out)
	}

	out = stringutils.EqualsAnyIgnoreCase("ABC", "abc", "abcd")
	if !out {
		t.Error(out)
	}

	out = stringutils.EqualsAny("ABC国", "abc国", "adfsf")
	if out {
		t.Error(out)
	}
}

func TestIndexOfFromIndex(t *testing.T) {
	out := stringutils.IndexOfFromIndex("abcdabcabc", "abc", 1)
	if out != 4 {
		t.Error(out)
	}

	out = stringutils.IndexOfFromIndex("abcdabcabc", "abce", 0)
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.IndexOfFromIndex("ABC国国国", "国", 2)
	if out != 3 {
		t.Error(out)
	}
}

func TestOrdinalIndexOf(t *testing.T) {

	out := stringutils.OrdinalIndexOf("abababab", "abab", 1, false)
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("abababab", "abab", 2, false)
	if out != 2 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("abababab", "abab", 3, false)
	if out != 4 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("abababab", "abab", 3, true)
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.OrdinalIndexOf("中国中国中国中国", "中国中国", 3, true)
	if out != 0 {
		t.Error(out)
	}
}

func TestContains(t *testing.T) {

	out := stringutils.Contains("ABC国国国", "国")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("ABC国国国", "21")
	if out != false {
		t.Error(out)
	}

	out = stringutils.Contains("", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("abc", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("abc", "a")
	if out != true {
		t.Error(out)
	}

	out = stringutils.Contains("abc", "z")
	if out != false {
		t.Error(out)
	}

}

func TestContainsIgnoreCase(t *testing.T) {

	out := stringutils.ContainsIgnoreCase("ABC国国国", "国")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "a")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "z")
	if out == true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "A")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsIgnoreCase("abc", "Z")
	if out == true {
		t.Error(out)
	}
}

func TestIsWhitespace(t *testing.T) {

	// space
	out := stringutils.IsWhitespace(" ")
	if out != true {
		t.Error(out)
	}

	// enter
	out = stringutils.IsWhitespace(`
`)
	if out != true {
		t.Error(out)
	}

	// tab
	out = stringutils.IsWhitespace("	")
	if out != true {
		t.Error(out)
	}

	// tab
	out = stringutils.IsWhitespace("d")
	if out == true {
		t.Error(out)
	}
}

func TestIndexOfAny(t *testing.T) {

	out := stringutils.IndexOfAny("", "ddd")
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.IndexOfAny("zzabyycdxx", "z", "a")
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.IndexOfAny("zzabyycdxx", "b", "y")
	if out != 3 {
		t.Error(out)
	}

	out = stringutils.IndexOfAny("中国Golang的粉丝", "黑", "丝")
	if out != 10 {
		t.Error(out)
	}
}

func TestContainsAny(t *testing.T) {

	out := stringutils.ContainsAny("", "d")
	if out == true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("zzabyycdxx", "z", "a")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("zzabyycdxx", "b", "y")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("zzabyycdxx", "z", "y")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("aba", "z")
	if out == true {
		t.Error(out)
	}

	out = stringutils.ContainsAny("中国Golang的粉丝", "黑", "丝")
	if out != true {
		t.Error(out)
	}
}

func TestIndexOfAnyBut(t *testing.T) {

	out := stringutils.IndexOfAnyBut("zzabyycdxx", "z", "a")
	if out != 3 {
		t.Error(out)
	}

	out = stringutils.IndexOfAnyBut("aba", "z")
	if out != 0 {
		t.Error(out)
	}

	out = stringutils.IndexOfAnyBut("aba", "a", "b")
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.IndexOfAnyBut("中国Golang的粉丝", "黑", "丝")
	if out != 0 {
		t.Error(out)
	}
}

func TestContainsOnly(t *testing.T) {
	out := stringutils.ContainsOnly("ab", "")
	if out != false {
		t.Error(out)
	}

	out = stringutils.ContainsOnly("abab", "a", "b", "c")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsOnly("ab1", "a", "b", "c")
	if out != false {
		t.Error(out)
	}

	out = stringutils.ContainsOnly("abz", "a", "b", "c")
	if out != false {
		t.Error(out)
	}
}

func TestContainsNone(t *testing.T) {
	out := stringutils.ContainsNone("ab", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsNone("abab", "x", "y", "z")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsNone("ab1", "x", "y", "z")
	if out != true {
		t.Error(out)
	}

	out = stringutils.ContainsNone("abz", "x", "y", "z")
	if out != false {
		t.Error(out)
	}
}

func TestLastIndexOfAny(t *testing.T) {
	out := stringutils.LastIndexOfAny("zzabyycdxx", "ab", "cd")
	if out != 6 {
		t.Error(out)
	}

	out = stringutils.LastIndexOfAny("zzabyycdxx", "cd", "ab")
	if out != 6 {
		t.Error(out)
	}

	out = stringutils.LastIndexOfAny("zzabyycdxx", "mn", "op")
	if out != -1 {
		t.Error(out)
	}

	out = stringutils.LastIndexOfAny("zzabyycdxx", "mn", "")
	if out != -1 {
		t.Error(out)
	}
}

func TestLeft(t *testing.T) {

	out := stringutils.Left("abc", 0)
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Left("abc", 2)
	if out != "ab" {
		t.Error(out)
	}

	out = stringutils.Left("abc", 4)
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Left("中国Golang的粉丝", 4)
	if out != "中国Go" {
		t.Error(out)
	}
}

func TestRight(t *testing.T) {

	out := stringutils.Right("abc", 0)
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Right("abc", 2)
	if out != "bc" {
		t.Error(out)
	}

	out = stringutils.Right("abc", 4)
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Right("中国Golang的粉丝", 4)
	if out != "g的粉丝" {
		t.Error(out)
	}
}

func TestMid(t *testing.T) {

	out := stringutils.Mid("abc", 0, 2)
	if out != "ab" {
		t.Error(out)
	}

	out = stringutils.Mid("abc", 0, 4)
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.Mid("abc", 2, 4)
	if out != "c" {
		t.Error(out)
	}

	out = stringutils.Mid("abc", 4, 2)
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Mid("abc", -2, 2)
	if out != "ab" {
		t.Error(out)
	}

	out = stringutils.Mid("中国Golang的粉丝", 2, 8)
	if out != "Golang的粉" {
		t.Error(out)
	}
}

func TestSubstringBefore(t *testing.T) {

	out := stringutils.SubstringBefore("abc", "a")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.SubstringBefore("abcba", "b")
	if out != "a" {
		t.Error(out)
	}

	out = stringutils.SubstringBefore("abc", "c")
	if out != "ab" {
		t.Error(out)
	}

	out = stringutils.SubstringBefore("abc", "d")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.SubstringBefore("中国Golang的粉丝", "丝")
	if out != "中国Golang的粉" {
		t.Error(out)
	}
}

func TestSubstringAfter(t *testing.T) {

	out := stringutils.SubstringAfter("abc", "a")
	if out != "bc" {
		t.Error(out)
	}

	out = stringutils.SubstringAfter("abcba", "b")
	if out != "cba" {
		t.Error(out)
	}
	out = stringutils.SubstringAfter("abc", "c")
	if out != "" {
		t.Error(out)
	}
	out = stringutils.SubstringAfter("abc", "d")
	if out != "" {
		t.Error(out)
	}
	out = stringutils.SubstringAfter("abc", "")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.SubstringAfter("中国Golang的粉丝", "国")
	if out != "Golang的粉丝" {
		t.Error(out)
	}

}

func TestIsNumeric(t *testing.T) {

	out := stringutils.IsNumeric("")
	if out != false {
		t.Error(out)
	}

	out = stringutils.IsNumeric("  ")
	if out != false {
		t.Error(out)
	}

	out = stringutils.IsNumeric("123")
	if out != true {
		t.Error(out)
	}

	out = stringutils.IsNumeric("12 3")
	if out != false {
		t.Error(out)
	}

	out = stringutils.IsNumeric("ab2c")
	if out != false {
		t.Error(out)
	}

	out = stringutils.IsNumeric("12-3")
	if out != false {
		t.Error(out)
	}
	out = stringutils.IsNumeric("12.3")
	if out != false {
		t.Error(out)
	}
	out = stringutils.IsNumeric("-123")
	if out != false {
		t.Error(out)
	}
	out = stringutils.IsNumeric("+123")
	if out != false {
		t.Error(out)
	}
}

func TestStartsWith(t *testing.T) {

	out := stringutils.StartsWith("", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.StartsWith("", "abc")
	if out != false {
		t.Error(out)
	}

	out = stringutils.StartsWith("abcdef", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.StartsWith("abcdef", "abc")
	if out != true {
		t.Error(out)
	}

	out = stringutils.StartsWith("ABCDEF", "abc")
	if out != false {
		t.Error(out)
	}

	out = stringutils.StartsWith("中国汉字ABCDEF", "中国汉字")
	if out != true {
		t.Error(out)
	}

}

func TestJoin(t *testing.T) {

	out := stringutils.Join([]string{}, "")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Join(nil, "abc")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.Join([]string{"1", "2", "3"}, ";")
	if out != "1;2;3" {
		t.Error(out)
	}

	out = stringutils.Join([]string{"1", "2", "3"}, "")
	if out != "123" {
		t.Error(out)
	}
}

func TestDeleteWhitespace(t *testing.T) {

	out := stringutils.DeleteWhitespace("")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.DeleteWhitespace("abc")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.DeleteWhitespace("   ab  c  ")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.DeleteWhitespace("   中国  汉 字  ")
	if out != "中国汉字" {
		t.Error(out)
	}
}

func TestRemoveStart(t *testing.T) {

	out := stringutils.RemoveStart("", "asdfsafd")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.RemoveStart("www.domain.com", "www.")
	if out != "domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveStart("domain.com", "www.")
	if out != "domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveStart("www.domain.com", "domain")
	if out != "www.domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveStart("abc", "")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.RemoveStart("中国汉字abc", "中国汉字")
	if out != "abc" {
		t.Error(out)
	}
}

func TestRemoveStartIgnoreCase(t *testing.T) {

	out := stringutils.RemoveStartIgnoreCase("", "asdfsafd")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.RemoveStartIgnoreCase("www.domain.com", "www.")
	if out != "domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveStartIgnoreCase("www.domain.com", "WWW.")
	if out != "domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveStartIgnoreCase("domain.com", "www.")
	if out != "domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveStartIgnoreCase("www.domain.com", "domain")
	if out != "www.domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveStartIgnoreCase("abc", "")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.RemoveStartIgnoreCase("中国汉字abc", "中国汉字")
	if out != "abc" {
		t.Error(out)
	}

}

func TestEndsWith(t *testing.T) {

	out := stringutils.EndsWith("", "def")
	if out != false {
		t.Error(out)
	}

	out = stringutils.EndsWith("abcdef", "def")
	if out != true {
		t.Error(out)
	}

	out = stringutils.EndsWith("ABCDEF", "def")
	if out != false {
		t.Error(out)
	}

	out = stringutils.EndsWith("ABCDEF", "cde")
	if out != false {
		t.Error(out)
	}

	out = stringutils.EndsWith("ABCDEF", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.EndsWith("ABCDEF中国汉字", "中国汉字")
	if out != true {
		t.Error(out)
	}
}

func TestEndsWithIgnoreCase(t *testing.T) {

	out := stringutils.EndsWithIgnoreCase("", "def")
	if out != false {
		t.Error(out)
	}

	out = stringutils.EndsWithIgnoreCase("abcdef", "")
	if out != true {
		t.Error(out)
	}

	out = stringutils.EndsWithIgnoreCase("abcdef", "def")
	if out != true {
		t.Error(out)
	}

	out = stringutils.EndsWithIgnoreCase("ABCDEF", "def")
	if out != true {
		t.Error(out)
	}

	out = stringutils.EndsWithIgnoreCase("ABCDEF", "cde")
	if out != false {
		t.Error(out)
	}
}

func TestRemoveEnd(t *testing.T) {

	out := stringutils.RemoveEnd("", "def")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.RemoveEnd("www.domain.com", ".com.")
	if out != "www.domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveEnd("www.domain.com", ".com")
	if out != "www.domain" {
		t.Error(out)
	}

	out = stringutils.RemoveEnd("www.domain.com", "domain")
	if out != "www.domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveEnd("abc", "")
	if out != "abc" {
		t.Error(out)
	}
}

func TestRemoveEndIgnoreCase(t *testing.T) {

	out := stringutils.RemoveEndIgnoreCase("", "def")
	if out != "" {
		t.Error(out)
	}

	out = stringutils.RemoveEndIgnoreCase("www.domain.com", ".com.")
	if out != "www.domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveEndIgnoreCase("www.domain.com", ".com")
	if out != "www.domain" {
		t.Error(out)
	}

	out = stringutils.RemoveEndIgnoreCase("www.domain.com", "domain")
	if out != "www.domain.com" {
		t.Error(out)
	}

	out = stringutils.RemoveEndIgnoreCase("abc", "")
	if out != "abc" {
		t.Error(out)
	}

	out = stringutils.RemoveEndIgnoreCase("www.domain.com", ".COM")
	if out != "www.domain" {
		t.Error(out)
	}

	out = stringutils.RemoveEndIgnoreCase("www.domain.COM", ".com")
	if out != "www.domain" {
		t.Error(out)
	}

	out = stringutils.RemoveEndIgnoreCase("www.domain.中国汉字", ".中国汉字")
	if out != "www.domain" {
		t.Error(out)
	}
}

func TestEndsWithAny(t *testing.T) {

	out := stringutils.EndsWithAny("")
	assert.Assert(t, out == false)

	out = stringutils.EndsWithAny("abcxyz", "")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAny("abcxyz", "xyz")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAny("abcxyz", "xyz", "abc")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAny("abcXYZ", "def", "XYZ")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAny("abcXYZ", "def", "xyz")
	assert.Assert(t, out == false)

	out = stringutils.EndsWithAny("abcx中国汉字yz", "def", "x中国汉字yz")
	assert.Assert(t, out == true)
}

func TestEndsWithAnyIgnoreCase(t *testing.T) {

	out := stringutils.EndsWithAnyIgnoreCase("")
	assert.Assert(t, out == false)

	out = stringutils.EndsWithAnyIgnoreCase("abcxyz", "")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAnyIgnoreCase("abcxyz", "xyz")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAnyIgnoreCase("abcxyz", "xyz", "abc")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAnyIgnoreCase("abcXYZ", "def", "XYZ")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAnyIgnoreCase("abcXYZ", "def", "xyz")
	assert.Assert(t, out == true)

	out = stringutils.EndsWithAnyIgnoreCase("abcx中国汉字yz", "def", "x中国汉字yz")
	assert.Assert(t, out == true)
}

func TestAppendIfMissing(t *testing.T) {

	out := stringutils.AppendIfMissing("abc", "")
	assert.Assert(t, out == "abc")

	out = stringutils.AppendIfMissing("", "xyz")
	assert.Assert(t, out == "xyz")

	out = stringutils.AppendIfMissing("abc", "xyz")
	assert.Assert(t, out == "abcxyz")

	out = stringutils.AppendIfMissing("abcxyz", "xyz")
	assert.Assert(t, out == "abcxyz")

	out = stringutils.AppendIfMissing("abcXYZ", "xyz")
	assert.Assert(t, out == "abcXYZxyz")

	out = stringutils.AppendIfMissing("abc", "xyz", "")
	assert.Assert(t, out == "abc")

	out = stringutils.AppendIfMissing("abc", "xyz", "mno")
	assert.Assert(t, out == "abcxyz")

	out = stringutils.AppendIfMissing("abcxyz", "xyz", "mno")
	assert.Assert(t, out == "abcxyz")

	out = stringutils.AppendIfMissing("abcmno", "xyz", "mno")
	assert.Assert(t, out == "abcmno")

	out = stringutils.AppendIfMissing("abcXYZ", "xyz", "mno")
	assert.Assert(t, out == "abcXYZxyz")

	out = stringutils.AppendIfMissing("abcMNO", "xyz", "mno")
	assert.Assert(t, out == "abcMNOxyz")
}

func TestAppendIfMissingIgnoreCase(t *testing.T) {
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("", "xyz") == "xyz")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abc", "xyz") == "abcxyz")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abcxyz", "xyz") == "abcxyz")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abcXYZ", "xyz") == "abcXYZ")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abc", "xyz", "") == "abc")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abc", "xyz", "mno") == "abcxyz")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abcxyz", "xyz", "mno") == "abcxyz")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abcmno", "xyz", "mno") == "abcmno")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abcXYZ", "xyz", "mno") == "abcXYZ")
	assert.Assert(t, stringutils.AppendIfMissingIgnoreCase("abcMNO", "xyz", "mno") == "abcMNO")
}

func TestPrependIfMissing(t *testing.T) {
	assert.Assert(t, stringutils.PrependIfMissing("", "xyz") == "xyz")
	assert.Assert(t, stringutils.PrependIfMissing("abc", "xyz") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissing("xyzabc", "xyz") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissing("XYZabc", "xyz") == "xyzXYZabc")
	assert.Assert(t, stringutils.PrependIfMissing("abc", "xyz", "") == "abc")
	assert.Assert(t, stringutils.PrependIfMissing("abc", "xyz", "mno") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissing("xyzabc", "xyz", "mno") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissing("mnoabc", "xyz", "mno") == "mnoabc")
	assert.Assert(t, stringutils.PrependIfMissing("XYZabc", "xyz", "mno") == "xyzXYZabc")
	assert.Assert(t, stringutils.PrependIfMissing("MNOabc", "xyz", "mno") == "xyzMNOabc")
}

func TestPrependIfMissingIgnoreCase(t *testing.T) {
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("", "xyz") == "xyz")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("abc", "xyz") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("xyzabc", "xyz") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("XYZabc", "xyz") == "XYZabc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("abc", "xyz", "") == "abc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("abc", "xyz", "mno") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("xyzabc", "xyz", "mno") == "xyzabc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("mnoabc", "xyz", "mno") == "mnoabc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("XYZabc", "xyz", "mno") == "XYZabc")
	assert.Assert(t, stringutils.PrependIfMissingIgnoreCase("MNOabc", "xyz", "mno") == "MNOabc")
}

func TestWrap(t *testing.T) {

	assert.Assert(t, stringutils.Wrap("", "asdfasdf") == "")
	assert.Assert(t, stringutils.Wrap("ab", "x") == "xabx")
	assert.Assert(t, stringutils.Wrap("ab", "\"") == "\"ab\"")
	assert.Assert(t, stringutils.Wrap("\"ab\"", "\"") == "\"\"ab\"\"")
	assert.Assert(t, stringutils.Wrap("ab", "'") == "'ab'")
	assert.Assert(t, stringutils.Wrap("'abcd'", "'") == "''abcd''")
	assert.Assert(t, stringutils.Wrap("\"abcd\"", "'") == "'\"abcd\"'")
	assert.Assert(t, stringutils.Wrap("'abcd'", "\"") == "\"'abcd'\"")
}

func TestWrapIfMissing(t *testing.T) {
	assert.Assert(t, stringutils.WrapIfMissing("", "asdfasdf") == "")
	assert.Assert(t, stringutils.WrapIfMissing("ab", "x") == "xabx")
	assert.Assert(t, stringutils.WrapIfMissing("ab", "\"") == "\"ab\"")
	assert.Assert(t, stringutils.WrapIfMissing("\"ab\"", "\"") == "\"ab\"")
	assert.Assert(t, stringutils.WrapIfMissing("ab", "'") == "'ab'")
	assert.Assert(t, stringutils.WrapIfMissing("'abcd'", "'") == "'abcd'")
	assert.Assert(t, stringutils.WrapIfMissing("\"abcd\"", "'") == "'\"abcd\"'")
	assert.Assert(t, stringutils.WrapIfMissing("'abcd'", "\"") == "\"'abcd'\"")
	assert.Assert(t, stringutils.WrapIfMissing("/", "/") == "/")
	assert.Assert(t, stringutils.WrapIfMissing("a/b/c", "/") == "/a/b/c/")
	assert.Assert(t, stringutils.WrapIfMissing("/a/b/c", "/") == "/a/b/c/")
	assert.Assert(t, stringutils.WrapIfMissing("a/b/c/", "/") == "/a/b/c/")
}

func TestLastIndexOf(t *testing.T) {
	// assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "a", 8) == 7)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "b", 8) == 5)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "ab", 8) == 4)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "b", 9) == 5)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "b", -1) == -1)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "a", 0) == 0)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "b", 0) == -1)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "b", 1) == -1)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "b", 2) == 2)
	assert.Assert(t, stringutils.LastIndexOf("aabaabaa", "ba", 2) == 2)
}

func TestUnwrap(t *testing.T) {

	assert.Assert(t, stringutils.Unwrap("'abc'", "'") == "abc")
	assert.Assert(t, stringutils.Unwrap("\"abc\"", "\"") == "abc")
	assert.Assert(t, stringutils.Unwrap("AABabcBAA", "AA") == "BabcB")
	assert.Assert(t, stringutils.Unwrap("A", "#") == "A")
	assert.Assert(t, stringutils.Unwrap("#A", "#") == "#A")
	assert.Assert(t, stringutils.Unwrap("A#", "#") == "A#")
}

func TestRand(t *testing.T) {

	out := stringutils.Rand(8, stringutils.STR_RAND_KIND_ALL)
	t.Log(out)
}
