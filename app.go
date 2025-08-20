package main

import (
	"context"
	"fmt"
	"qouteswidget/qoutes"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	latestQuote qoutes.Quote
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Set initial quote
	a.latestQuote = qoutes.GetRandomQuote()
	// Start goroutine to update quote every 5 minutes
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				a.latestQuote = qoutes.GetRandomQuote()
				// Emit event to frontend
				runtime.EventsEmit(a.ctx, "newQuote", a.latestQuote)
			case <-ctx.Done():
				return
			}
		}
	}()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	fmt.Println("Greet called with name:", name)
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetRandomQuote returns a new random quote and updates the latest quote
func (a *App) GetRandomQuote() qoutes.Quote {
	a.latestQuote = qoutes.GetRandomQuote()
	return a.latestQuote
}

// GetLatestQuote returns the most recently generated quote
func (a *App) GetLatestQuote() qoutes.Quote {
	return a.latestQuote
}
