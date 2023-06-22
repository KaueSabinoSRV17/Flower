# Flower 🌻

A CLI tool made to help developers using Git Flow and Conventional Commits.

## How to Install

Currently, it can be Installed via:

### Brew

```bash
brew install KaueSabinoSRV17/homebrew-flower/flower
```

### Linux Binary

```bash
wget https://github.com/KaueSabinoSRV17/Flower/releases/download/1.0.6/Flower_1.0.6_linux_amd64.tar.gz
tar -xzf Flower_1.0.6_linux_amd64.tar.gz
mv ./flow /usr/local/bin/flow
rm Flower_1.0.6_linux_amd64.tar.gz
```

## How to Use

Currently, it has only the `commit` command:

```bash
flow commit
```

It will ask you what files you are going to add to git, what is going to be the Conventional Commit prefix
and the Commit Message. After that, you can run `git log` to validate the commit.

## Contribuitions

### Project Structure

The Structure can and probaly should be changed, it would be very welcome some advices on how to organize
the files and folders properly, since i come from `Typescript` Development.

#### .github directory

Where we store `yaml` files for the ci/cd runs in Github Actions. Currently, i use `goreleaser` to automate
the publishing to `brew`.

#### cmd directory

Where we store `Go` files generated by `Cobra`. They add the functionality of a CLI to the business logic inside
`use_cases`.

#### infra

Where we store `terraform` files that describe and mainteen any infrastructure that we need to publish the CLI.
Currently there is only a `S3 Bucket` in `AWS`, that we are trying to use as a repository for `apt`, making it
easier to install on Linux.

#### use_cases directory

Where we store business logic to implement the `Use Cases` of the CLI (Currently only Conventional Commits).