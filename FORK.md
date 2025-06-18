# mislav/go-readability

This is a fork of `github.com/go-shiori/go-readability` that adds performance optimizations and some fixes for extracting correct article content.

## Changes

- Merge pull request #1 from mislav/dependabot/github_actions/all-ad97c67762 - 9b631b1
- Merge pull request #2 from mislav/dependabot/go_modules/all-136686d3d5 - 3e37e5f
- Merge pull request #3 from mislav/ci-tweaks - 82828cc
- Merge branch 'counting-optimizations' - ef9026a
- Merge remote-tracking branch 'origin/master' - 95cc3ac
- Merge branch 'generate-test-timestamps' - 9ebfc8b
- Merge branch 'log-outerhtml' - ad84ce8
- Merge branch 'preserve-div-ids' - ee91de6
- Merge branch 'verbose-flag' - f5ecec2
- Merge branch 'test-parser-truncate-diff' - df47d01
- Merge branch 'figure-fix' - a44e548
- Merge pull request #4 from mislav/linter-fixes - 31d0ef0
- Merge pull request #5 from mislav/parse-and-mutate - daa20d1

### My pull requests to upstream

- [#100: Fix extracting images from figures on Medium](https://github.com/go-shiori/go-readability/pull/100)
- [#99: Optimize traversing the DOM when analyzing text content](https://github.com/go-shiori/go-readability/pull/99)
- [#98: Optimize logging HTML nodes by skipping costly compute when Debug is off](https://github.com/go-shiori/go-readability/pull/98)
- [#97: script/generate-test: add publishedTime, modifiedTime to expected metadata](https://github.com/go-shiori/go-readability/pull/97)
- [#96: Test_parser: make failures for text content mismatches more readable](https://github.com/go-shiori/go-readability/pull/96)
- [#95: Preserve IDs & class names of unwrapped DIVs](https://github.com/go-shiori/go-readability/pull/95)

## Benchmark

The benchmark measures the performance of pasing a very large HTML document (`test-pages/wikipedia-2/source.html`).

Before:
~~~
BenchmarkParser-8   	      24	  49127665 ns/op	73733471 B/op	  201173 allocs/op
~~~

After:
~~~
BenchmarkParser-8   	      43	  33544991 ns/op	 7943065 B/op	  106772 allocs/op
~~~
