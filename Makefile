INSTANCE = user

test:
	@go get github.com/modocache/gover
	@go test -v -covermode=count -coverprofile=profile.coverprofile
	@go test -v -covermode=count -coverprofile=users.coverprofile ./users
	@gover
	mv gover.coverprofile cover.profile
	@rm *.coverprofile
	@go mod tidy
