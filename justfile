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

[working-directory("rs")]
rs *args:
  cargo {{args}}

fmt:
  just ts biome check --write --unsafe
  just py run ruff check --fix
  just go fmt go/...
  just rs fmt

lint lang="*":
  #!/usr/bin/env bash
  set -euxo pipefail # e = exit on error, u = treat unset variables as errors, x = print commands, o = print options, pipefail = exit on error in a pipeline
  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "ts" ]; then
    just ts biome check
  fi
  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "py" ]; then
    just py run ruff check
  fi
  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "go" ]; then
    just go vet go/...
    if [ "$(just go fmt go/... | wc -l)" -gt 0 ]; then
      echo "The following files are not formatted:"
      just go fmt go/...
      exit 1
    fi
  fi
  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "rs" ]; then
    just rs clippy -- -D warnings && just rs fmt -- --check
  fi

test lang="*":
  #!/usr/bin/env bash
  set -euxo pipefail # e = exit on error, u = treat unset variables as errors, x = print commands, o = print options, pipefail = exit on error in a pipeline

  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "ts" ]; then
    just ts test
  fi
  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "py" ]; then
    just py run pytest
  fi
  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "go" ]; then
    just go test ./...
  fi
  if [ "{{lang}}" = "*" ] || [ "{{lang}}" = "rs" ]; then
    just rs test
  fi
