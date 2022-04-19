module example.com/urlshort

go 1.18

replace example.com/handler => ../handler

require (
	example.com/handler v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)
