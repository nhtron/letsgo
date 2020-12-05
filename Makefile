watch:
	reflex -r '\.(go|gohtml)' -s -- sh -c "go run ./cmd/web/*"
watch-log:
	reflex -r '\.(go|gohtml)' -s -- sh -c "go run ./cmd/web/* >>/tmp/info.log 2>>/tmp/error.log"