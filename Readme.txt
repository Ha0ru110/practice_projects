GOOS=linux GOARCH=amd64 go build  equation_solver.go

GOOS=darwin GOARCH=amd64 go build  -o=equation_solver_darwin  equation_solver.go

GOOS=windows GOARCH=amd64 go build  -o=equation_solver_windows_version.exe  equation_solver.go