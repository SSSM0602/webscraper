package main
import (
	"testing"
	"reflect"
	"net/url"
)

func TestGetH1FromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h1>Test Title</h1></body></html>"
	actual := getH1FromHTML(inputBody)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetH1FromHTMLNested(t *testing.T) {
	inputBody := `<html><body><header><div><h1>Nested Title</h1></div></header></body></html>`
	actual := getH1FromHTML(inputBody)
	expected := "Nested Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetH1FromHTMLWithSpecialChars(t *testing.T) {
	inputBody := `<html><body><h1>Title with &amp; Symbol</h1></body></html>`
	actual := getH1FromHTML(inputBody)
	expected := "Title with & Symbol" // entities decoded by goquery

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLMultipleInMain(t *testing.T) {
	inputBody := `<html><body>
		<main>
			<p>First in main.</p>
			<p>Second in main.</p>
		</main>
		<p>Outside paragraph.</p>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "First in main." // should pick first <p> inside <main>

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLDeepNested(t *testing.T) {
	inputBody := `<html><body>
		<main>
			<section><article><p>Deep nested main paragraph.</p></article></section>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Deep nested main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="https://blog.boot.dev"><span>Boot.dev</span></a></body></html>`

    baseURL, err := url.Parse(inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://example.com"
	inputBody := `<html><body>
		<a href="/about">About</a>
	</body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://example.com/about"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLMultipleLinks(t *testing.T) {
	inputURL := "https://example.org"
	inputBody := `<html><body>
		<a href="/page">External</a>
		<a href="/contact">Contact</a>
		<a href="/blog/post">Blog Post</a>
	</body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://example.org/page",
		"https://example.org/contact",
		"https://example.org/blog/post",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
func TestGetImagesFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="/logo.png" alt="Logo"></body></html>`

    baseURL, err := url.Parse(inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
