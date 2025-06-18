package readability

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	fp "path/filepath"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/go-shiori/dom"
	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/net/html"
)

var (
	fakeHostURL, _ = url.ParseRequestURI("http://fakehost/test/page.html")
)

type ExpectedMetadata struct {
	Title         string `json:"title,omitempty"`
	Byline        string `json:"byline,omitempty"`
	Excerpt       string `json:"excerpt,omitempty"`
	Language      string `json:"language,omitempty"`
	SiteName      string `json:"siteName,omitempty"`
	Readerable    bool   `json:"readerable"`
	PublishedTime string `json:"publishedTime,omitempty"`
	ModifiedTime  string `json:"modifiedTime,omitempty"`
}

func Test_parser(t *testing.T) {
	testDir := "test-pages"
	testItems, err := os.ReadDir(testDir)
	if err != nil {
		t.Errorf("\nfailed to read test directory")
	}

	for _, item := range testItems {
		if !item.IsDir() {
			continue
		}

		itemName := item.Name()
		t.Run(itemName, func(t1 *testing.T) {
			// Prepare path
			sourcePath := fp.Join(testDir, itemName, "source.html")
			expectedPath := fp.Join(testDir, itemName, "expected.html")
			expectedMetaPath := fp.Join(testDir, itemName, "expected-metadata.json")

			// Extract source file
			article, isReaderable, extractedDoc, err := extractSourceFile(sourcePath)
			if err != nil {
				t1.Error(err)
			}

			// Decode expected file
			expectedDoc, err := decodeExpectedFile(expectedPath)
			if err != nil {
				t1.Error(err)
			}

			// Decode expected metadata
			metadata, err := decodeExpectedMetadata(expectedMetaPath)
			if err != nil {
				t1.Error(err)
			}

			// Compare extraction result
			err = compareArticleContent(extractedDoc, expectedDoc)
			if err != nil {
				t1.Error(err)
			}

			// Check metadata
			if metadata.Byline != article.Byline {
				t1.Errorf("byline, want %q got %q\n", metadata.Byline, article.Byline)
			}

			if metadata.Excerpt != article.Excerpt {
				t1.Errorf("excerpt, want %q got %q\n", metadata.Excerpt, article.Excerpt)
			}

			if metadata.SiteName != article.SiteName {
				t1.Errorf("sitename, want %q got %q\n", metadata.SiteName, article.SiteName)
			}

			if metadata.Title != article.Title {
				t1.Errorf("title, want %q got %q\n", metadata.Title, article.Title)
			}

			if metadata.Readerable != isReaderable {
				t1.Errorf("readerable, want %v got %v\n", metadata.Readerable, isReaderable)
			}

			if metadata.Language != article.Language {
				t1.Errorf("language, want %q got %q\n", metadata.Language, article.Language)
			}

			if !timesAreEqual(metadata.PublishedTime, article.PublishedTime) {
				t1.Errorf("date published, want %q got %q\n", metadata.PublishedTime, article.PublishedTime)
			}

			if !timesAreEqual(metadata.ModifiedTime, article.ModifiedTime) {
				t1.Errorf("date modified, want %q got %q\n", metadata.ModifiedTime, article.ModifiedTime)
			}

		})
	}
}

func Test_countCharsAndCommas(t *testing.T) {
	createElement := func(tagName string, children ...*html.Node) *html.Node {
		node := &html.Node{
			Type: html.ElementNode,
			Data: tagName,
		}
		for _, c := range children {
			node.AppendChild(c)
		}
		return node
	}
	textNode := func(text string) *html.Node {
		return &html.Node{
			Type: html.TextNode,
			Data: text,
		}
	}

	tests := []struct {
		name   string
		node   *html.Node
		chars  int
		commas int
	}{
		{
			name: "traverse node descendants",
			node: createElement("div",
				textNode("hello again "),
				createElement("b", textNode("world,")),
				textNode(" peace"),
			),
			chars:  24,
			commas: 1,
		},
		{
			name:   "commas in different languages",
			node:   textNode("now,its،a mixed﹐commas︐from︑various⹁place⸴and⸲country，"),
			chars:  54,
			commas: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChars, gotCommas := countCharsAndCommas(tt.node)
			if gotChars != tt.chars {
				t.Errorf("countCharsAndCommas() chars = %v, want %v", gotChars, tt.chars)
			}
			if gotCommas != tt.commas {
				t.Errorf("countCharsAndCommas() commas = %v, want %v", gotCommas, tt.commas)
			}
		})
	}
}

