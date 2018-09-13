package main

import "testing"

func Test_findSubstring(t *testing.T) {
	t.Logf("result=%v\n", findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	t.Logf("result=%v\n", findSubstring("wordgoodstudentgoodword", []string{"word", "student"}))
	t.Logf("result=%v\n", findSubstring("a", []string{}))
	t.Logf("result=%v\n", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	t.Logf("result=%v\n", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}
