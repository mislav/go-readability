# mislav/go-readability

This is a fork of `github.com/go-shiori/go-readability` that adds performance optimizations and some fixes for extracting correct article content.

## Changes

- Merge pull request #1 from mislav/dependabot/github_actions/all-ad97c67762 - mislav/go-readability@9b631b1
- Merge pull request #2 from mislav/dependabot/go_modules/all-136686d3d5 - mislav/go-readability@3e37e5f
- Merge pull request #3 from mislav/ci-tweaks - mislav/go-readability@82828cc
- Merge branch 'counting-optimizations' - mislav/go-readability@ef9026a
- Merge remote-tracking branch 'origin/master' - mislav/go-readability@95cc3ac
- Merge branch 'generate-test-timestamps' - mislav/go-readability@9ebfc8b
- Merge branch 'log-outerhtml' - mislav/go-readability@ad84ce8
- Merge branch 'preserve-div-ids' - mislav/go-readability@ee91de6
- Merge branch 'verbose-flag' - mislav/go-readability@f5ecec2
- Merge branch 'test-parser-truncate-diff' - mislav/go-readability@df47d01
- Merge branch 'figure-fix' - mislav/go-readability@a44e548
- Merge pull request #4 from mislav/linter-fixes - mislav/go-readability@31d0ef0
- Merge pull request #5 from mislav/parse-and-mutate - mislav/go-readability@daa20d1

### My pull requests to upstream

- [#100: Fix extracting images from figures on Medium](https://github.com/go-shiori/go-readability/pull/100)
- [#99: Optimize traversing the DOM when analyzing text content](https://github.com/go-shiori/go-readability/pull/99)
- [#98: Optimize logging HTML nodes by skipping costly compute when Debug is off](https://github.com/go-shiori/go-readability/pull/98)
- [#97: script/generate-test: add publishedTime, modifiedTime to expected metadata](https://github.com/go-shiori/go-readability/pull/97)
- [#96: Test_parser: make failures for text content mismatches more readable](https://github.com/go-shiori/go-readability/pull/96)
- [#95: Preserve IDs & class names of unwrapped DIVs](https://github.com/go-shiori/go-readability/pull/95)

## Benchmark

The benchmark measures the performance of parsing a very large HTML document (`test-pages/wikipedia-2/source.html`):

~~~
before: BenchmarkParser-8   	      24	  50281588 ns/op	73722886 B/op	  201173 allocs/op
after : BenchmarkParser-8   	      39	  28989067 ns/op	 7935952 B/op	  106772 allocs/op
~~~