func extractSourceFile(path string) (Article, bool, *html.Node, error) {
	// Open source file
	f, err := os.Open(path)
	if err != nil {
		return Article{}, false, nil, fmt.Errorf("failed to open source: %v", err)
	}
	defer f.Close()

	// Decode to HTML
	originalDoc, err := html.Parse(f)
	if err != nil {
		return Article{}, false, nil, fmt.Errorf("failed to decode source: %v", err)
	}

	parser := NewParser()
	isReaderable := parser.CheckDocument(originalDoc)

	// Extract readable article
	article, err := parser.ParseAndMutate(originalDoc, fakeHostURL)
	if err != nil {
		return Article{}, false, nil, fmt.Errorf("failed to extract source: %v", err)
	}

	// At this point, we have article.Node which we could potentially return as extractedDoc to
	// compare with the DOM in "expected.html", but there could potentially be differences due to
	// how some HTML nodes might get dropped during "expected.html" parsing. The safest way to
	// compensate for that is to re-parse the HTML representation of the readable DOM.
	extractedDoc, err := dom.Parse(strings.NewReader(article.Content))
	if err != nil {
		return Article{}, false, nil, fmt.Errorf("failed to parse extract to HTML: %v", err)
	}

	return article, isReaderable, extractedDoc, nil
}

func decodeExpectedFile(path string) (*html.Node, error) {
	// Open expected file
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open expected: %v", err)
	}
	defer f.Close()

	// Parse file into HTML document
	doc, err := dom.Parse(f)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expected to HTML: %v", err)
	}

	return doc, nil
}

func decodeExpectedMetadata(path string) (ExpectedMetadata, error) {
	var zero ExpectedMetadata

	// Open expected file
	f, err := os.Open(path)
	if err != nil {
		return zero, fmt.Errorf("failed to open metadata: %v", err)
	}
	defer f.Close()

	// Parse file into map
	var result ExpectedMetadata
	err = json.NewDecoder(f).Decode(&result)
	return result, err
}

func compareArticleContent(result, expected *html.Node) error {
	// Make sure number of nodes is same
	resultNodesCount := len(dom.Children(result))
	expectedNodesCount := len(dom.Children(expected))
	if resultNodesCount != expectedNodesCount {
		return fmt.Errorf("number of nodes is different, want %d got %d",
			expectedNodesCount, resultNodesCount)
	}

	resultNode := result
	expectedNode := expected
	for resultNode != nil && expectedNode != nil {
		// Get node excerpt
		resultExcerpt := getNodeExcerpt(resultNode)
		expectedExcerpt := getNodeExcerpt(expectedNode)

		// Compare tag name
		resultTagName := dom.TagName(resultNode)
		expectedTagName := dom.TagName(expectedNode)
		if resultTagName != expectedTagName {
			return fmt.Errorf("tag name is different\n"+
				"want    : %s (%s)\n"+
				"got     : %s (%s)",
				expectedTagName, expectedExcerpt,
				resultTagName, resultExcerpt)
		}

		// Compare attributes
		resultAttrCount := len(resultNode.Attr)
		expectedAttrCount := len(expectedNode.Attr)
		if resultAttrCount != expectedAttrCount {
			return fmt.Errorf("number of attributes is different\n"+
				"want    : %d (%s)\n"+
				"got     : %d (%s)",
				expectedAttrCount, expectedExcerpt,
				resultAttrCount, resultExcerpt)
		}

		for _, resultAttr := range resultNode.Attr {
			expectedAttrVal := dom.GetAttribute(expectedNode, resultAttr.Key)
			switch resultAttr.Key {
			case "href", "src":
				resultAttr.Val = strings.TrimSuffix(resultAttr.Val, "/")
				expectedAttrVal = strings.TrimSuffix(expectedAttrVal, "/")
			}

			if resultAttr.Val != expectedAttrVal {
				return fmt.Errorf("attribute %s is different\n"+
					"want    : %s (%s)\n"+
					"got     : %s (%s)",
					resultAttr.Key, expectedAttrVal, expectedExcerpt,
					resultAttr.Val, resultExcerpt)
			}
		}

		// Compare text content
		resultText := strings.TrimSpace(dom.TextContent(resultNode))
		expectedText := strings.TrimSpace(dom.TextContent(expectedNode))

		resultText = strings.Join(strings.Fields(resultText), " ")
		expectedText = strings.Join(strings.Fields(expectedText), " ")

		comparator := diffmatchpatch.New()
		diffs := comparator.DiffMain(resultText, expectedText, false)
		wordDiff := truncatedPrettyDiff(diffs)

		if len(diffs) > 1 {
			return fmt.Errorf("text content is different:\nnode: %s\ntext: %s", resultExcerpt, wordDiff)
		}

		// Move to next node
		ps := Parser{}
		resultNode = ps.getNextNode(resultNode, false)
		expectedNode = ps.getNextNode(expectedNode, false)
	}

	return nil
}

