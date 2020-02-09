package beego2ts

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"time"
)

func AA() {
	const a = "/Users/sage/work/api"
	const b = "/Users/sage/package/go/src"
	r, err := git.PlainOpen(a)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	//status, err := w.Status()
	//CheckIfError(err)
	//if len(status) == 0 {
	//	return
	//}

	_, err = w.Add(".")
	CheckIfError(err)

	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Sage",
			Email: "zhang@saget.cn",
			When:  time.Now(),
		},
	})
	CheckIfError(err)
	fmt.Println("commit", commit)
	// Prints the current HEAD to verify that all worked well.
	//	//Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)
	//
	fmt.Println("obj", obj)

	err = r.Push(&git.PushOptions{})
	CheckIfError(err)
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
