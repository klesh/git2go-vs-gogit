package main

import (
	"fmt"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func GetCommitsByGoGit(repoPath string) {
	fmt.Println("GetCommitsByGoGit")
	repo, err := gogit.PlainOpen(repoPath)
	if err != nil {
		panic(err)
	}

	iter, err := repo.CommitObjects()
	if err != nil {
		panic(err)
	}

	iter.ForEach(func(commit *object.Commit) error {
		insertions := 0
		deletions := 0
		parents := ""

		fileStats, err := commit.Stats()
		if err != nil {
			panic(err)
		}

		for _, fileStat := range fileStats {
			insertions += fileStat.Addition
			deletions += fileStat.Deletion
		}

		parentIter := commit.Parents()

		parentIter.ForEach(func(parentCommit *object.Commit) error {
			if parents != "" {
				parents += ", "
			}
			parents += parentCommit.Hash.String()
			return nil
		})

		fmt.Printf("%v %v   added: %v  deleted: %v parents: %v\n", commit.Hash.String(), commit.Author.Email, insertions, deletions, parents)
		return nil
	})
}
