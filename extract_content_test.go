package main
import "testing"

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