func truncatedPrettyDiff(diffs []diffmatchpatch.Diff) string {
	var buf bytes.Buffer
	for i, d := range diffs {
		switch d.Type {
		case diffmatchpatch.DiffInsert:
			buf.WriteString("\x1B[32m")
			buf.WriteString("{+")
			writeTruncatedText(&buf, d.Text, 40, truncateMiddle)
			buf.WriteString("+}")
			buf.WriteString("\x1B[m")
		case diffmatchpatch.DiffDelete:
			buf.WriteString("\x1B[31m")
			buf.WriteString("[-")
			writeTruncatedText(&buf, d.Text, 40, truncateMiddle)
			buf.WriteString("-]")
			buf.WriteString("\x1B[m")
		case diffmatchpatch.DiffEqual:
			tt := truncateMiddle
			if i == 0 {
				tt = truncateLeft
			} else if i == len(diffs)-1 {
				tt = truncateRight
			}
			writeTruncatedText(&buf, d.Text, 40, tt)
		default:
			panic("diff type not implemented: " + d.Type.String())
		}
	}
	return buf.String()
}

type truncateType int

const (
	truncateLeft truncateType = iota
	truncateRight
	truncateMiddle
)

func writeTruncatedText(buf *bytes.Buffer, text string, limit int, where truncateType) {
	n := utf8.RuneCountInString(text)
	if n <= limit {
		buf.WriteString(text)
		return
	}
	textRunes := []rune(text)
	switch where {
	case truncateLeft:
		buf.WriteString("<...>")
		buf.WriteString(string(textRunes[n-limit:]))
	case truncateRight:
		buf.WriteString(string(textRunes[:limit]))
		buf.WriteString("<...>")
	case truncateMiddle:
		buf.WriteString(string(textRunes[:limit/2]))
		buf.WriteString("<...>")
		buf.WriteString(string(textRunes[n-(limit/2):]))
	}
}

func getNodeExcerpt(node *html.Node) string {
	outer := dom.OuterHTML(node)
	outer = strings.Join(strings.Fields(outer), " ")
	if len(outer) < 120 {
		return outer
	}
	return outer[:120]
}

func timesAreEqual(metadataTimeString string, parsedTime *time.Time) bool {
	if metadataTimeString == "" && parsedTime == nil {
		return true
	}

	if metadataTimeString == "" || parsedTime == nil {
		return false
	}

	ps := Parser{}
	metadataTime := ps.getParsedDate(metadataTimeString)
	return metadataTime.Equal(*parsedTime)
}
