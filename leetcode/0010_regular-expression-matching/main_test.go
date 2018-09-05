package main

import "testing"

func Test_isMatch4(t *testing.T) {
	t.Logf("%t", isMatch4("aa", "a"))                    //false
	t.Logf("%t", isMatch4("aa", "a*"))                   //true
	t.Logf("%t", isMatch4("ab", ".*"))                   //true
	t.Logf("%t", isMatch4("aab", "c*a*b"))               //true
	t.Logf("%t", isMatch4("mississippi", "mis*is*p*."))  //false
	t.Logf("%t", isMatch4("abcd", "d*"))                 //false
	t.Logf("%t", isMatch4("a", ""))                      //false
	t.Logf("%t", isMatch4("mississippi", "mis*is*ip*.")) //true
	t.Logf("%t", isMatch4("ab", ".*c"))                  //false
	t.Logf("%t", isMatch4("aaa", "a*a"))                 //true
	t.Logf("%t", isMatch4("aaa", "ab*ac*a"))             //true
	t.Logf("%t", isMatch4("aaa", "ab*a*c*a"))            //true
	t.Logf("%t", isMatch4("aa", "aa"))                   //true
	t.Logf("%t", isMatch4("a", ".*..a*"))                //false

	t.Logf("%t", isMatch4("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s")) //true
}
