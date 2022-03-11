package main
////
//import (
//	//"github.com/miiy/gogoweb/api/routes"
//	"github.com/gin-gonic/gin"
//	"github.com/miiy/go-web/examples/web-server/internal/app/auth"
//	"github.com/miiy/go-web/pkg/router"
//
//	"github.com/miiy/go-web/examples/web-server/internal/app/health"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestPingRoute(t *testing.T) {
//	r := &router.Router{
//		Engine: gin.Default(),
//	}
//	// register router
//	m := auth.NewModule(r, nil, nil, nil, nil)
//
//	m.RegisterRouter()
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/health/liveness", nil)
//	engine.ServeHTTP(w, req)
//
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, "Ol", w.Body.String())
//}