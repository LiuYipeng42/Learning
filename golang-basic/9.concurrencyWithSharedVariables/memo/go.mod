module memo

go 1.19

replace memo1 => ./memo1

replace memo2 => ./memo2

require (
	memo1 v0.0.0-00010101000000-000000000000
	memo2 v0.0.0-00010101000000-000000000000
)
