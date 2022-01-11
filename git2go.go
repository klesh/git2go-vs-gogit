package main

import (
	"fmt"

	git2go "github.com/libgit2/git2go/v28"
)

func GetCommitsByGit2Go(repoPath string) {
	fmt.Println("GetCommitsByGit2Go")
	repo, err := git2go.OpenRepository(repoPath)
	if err != nil {
		panic(err)
	}
	odb, err := repo.Odb()
	if err != nil {
		panic(err)
	}
	err = odb.ForEach(func(id *git2go.Oid) error {
		if id == nil {
			return nil
		}
		commit, err := repo.LookupCommit(id)
		if err != nil {
			fmt.Printf("failed commit: %v\n", id)
			return nil
		}
		insertions := 0
		deletions := 0
		parents := ""
		for i := uint(0); i < commit.ParentCount(); i++ {
			parent := commit.Parent(i)
			parentTree, err := parent.Tree()
			if err != nil {
				panic(err)
			}
			tree, err := commit.Tree()
			if err != nil {
				panic(err)
			}
			diff, err := repo.DiffTreeToTree(parentTree, tree, &git2go.DiffOptions{
				Flags: git2go.DiffIgnoreSubmodules | git2go.DiffIgnoreFilemode,
			})
			if err != nil {
				panic(err)
			}
			stats, err := diff.Stats()
			if err != nil {
				panic(err)
			}
			insertions += stats.Insertions()
			deletions += stats.Deletions()

			if parents != "" {
				parents += ", "
			}
			parents += fmt.Sprintf("%v", parent.Id())
		}
		fmt.Printf("%v %v   added: %v  deleted: %v parents: %v\n", commit.Id(), commit.Author().Email, insertions, deletions, parents)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
