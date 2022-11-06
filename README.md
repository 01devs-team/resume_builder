# Resume builder

A simple web application to build your resume with VCS ability

# How to run

Project has Makefile which runs docker-compose in each folder. to run frontend or backend run below commands:

- frontend: `make front`
- backend: `make back`

# Contribute

All contribute to project are welcomed but there are few rules that is good to keep in mind

## Commit convention

since frontend and backend are on the same repo, commit messages must show that which of them are affected so commits must have below structure:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

The commit contains the following structural elements:

- **fix**: a commit of the type fix patches a bug in codebase.
- **feat**: a commit of the type feat introduces a new feature to the codebase.
- **BREAKING CHANGE**: a commit that has a footer `BREAKING CHANGE:`, or appends a `!` after the type/scope, introduces a breaking change. A BREAKING CHANGE can be part of commits of any type.
- types other than fix: and feat: are allowed, for example [@commitlint/config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional) (based on the the Angular convention) recommends `build:`, `chore:`, `ci:`, `docs:`, `style:`, `refactor:`, `perf:`, `test:`, and others.
  footers other than `BREAKING CHANGE: <description>` may be provided and follow a convention similar to [git trailer format](https://git-scm.com/docs/git-interpret-trailers).
