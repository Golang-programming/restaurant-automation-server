package middleware

import (
	"net/http"
)

func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// tenantID := r.Header.Get("X-Tenant-ID") // Or derive it from URL, subdomain, etc.

		// err := database.SetDB(tenantID)
		// if err != nil {
		// 	http.Error(w, "Tenant not found", http.StatusNotFound)
		// 	return
		// }

		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}
