package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestAuthMiddlewareValidToken(t *testing.T) {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK) 
    })

    testHandler := AuthMiddleware(handler)

    r, _ := http.NewRequest("GET", "/", nil)
    r.Header.Set("Authorization", AuthToken) 
    w := httptest.NewRecorder()
    testHandler.ServeHTTP(w, r)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status OK for valid token, got %v", w.Code)
    }
}

func TestAuthMiddlewareInvalidToken(t *testing.T) {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    })

    testHandler := AuthMiddleware(handler)

    r, _ := http.NewRequest("GET", "/", nil)
    r.Header.Set("Authorization", "InvalidToken")

    w := httptest.NewRecorder()
    testHandler.ServeHTTP(w, r)

    if w.Code != http.StatusUnauthorized {
        t.Errorf("Expected status Unauthorized for invalid token, got %v", w.Code)
    }
}
