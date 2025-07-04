package re2go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsByline(t *testing.T) {
	assert.True(t, IsByline(`<h4 class="article-byline">`))
	assert.True(t, IsByline(`<meta name="author" content="Die Tagespost" />`))
	assert.True(t, IsByline(`<span class="dateline">CAIRO — Agustine Myer</span>`))
	assert.True(t, IsByline(`<h3 class="writtenbynames">ΣΥΝΤΑΚΤΙΚΗ ΟΜΑΔΑ</h3>`))
	assert.True(t, IsByline(`<span class="meta-prep-author">Publiziert am</span>`))

	assert.False(t, IsByline(`<h4 class="article-line">`))
	assert.False(t, IsByline(`<meta name="autor" content="Die Tagespost" />`))
	assert.False(t, IsByline(`<span class="date">CAIRO — Agustine Myer</span>`))
	assert.False(t, IsByline(`<h3 class="bynames">ΣΥΝΤΑΚΤΙΚΗ ΟΜΑΔΑ</h3>`))
	assert.False(t, IsByline(`<span class="meta-autor">Publiziert am</span>`))
}

func Test_IsPositiveClass(t *testing.T) {
	assert.True(t, IsPositiveClass(`<section class="article">Random content here</section>`))
	assert.True(t, IsPositiveClass(`<body class="body">Some body content</body>`))
	assert.True(t, IsPositiveClass(`<div class="content">This is content inside a div</div>`))
	assert.True(t, IsPositiveClass(`<article class="entry">An entry in the article</article>`))
	assert.True(t, IsPositiveClass(`<span class="hentry">Highlight this entry</span>`))
	assert.True(t, IsPositiveClass(`<header class="h-entry">Header for h-entry</header>`))
	assert.True(t, IsPositiveClass(`<main class="main">Main section content</main>`))
	assert.True(t, IsPositiveClass(`<nav class="page">Page navigation content</nav>`))
	assert.True(t, IsPositiveClass(`<ul class="pagination">Pagination list</ul>`))
	assert.True(t, IsPositiveClass(`<aside class="post">This is a post</aside>`))
	assert.True(t, IsPositiveClass(`<p class="text">Some paragraph text</p>`))
	assert.True(t, IsPositiveClass(`<article class="blog">Blog article content</article>`))
	assert.True(t, IsPositiveClass(`<section class="story">A story section</section>`))

	assert.False(t, IsPositiveClass(`<header class="header">Header here</header>`))
	assert.False(t, IsPositiveClass(`<footer class="footer">Footer section</footer>`))
	assert.False(t, IsPositiveClass(`<div class="container">This inside a container</div>`))
	assert.False(t, IsPositiveClass(`<section class="sidebar">This is a sidebar</section>`))
	assert.False(t, IsPositiveClass(`<nav class="navigation">Navigation links</nav>`))
	assert.False(t, IsPositiveClass(`<p class="description">Paragraph description</p>`))
	assert.False(t, IsPositiveClass(`<div class="news">Latest news</div>`))
	assert.False(t, IsPositiveClass(`<aside class="widget">A widget section</aside>`))
	assert.False(t, IsPositiveClass(`<div class="layout">Side layout</div>`))
	assert.False(t, IsPositiveClass(`<section class="gallery">Gallery of images</section>`))
}

