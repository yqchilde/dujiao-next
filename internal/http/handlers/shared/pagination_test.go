package shared

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParsePagination(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		target       string
		wantPage     int
		wantPageSize int
	}{
		{
			name:         "uses query values",
			target:       "/?page=3&page_size=40",
			wantPage:     3,
			wantPageSize: 40,
		},
		{
			name:         "uses defaults when query is absent",
			target:       "/",
			wantPage:     1,
			wantPageSize: 20,
		},
		{
			name:         "normalizes invalid values",
			target:       "/?page=bad&page_size=bad",
			wantPage:     1,
			wantPageSize: 20,
		},
		{
			name:         "bounds values",
			target:       "/?page=0&page_size=300",
			wantPage:     1,
			wantPageSize: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request = httptest.NewRequest("GET", tt.target, nil)

			gotPage, gotPageSize := ParsePagination(c)

			if gotPage != tt.wantPage || gotPageSize != tt.wantPageSize {
				t.Fatalf("ParsePagination() = (%d, %d), want (%d, %d)", gotPage, gotPageSize, tt.wantPage, tt.wantPageSize)
			}
		})
	}
}
