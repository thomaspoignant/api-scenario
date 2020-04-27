
go test . -coverprofile cover.out.tmp
cat cover.out.tmp | grep -v "_generated.go" > cover.out
tool cover -func cover.out

