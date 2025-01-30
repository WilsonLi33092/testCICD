module github.com/WilsonLi33092/testCICD/actions/main

go 1.23.5

replace github.com/WilsonLi33092/testCICD/add => ./add

require (
	github.com/WilsonLi33092/testCICD/add v0.0.0-00010101000000-000000000000
	github.com/WilsonLi33092/testCICD/calculations v0.0.0-00010101000000-000000000000
	github.com/WilsonLi33092/testCICD/subtract v0.0.0-00010101000000-000000000000
)

replace github.com/WilsonLi33092/testCICD/subtract => ./subtract

replace github.com/WilsonLi33092/testCICD/calculations => ./calculations
