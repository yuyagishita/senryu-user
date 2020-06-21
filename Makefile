test:
	@go get github.com/modocache/gover
	@go test -v -covermode=count -coverprofile=profile.coverprofile
	@go test -v -covermode=count -coverprofile=users.coverprofile ./users
	# @go test -v -covermode=count -coverprofile=mongodb.coverprofile ./db/mongodb
	@gover
	mv gover.coverprofile cover.profile
	@rm *.coverprofile
	@go mod tidy
