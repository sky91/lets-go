module github.com/sky91/lets-go/ipcola

go 1.21

require (
	github.com/pkg/errors v0.9.1
	github.com/sky91/lets-go/gox v0.0.0
)

replace (
	github.com/sky91/lets-go/gox  => ../gox
)