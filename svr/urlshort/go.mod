module example.com/urlshort

go 1.18

replace example.com/handler => ../handler

require (
	example.com/handler v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)
