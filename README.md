# git-boil
A simple boiler plate for setting up nested git remotes

## Installation

```
$ go get https://github.com/Kjoedicker/git-boil
```

## Configuration

`git-boil` relies on `conf.yaml` which is stored in `$HOME/.config/git-boil/conf.yaml`.

This file contains the logistics of which repos you commonly work with, as set per your preference.

```
remotes: 
  - repo:
      name: <repo-name>
      url: <url>
  - repo:
      name: <repo-name>
      url: <url>

```

## Usage
`git-boil` generates remote urls and assigns them to the `origin` remote handle. 

`git-boil` simply asks for a project name, which it uses to infer the remote paths based on the logistics provided in `conf.yaml`.

For example

```
$ ./git-boil example-project
$ git remote show origin
* remote origin
  Fetch URL: https://github.com/owner/example-project.git
  Push  URL: https://github.com/owner/example-project.git
  Push  URL: https://<alternative-url_1>/owner/example-project.git
```