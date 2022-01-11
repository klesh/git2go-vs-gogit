# Intro

Benchmark:

- [go-git/go-git: A highly extensible Git implementation in pure Go.](https://github.com/go-git/go-git)
- [libgit2/git2go: Git to Go; bindings for libgit2. Like McDonald's but tastier.](https://github.com/libgit2/git2go)


# Step to launch

1. Clone a target repo into your local drive:

```
git clone https://github.com/pingcap/tidb.git  /path/to/your/local/repo
```
2. Update `LOCAL_REPO_PATH` inside `benchmark_test.go`
```
const LOCAL_REPO_PATH = "/path/to/your/local/repo"
```
3. Benchmark for `git2go`
```
go test -bench Git2Go
```
4. Benchmark for `gogit`
```
go test -bench GoGit
```
