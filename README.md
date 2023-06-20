# Flower

A CLI tool made to help developers using Git Flow and Conventional Commits.

## How to Install

Currently, it can be Installed via

### Brew

````bash
brew install KaueSabinoSRV17/homebrew-flower/flower
``````

### Linux Binary

```bash
wget https://github.com/KaueSabinoSRV17/Flower/releases/download/1.0.5/Flower_1.0.5_linux_amd64.tar.gz
tar -xzf Flower_1.0.5_linux_amd64.tar.gz
mv ./flow /usr/local/bin/flow
rm Flower_1.0.5_linux_amd64.tar.gz
```

## How to Use

Currently, it has only the `commit` command:

````bash
flow commit
``````

It will ask you what files you are going to add to git, what is going to be the Conventional Commit prefix
and the Commit Message. After that, you can run `git log` to validate the commit
