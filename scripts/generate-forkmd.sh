#!/bin/bash
set -euo pipefail

run_benchmark() {
  go test -benchmem -run='^$' -bench '^BenchmarkParser$' . | grep BenchmarkParser-
}

sed -i.bak 's/ParseAndMutate/ParseDocument/' bench_test.go
git fetch upstream
git checkout -q upstream/master
bench_before="$(run_benchmark)"
git switch -
mv -v bench_test.go.bak bench_test.go
bench_after="$(run_benchmark)"

{
  cat <<MARKDOWN
# mislav/go-readability

This is a fork of \`github.com/go-shiori/go-readability\` that adds performance optimizations, compatibility with [Readability.js 0.6.0](https://github.com/mozilla/readability/blob/main/CHANGELOG.md#060---2025-03-03), and some fixes for extracting article contents such as images.

## Changes

MARKDOWN

  git log --reverse --format='- %s - mislav/go-readability@%h' --merges upstream/master... | grep -vwF fork.md

  cat <<MARKDOWN

### My pull requests to upstream

MARKDOWN

  gh pr list -R go-shiori/go-readability --author=@me --state=all --json number,title,url \
    --template '{{range .}}- [{{printf "#%v: %s" .number .title}}]({{.url}}){{"\n"}}{{end}}'

  cat <<MARKDOWN

## Benchmark

The benchmark measures the performance of parsing a very large HTML document (\`test-pages/wikipedia-2/source.html\`):

~~~
before: ${bench_before}
after : ${bench_after}
~~~
MARKDOWN
} > FORK.md
