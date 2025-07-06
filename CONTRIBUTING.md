# Contributing to LLML

It could be cool if each foundation lab released their own set of formatters optimized for their LLMs.

We love your input! We want to make contributing to LLML as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## We Develop with GitHub

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

## We Use [GitHub Flow](https://guides.github.com/introduction/flow/index.html)

Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Any contributions you make will be under the MIT Software License

In short, when you submit code changes, your submissions are understood to be under the same [MIT License](LICENSE) that covers the project. Feel free to contact the maintainers if that's a concern.

## Report bugs using GitHub's [issues](https://github.com/zenbase-ai/llml/issues)

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/zenbase-ai/llml/issues/new); it's that easy!

## Write bug reports with detail, background, and sample code

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

## Development Process

### Prerequisites

1. Install [mise](https://mise.jdx.dev/) for tool management
2. Run `mise install` to get the required development tools

### Setting Up Development Environment

```bash
# Clone the repository
git clone https://github.com/zenbase-ai/llml.git
cd llml

# Install development tools
mise install

# Set up language-specific dependencies
cd py && uv install && cd ..
cd ts && bun install && cd ..
cd rs && cargo build && cd ..
cd go && go mod download && cd ..
```

### Running Tests

We maintain consistent behavior across all four language implementations. Always test your changes:

```bash
# Run all tests
just test

# Run language-specific tests
just test py
just test ts
just test rs
just test go
```

### Code Style

Each language follows its ecosystem's conventions:

- **Python**: Black formatting, ruff linting
- **TypeScript**: Biome formatting and linting
- **Rust**: rustfmt and clippy
- **Go**: gofmt and standard linting

Run linters before submitting:

```bash
# Run all linters
just lint

# Run language-specific linters
just lint py
just lint ts
just lint rs
just lint go
```

## Multi-Language Consistency

LLML is implemented in four languages. When making changes:

1. **Consistency is key**: All implementations must produce identical output for the same input
2. **Test across languages**: Add the same test case to all four test suites
3. **Update all READMEs**: Keep documentation synchronized across implementations

### Adding a New Feature

1. Start with one language implementation
2. Write comprehensive tests
3. Port to the other three languages
4. Ensure all tests pass across all languages
5. Update documentation in all relevant places

## Testing Guidelines

1. **Test Coverage**: Aim for 90%+ coverage
2. **Edge Cases**: Test empty values, special characters, deeply nested structures
3. **Cross-Language**: Ensure your test exists in all four language test suites
4. **Integration**: Test real-world scenarios, not just unit tests

Example test structure:
```python
def test_my_feature():
    input_data = {"key": "value"}
    expected = "<key>value</key>"
    assert llml(input_data) == expected
```

## Documentation

- Update the main README.md for user-facing changes
- Update language-specific READMEs for implementation details
- Add examples to the relevant example directories
- Keep the FAQ updated with common questions

## Pull Request Process

1. **Title**: Use a clear, descriptive title
2. **Description**: Explain what changes you made and why
3. **Testing**: Describe how you tested your changes
4. **Screenshots**: If applicable, add screenshots to help explain your changes
5. **Breaking Changes**: Clearly mark any breaking changes

### PR Checklist

- [ ] Tests pass (`just test`)
- [ ] Linting passes (`just lint`)
- [ ] Documentation updated
- [ ] Consistent across all language implementations
- [ ] No breaking changes (or clearly documented)

## Community

- Be respectful and inclusive
- Welcome newcomers and help them get started
- Focus on what is best for the community
- Show empathy towards other community members

## Questions?

Feel free to open an issue with the "question" label or reach out to the maintainers.

## Recognition

Contributors will be recognized in our README. We appreciate every contribution, no matter how small!

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
