package config

import (
	"net/http"
	"os"
	"strconv"
)

type CookieInterface interface {
	Name() string
	Domain() string
	Path() string
	Secure() bool
	HttpOnly() bool
	SameSite() http.SameSite
	Expires() int
}

type Cookie struct {
	name     string
	domain   string
	path     string
	secure   bool
	httpOnly bool
	sameSite http.SameSite
	expires  int
}

func initCookieConfig() *Cookie {
	name := os.Getenv("COOKIE_NAME")
	if name == "" {
		name = "cooooookie"
	}

	domain := os.Getenv("COOKIE_DOMAIN")
	if domain == "" {
		domain = "localhost"
	}

	path := os.Getenv("COOKIE_PATH")
	if path == "" {
		path = "/"
	}

	secure, err := strconv.ParseBool(os.Getenv("COOKIE_SECURE"))
	if err != nil {
		secure = false
	}

	httpOnly, err := strconv.ParseBool(os.Getenv("COOKIE_HTTPONLY"))
	if err != nil {
		httpOnly = true
	}

	sameSite, _ := strconv.Atoi(os.Getenv("COOKIE_SAMESITE"))
	if sameSite == 0 {
		sameSite = 2
	}

	expires, _ := strconv.Atoi(os.Getenv("COOKIE_EXPIRES"))
	if expires == 0 {
		expires = 192
	}

	return &Cookie{
		name:     name,
		domain:   domain,
		path:     path,
		secure:   secure,
		httpOnly: httpOnly,
		sameSite: http.SameSite(sameSite),
		expires:  expires,
	}
}

func (cookie *Cookie) Name() string {
	return cookie.name
}

func (cookie *Cookie) Domain() string {
	return cookie.domain
}

func (cookie *Cookie) Path() string {
	return cookie.path
}

func (cookie *Cookie) Secure() bool {
	return cookie.secure
}

func (cookie *Cookie) HttpOnly() bool {
	return cookie.httpOnly
}

func (cookie *Cookie) SameSite() http.SameSite {
	return cookie.sameSite
}

func (cookie *Cookie) Expires() int {
	return cookie.expires
}
