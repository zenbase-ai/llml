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
  just py run ruff format
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

publish lang version:
  #!/usr/bin/env bash
  set -euxo pipefail

  if [ "{{lang}}" = "ts" ]; then
    cd ts
    pnpm version {{version}}
    pnpm publish --access public
  elif [ "{{lang}}" = "py" ]; then
    cd py
    CURRENT_VERSION=$(rg '^version = "([^"]+)"' pyproject.toml -r '$1')
    NEW_VERSION=$(echo "$CURRENT_VERSION" | ruby ../bin/bump.rb {{version}})
    uvx --from=toml-cli toml set --toml-path=pyproject.toml project.version "$NEW_VERSION"
    uv build
    uv publish
  elif [ "{{lang}}" = "go" ]; then
    cd go
    # Get latest go tag, extract version, bump it
    LATEST_TAG=$(git tag -l "go/v*" --sort=-version:refname | head -1 || echo "go/v0.0.0")
    CURRENT_VERSION=$(echo "$LATEST_TAG" | rg -o 'v(\d+\.\d+\.\d+)' -r '$1')
    NEW_VERSION=$(echo "$CURRENT_VERSION" | ruby ../bin/bump.rb {{version}})
    NEW_TAG="go/v$NEW_VERSION"
    git tag "$NEW_TAG"
    git push --tags
    echo "Published Go module with tag: $NEW_TAG"
  elif [ "{{lang}}" = "rs" ]; then
    cd rs
    CURRENT_VERSION=$(rg '^version = "([^"]+)"' Cargo.toml -r '$1')
    NEW_VERSION=$(echo "$CURRENT_VERSION" | ruby ../bin/bump.rb {{version}})
    sed -i "s/^version = \".*\"/version = \"$NEW_VERSION\"/" Cargo.toml
    cargo publish
  fi
