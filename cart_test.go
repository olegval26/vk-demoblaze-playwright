package demoblaze

import (
	"testing"
	"time"

	"github.com/mxschmitt/playwright-go"
)

const (
	baseURL     = "https://www.demoblaze.com"
	productName = "Samsung galaxy s6"
)

func TestSamsungGalaxyS6IsOnlyItemInCart(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("start Playwright: %v", err)
	}
	defer func() {
		if err := pw.Stop(); err != nil {
			t.Errorf("stop Playwright: %v", err)
		}
	}()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		t.Fatalf("launch Chromium: %v", err)
	}
	defer func() {
		if err := browser.Close(); err != nil {
			t.Errorf("close browser: %v", err)
		}
	}()

	context, err := browser.NewContext()
	if err != nil {
		t.Fatalf("create isolated browser context: %v", err)
	}
	defer func() {
		if err := context.Close(); err != nil {
			t.Errorf("close browser context: %v", err)
		}
	}()

	page, err := context.NewPage()
	if err != nil {
		t.Fatalf("open browser page: %v", err)
	}
	if _, err := page.Goto(baseURL); err != nil {
		t.Fatalf("open DemoBlaze: %v", err)
	}

	assertions := playwright.NewPlaywrightAssertions(15_000)
	productLink := page.GetByRole(playwright.AriaRole("link"), playwright.PageGetByRoleOptions{
		Name:  productName,
		Exact: playwright.Bool(true),
	})
	if err := productLink.Click(); err != nil {
		t.Fatalf("open %q product page: %v", productName, err)
	}

	productHeading := page.GetByRole(playwright.AriaRole("heading"), playwright.PageGetByRoleOptions{
		Name:  productName,
		Exact: playwright.Bool(true),
	})
	if err := assertions.Locator(productHeading).ToBeVisible(); err != nil {
		t.Fatalf("product page did not show %q: %v", productName, err)
	}

	dialogMessages := make(chan string, 1)
	dialogErrors := make(chan error, 1)
	page.OnDialog(func(dialog playwright.Dialog) {
		dialogMessages <- dialog.Message()
		dialogErrors <- dialog.Accept()
	})
	addToCart := page.GetByRole(playwright.AriaRole("link"), playwright.PageGetByRoleOptions{
		Name:  "Add to cart",
		Exact: playwright.Bool(true),
	})
	if err := addToCart.Click(); err != nil {
		t.Fatalf("click Add to cart: %v", err)
	}
	select {
	case got := <-dialogMessages:
		if got != "Product added" {
			t.Fatalf("unexpected add-to-cart dialog: got %q, want %q", got, "Product added")
		}
	case <-time.After(10 * time.Second):
		t.Fatal("the add-to-cart confirmation did not appear")
	}
	select {
	case err := <-dialogErrors:
		if err != nil {
			t.Fatalf("accept add-to-cart dialog: %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("the add-to-cart confirmation was not accepted")
	}

	cartLink := page.GetByRole(playwright.AriaRole("link"), playwright.PageGetByRoleOptions{
		Name:  "Cart",
		Exact: playwright.Bool(true),
	})
	if err := cartLink.Click(); err != nil {
		t.Fatalf("open cart: %v", err)
	}

	cartRows := page.Locator("#tbodyid tr")
	if err := assertions.Locator(cartRows).ToHaveCount(1); err != nil {
		t.Fatalf("cart should contain exactly one item: %v", err)
	}
	if err := assertions.Locator(cartRows.First()).ToContainText(productName); err != nil {
		t.Fatalf("the cart item should be %q: %v", productName, err)
	}
}
