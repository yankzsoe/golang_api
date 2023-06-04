go clean

go build -ldflags="-X main.ApiVersion=1.0.1 -X main.Environment=DEV" -o ./golang_api.exe

golang_api.exe