func Test_IsNegativeClass(t *testing.T) {
	assert.True(t, IsNegativeClass(`<div class="ad-banner">Advertisement banner content</div>`))
	assert.True(t, IsNegativeClass(`<section class="hidden">Hidden section</section>`))
	assert.True(t, IsNegativeClass(`<div class="-ad-">Ad content here</div>`))
	assert.True(t, IsNegativeClass(`hid`))
	assert.True(t, IsNegativeClass(`hid class`))
	assert.True(t, IsNegativeClass(`class hid`))
	assert.True(t, IsNegativeClass(`class hid good`))
	assert.True(t, IsNegativeClass(`<section class="hid">Again, hid match</section>`))
	assert.True(t, IsNegativeClass(`<div class="banner">Banner content</div>`))
	assert.True(t, IsNegativeClass(`<aside class="combx">Comb box content</aside>`))
	assert.True(t, IsNegativeClass(`<section class="comment">User comments here</section>`))
	assert.True(t, IsNegativeClass(`<div class="com-">Com- prefix example</div>`))
	assert.True(t, IsNegativeClass(`<section class="contact">Contact information</section>`))
	assert.True(t, IsNegativeClass(`<footer class="foot">Footer section</footer>`))
	assert.True(t, IsNegativeClass(`<section class="footer">Footer content</section>`))
	assert.True(t, IsNegativeClass(`<footer class="footnote">Footnote details</footer>`))
	assert.True(t, IsNegativeClass(`<div class="gdpr">GDPR compliance text</div>`))
	assert.True(t, IsNegativeClass(`<header class="masthead">Masthead content</header>`))
	assert.True(t, IsNegativeClass(`<div class="media">Media gallery</div>`))
	assert.True(t, IsNegativeClass(`<section class="meta">Meta information</section>`))
	assert.True(t, IsNegativeClass(`<div class="outbrain">Outbrain recommendations</div>`))
	assert.True(t, IsNegativeClass(`<section class="promo">Promotional content</section>`))
	assert.True(t, IsNegativeClass(`<div class="related">Related articles</div>`))
	assert.True(t, IsNegativeClass(`<section class="scroll">Scrolling section</section>`))
	assert.True(t, IsNegativeClass(`<div class="share">Sharing tools</div>`))
	assert.True(t, IsNegativeClass(`<aside class="shoutbox">Shoutbox chat</aside>`))
	assert.True(t, IsNegativeClass(`<nav class="sidebar">Sidebar links</nav>`))
	assert.True(t, IsNegativeClass(`<section class="skyscraper">Skyscraper ad</section>`))
	assert.True(t, IsNegativeClass(`<div class="sponsor">Sponsored content</div>`))
	assert.True(t, IsNegativeClass(`<section class="shopping">Shopping cart</section>`))
	assert.True(t, IsNegativeClass(`<div class="tags">Tag list</div>`))
	assert.False(t, IsNegativeClass(`<div class="tool">Tools and settings</div>`))
	assert.True(t, IsNegativeClass(`<aside class="widget">Widget features</aside>`))

	assert.False(t, IsNegativeClass(`<header class="navbar">Navigation bar</header>`))
	assert.False(t, IsNegativeClass(`<section class="overview">Overview section content</section>`))
	assert.False(t, IsNegativeClass(`<div class="gallery">Gallery of images</div>`))
	assert.False(t, IsNegativeClass(`<aside class="support">Support section</aside>`))
	assert.False(t, IsNegativeClass(`<div class="catalog">Product catalog</div>`))
	assert.False(t, IsNegativeClass(`<nav class="user-menu">User menu links</nav>`))
	assert.False(t, IsNegativeClass(`<article class="news-feed">Latest news</article>`))
	assert.False(t, IsNegativeClass(`<section class="details">Detailed section content</section>`))
	assert.False(t, IsNegativeClass(`<div class="profile">User profile content</div>`))
}

