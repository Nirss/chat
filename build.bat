go build
set GOARCH=wasm
set GOOS=js
go build -o web/app.wasm
chat.exe