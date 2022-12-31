# Codacy GoLang Tools Engine

[![CircleCI](https://circleci.com/gh/codacy/codacy-engine-golang-seed.svg?style=svg)](https://circleci.com/gh/codacy/codacy-engine-golang-seed)

Framework to help integration with external analysis GoLang tools at Codacy.

## Usage

Add to your go tool dependencies:

```bash
go get github.com/codacy/codacy-engine-golang-seed
```


### Tool Implementation


In order to use Golang engine seed, you must create a `struct` that implements `ToolImplementation` interface:

```
// ToolImplementation interface to implement the tool
type ToolImplementation interface {
	Run(tool Tool, sourceDir string) ([]Issue, error)
}
```

Then, on the main function you just need to instantiate it and pass to engine's `StartTool` method.

As an example:

```
// GoTool tool implementation
type GoTool struct {
}

// Run runs the tool implementation
func (i GoTool) Run(tool codacy.Tool, sourceDir string) ([]codacy.Issue, error) {
    // implement tool
    // ...
	return result, nil
}

func main() {
	implementation := GoTool{}

	codacy.StartTool(implementation)
}

```

## Docs

### How to integrate an external analysis tool on Codacy

By creating a docker and writing code to handle the tool invocation and output,
you can integrate the tool of your choice on Codacy!

> To know more about dockers, and how to write a docker file please refer to [https://docs.docker.com/reference/builder/](https://docs.docker.com/reference/builder/)

In this tutorial, we explain how you can integrate an analysis tool of your choice in Codacy.
You can check the code of an already implemented tool and, if you wish, fork it to start your own.
You are free to modify and use it for your own tools.

#### Requirements

* Docker definition with the tool you want to integrate
* Define the documentation for the patterns provided by the tool

## Test

Follow the instructions at [codacy-plugins-test](https://github.com/codacy/codacy-plugins-test/blob/master/README.md#test-definition).

## Submit the docker

**Running the docker**
```bash
docker run -t \
--net=none \
--privileged=false \
--cap-drop=ALL \
--user=docker \
--rm=true \
-v <PATH-TO-FOLDER-WITH-FILES-TO-CHECK>:/src:ro \
<YOUR-DOCKER-NAME>:<YOUR-DOCKER-VERSION>
```

**Docker restrictions**
* Docker image size should not exceed 500MB
* Docker should contain a non-root user named docker with UID/GID 2004
* All the source code of the docker must be public
* The docker base must officially be supported on DockerHub
* Your docker must be provided in a repository through a public git host (ex: GitHub, Bitbucket, ...)

**Docker submission**
* To submit the docker you should send an email to support@codacy.com with the link to the git repository with your docker definition.
* The docker will then be subjected to a review by our team and we will then contact you with more details.

If you have any question or suggestion regarding this guide please contact us at support@codacy.com.

## What is Codacy

[Codacy](https://www.codacy.com/) is an Automated Code Review Tool that monitors your technical debt, helps you improve your code quality, teaches best practices to your developers, and helps you save time in Code Reviews.

### Among Codacy’s features

* Identify new Static Analysis issues
* Commit and Pull Request Analysis with GitHub, BitBucket/Stash, GitLab (and also direct git repositories)
* Auto-comments on Commits and Pull Requests
* Integrations with Slack, HipChat, Jira, YouTrack
* Track issues in Code Style, Security, Error Proneness, Performance, Unused Code and other categories

Codacy also helps keep track of Code Coverage, Code Duplication, and Code Complexity.

Codacy supports PHP, Python, Ruby, Java, JavaScript, and Scala, among others.

### Free for Open Source

Codacy is free for Open Source projects.
test
