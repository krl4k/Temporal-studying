# Temporal application for studying purposes

## Why Temporal?

* Temporal SDKs empower developers to focus on creating reliable and scalable business logic.
* Eliminates the need for building custom supervisor systems for reliability and fault-tolerance.
* Temporal SDK offers a unified library, abstracting the complexities of how Temporal manages distributed systems.

## Temporal Application 
A Temporal Application is the code you write, comprised of:
* Workflow Definitions
  * A Workflow Definition is the code that defines the constraints of a Workflow Execution.
* Activity Definitions
  * An Activity Definition is the code that defines the constraints of an Activity Task Execution.
* code used to configure Temporal Clients
  * 
* code used to configure and start Workers. 



## Requirements
* Goland
* Temporal environment


## How to run

#### Install and Run Temporal server
```bash
# install temporal
brew install temporal 

# run temporal server
temporal server start-dev

# create namespace
temporal operator namespace create backgroundcheck_namespace
```

#### WEB UI address
```
http://localhost:8233/
```

#### Start workflow using CLI
```bash
# start workflow
temporal workflow start \
 --task-queue backgroundcheck-boilerplate-task-queue-local \
 --type BackgroundCheck \
 --input '"555-55-5555"' \
 --namespace backgroundcheck_namespace \
 --workflow-id backgroundcheck_workflow
```

#### List workflows using CLI
```bash
temporal workflow list \
--namespace backgroundcheck_namespace
```

#### Save workflow history to file
```bash
temporal workflow show \
 --workflow-id backgroundcheck_workflow \
 --namespace backgroundcheck_namespace \
 --output json > backgroundcheck_workflow_event_history.json
```

# Links
* [Source code of the tutorial](https://docs.temporal.io/dev-guide/go/project-setup#boilerplate-project)
