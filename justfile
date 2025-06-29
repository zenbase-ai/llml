set dotenv-load := true
set dotenv-path := ".env"
set ignore-comments := true

claude *args:
  pnpx @anthropic-ai/claude-code {{args}}

gemini *args:
  pnpx @google/gemini-cli {{args}}

[working-directory("ts")]
ts *args:
  bun {{args}}

[working-directory("py")]
py *args:
  uv {{args}}

[working-directory("go")]
go *args:
  go {{args}}

[working-directory(".")]
test:
  ts test
  py test
  go test ./...