func Test_IsUnlikelyCandidates(t *testing.T) {
	assert.True(t, IsUnlikelyCandidates(`<div class="ad-banner">Ad banner content</div>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="-ad-">Ad-related content</section>`))
	assert.True(t, IsUnlikelyCandidates(`<article class="ai2html">AI to HTML conversion content</article>`))
	assert.True(t, IsUnlikelyCandidates(`<nav class="banner">Banner navigation</nav>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="breadcrumbs">Breadcrumbs navigation</section>`))
	assert.True(t, IsUnlikelyCandidates(`<aside class="combx">Comb box content</aside>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="comment">Comment section</section>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="community">Community forum</div>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="cover-wrap">Cover wrap for article</div>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="disqus">Disqus comment section</section>`))
	assert.True(t, IsUnlikelyCandidates(`<aside class="extra">Extra content</aside>`))
	assert.True(t, IsUnlikelyCandidates(`<footer class="footer">Footer section</footer>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="gdpr">GDPR compliance</div>`))
	assert.True(t, IsUnlikelyCandidates(`<header class="header">Header content</header>`))
	assert.True(t, IsUnlikelyCandidates(`<aside class="legends">Legends and explanations</aside>`))
	assert.True(t, IsUnlikelyCandidates(`<nav class="menu">Menu navigation</nav>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="related">Related articles</section>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="remark">Remark section</div>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="replies">Replies to comments</section>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="rss">RSS feed</div>`))
	assert.True(t, IsUnlikelyCandidates(`<aside class="shoutbox">Shoutbox chat</aside>`))
	assert.True(t, IsUnlikelyCandidates(`<nav class="sidebar">Sidebar content</nav>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="skyscraper">Skyscraper ad</section>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="social">Social media links</div>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="sponsor">Sponsored content</section>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="supplemental">Supplemental information</div>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="ad-break">Ad break content</div>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="agegate">Age verification</div>`))
	assert.True(t, IsUnlikelyCandidates(`<nav class="pagination">Pagination links</nav>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="pager">Pager navigation</div>`))
	assert.True(t, IsUnlikelyCandidates(`<section class="popup">Popup content</section>`))
	assert.True(t, IsUnlikelyCandidates(`<div class="yom-remote">Yom remote content</div>`))

	assert.False(t, IsUnlikelyCandidates(`<div class="container">Main container</div>`))
	assert.False(t, IsUnlikelyCandidates(`<section class="overview">Overview section</section>`))
	assert.False(t, IsUnlikelyCandidates(`<article class="newsfeed">Newsfeed content</article>`))
	assert.False(t, IsUnlikelyCandidates(`<section class="gallery">Image gallery</section>`))
	assert.False(t, IsUnlikelyCandidates(`<div class="catalog">Product catalog</div>`))
	assert.False(t, IsUnlikelyCandidates(`<section class="summary">Summary content</section>`))
}

func Test_MaybeItsACandidate(t *testing.T) {
	assert.True(t, MaybeItsACandidate(`<section class="and">Logical and condition</section>`))
	assert.True(t, MaybeItsACandidate(`<article class="article">Article content</article>`))
	assert.True(t, MaybeItsACandidate(`<body class="body">Body of the document</body>`))
	assert.True(t, MaybeItsACandidate(`<div class="column">Column layout</div>`))
	assert.True(t, MaybeItsACandidate(`<section class="content">Main content here</section>`))
	assert.True(t, MaybeItsACandidate(`<main class="main">Main section</main>`))
	assert.True(t, MaybeItsACandidate(`<div class="shadow">Shadow effect</div>`))

	assert.False(t, MaybeItsACandidate(`<header class="header">Header section</header>`))
	assert.False(t, MaybeItsACandidate(`<div class="navbar">Navigation bar</div>`))
	assert.False(t, MaybeItsACandidate(`<section class="footer">Footer section</section>`))
	assert.False(t, MaybeItsACandidate(`<nav class="menu">Menu navigation</nav>`))
	assert.False(t, MaybeItsACandidate(`<section class="gallery">Photo gallery</section>`))
	assert.False(t, MaybeItsACandidate(`<p class="text">Paragraph text</p>`))
}

func Test_NormalizeSpaces(t *testing.T) {
	assert.Equal(t, "some sentence", NormalizeSpaces("some   sentence"))
	assert.Equal(t, "with tabs", NormalizeSpaces("with \t \ttabs"))
	assert.Equal(t, " single space is ok ", NormalizeSpaces(" single space is ok "))
	assert.Equal(t, " multi space removed ", NormalizeSpaces("   multi   space   removed   "))
}
