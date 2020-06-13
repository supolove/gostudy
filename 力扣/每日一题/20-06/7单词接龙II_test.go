package _0_06

import "testing"

/*
https://leetcode-cn.com/problems/word-ladder-ii/

*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	var ans [][]string

	containWord := false
	m := map[string]bool{}
	for _, word := range wordList {
		if word == endWord {
			containWord = true
			m[word] = true
		}
	}
	if !containWord {
		return ans
	}
	return ans
}

func bfs(beginWord, endWord string, wordList map[string]bool, ans [][]string) {
	//var queue []string
	//var mpath []string
	//mpath = append(mpath,beginWord)

}

func Test_findLadders(t *testing.T) {

}
