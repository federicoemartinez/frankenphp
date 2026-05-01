package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here.
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/dunglas/frankenphp/caddy"
	_ "github.com/dunglas/mercure/caddy"
	_ "github.com/dunglas/vulcain/caddy"

	// Windows SSO via SSPI (registers http.handlers.sspi_auth Caddyfile directive).
	// Windows-only at runtime; on other platforms the package compiles as a stub
	// that errors at provision time so the build itself stays cross-platform.
	_ "github.com/federicoemartinez/caddy-sspi"
)

func main() {
	caddycmd.Main()
}
