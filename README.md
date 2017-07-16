# aws-architect

Simple tool useful to define smart architectures on Amazon Web Service.


### Objectives

Objectives of project are :

* Define an Amazon Web Services accelerator
* Support complex CloudFormation operations
* Integrate multiple services
* Provide an AWS client agnostic application


### Pre-requisites

N.A.


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
