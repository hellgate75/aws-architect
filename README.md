# AWS Architect tool
<p align="center"><img src="https://travis-ci.org/hellgate75/aws-architect.svg?branch=master" alt="trevis-ci" width="98" height="20" />&nbsp;<a href="https://travis-ci.org/hellgate75/aws-architect">Check last build on Travis-CI</a></p><br/>

Simple tool useful to define smart architectures on Amazon Web Service.


### Why prefer this tool to AWS client?

Answer is very simple, we prefer this tool to AWS client because :

* We do not want to maintain aws client, hiss versioning or BOTO libraries, and updating this project you will update to latest aws sdk
* We want an an accelerator do deploy faster AWS infrastructure.
* We do not want do strive our mind with aws client complex api
* We want ready build-in procedures to access services
* We want a tool easy to integrate with new features
* We are providing a plugin interface to develop different scenarios
* We want integrated AWS Architect features with AWS Developer operations
* We want to access AWS with simplified guided procedures and batch YAML scripts to deploy multiple complex architectures cross-region, maintaining versioning of cloud artifacts and stacks


### Objectives

Project objectives are :

* Provide an AWS client agnostic application
* Define an Amazon Web Services accelerator
* Support complex CloudFormation operations
* Integrate multiple AWS services
* Flatten distance between DevOps and CloudFormation
* Create new operative system working in both DevOps, SysOps, and Architectures
* YAML batch language for architecture and multi-region stacks
* Plugin interface to fast develop new features or simply pull from repository new features, without any integration effort


### Pre-requisites

You can use an AWS credentials file to specify your credentials. This is a special, INI-formatted file stored under your HOME directory, and is a good way to manage credentials for your development environment. The file should be placed at `~/.aws/credentials`, where `~` represents your HOME directory.

Using an AWS credentials file offers a few benefits:

    Your projects' credentials are stored outside of your projects, so there is no chance of accidentally committing them into version control.
    You can define and name multiple sets of credentials in one place.
    You can easily reuse the same credentials between projects.

The format of the AWS credentials file should look something like the following:
```bash

[default]
aws_access_key_id = YOUR_AWS_ACCESS_KEY_ID
aws_secret_access_key = YOUR_AWS_SECRET_ACCESS_KEY

[project1]
aws_access_key_id = ANOTHER_AWS_ACCESS_KEY_ID
aws_secret_access_key = ANOTHER_AWS_SECRET_ACCESS_KEY

```

Each section (e.g., [default], [project1]), represents a separate credential profile. Profiles can be referenced from a SDK configuration flag, or when you are instantiating a client, using the profile option:

```bash

 -profile <profile_name>

```


### Install client

To install client use following command :

```bash
    go get -u github.com/golang/dep/cmd/dep
    go get -u github.com/hellgate75/aws-architect
    dep ensure -update
```


### Run command

To execute client use following command :

```bash
    aws-architect
```

With this command you will have the list of available sub-commands. In case you want inspect sub-command sample usage you can execute following command :

 ```bash
     aws-architect help <sub-command> 
 ```


### Some samples

Here a list of common used configuration file samples:

* [YAML CORs configuration file](/samples/cors.yml)


### Issues

We track and improve code accordingly to plan and issues you create on project issues tracker at :
[Issues Tracker](https://github.com/hellgate75/aws-architect/issues)


Thanks you for advertising us on library bugs or improvements.


### License

[LGPL-v3](/LICENSE)

Extra clauses :

* This software is not re-distributable for commercial use prior authors written authorization 
* Extension of this library is prohibited without prior written authorization of authors
* Authors decline any responsibility on un-appropriate use of the library
* This software is covered by copyright rules, if you want know more write authors at hellgate75@gmail.com
* Authors will maintain this library on discontinuos periods of time, now blueprints or milestone are provided at the moment
* Users are not authorized for commercial use of library without a specific license from authors, and business packages for company LDAP/SAM integration
* This library is not delivered in Gopkg, because we identify in it a potential AWS integrated tool-set, in a future time
