package main

import "testing"

const LOCAL_REPO_PATH = "/home/klesh/Projects/merico/tidb"

func BenchmarkGit2Go(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetCommitsByGit2Go(LOCAL_REPO_PATH)
	}
}

func BenchmarkGoGit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetCommitsByGoGit(LOCAL_REPO_PATH)
	}
}
