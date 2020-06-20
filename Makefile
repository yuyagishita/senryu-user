test:
	@go test -v -covermode=count -coverprofile=users.coverprofile ./users
	@rm users.coverprofile
