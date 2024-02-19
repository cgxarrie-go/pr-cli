# prq
Interaction with pull requests from command line
List and create PRs in Azure or Github repositories


## Installation

- Use the binaries the [latest release](https://github.com/cgxarrie-go/prq/releases/tag/v0.2.6)

Allocate the binay in a folder accessible from command line

    -prq.exe for Windows

    -prq for Mac

-Use go install to install the latest version
```
    go install github.com/cgxarrie-go/prq@latest
```
    


## Usage

### Config commands
- prq config : display config
- prq config azpat : set PAT in Azure config
- prq config ghpat : set PAT in Github config
- prq config remotes -a **remote** : Add a remote to config
- prq config remotes -r **remote** : Remove remote from config

### List PR commands 
List will list all active PRs in the remote of the current folder's local git 
repository (Azure ot Github)

- prq list : Lists all PR in status Active in the repository in the current directory
- prq list --option d : Lists all PR in status Active in all the repositories found in the current directory tree
- prq list --option c : Lists all PR in status Active in all the repositories in config remotes

- prq list --filter frank : Lists all PR in status Active in the repository in the current directory with the word frank in the title, authoror status

### Create PR commands 
- prq create : creates a draft PR from current branch to default destination 
brnach with default title

default destination branch is **master** in Azure and **main** in Github
deafult title is **PR from spurce-branch-name to destination-branch-name**

#### modifiers
-d : specify destination brnach of the PR
-t : set the title of the PR

- prq create -d **branchname** : creates a draft PR from current branch to **branchname** with default title
- prq create -t **pr-title** : crecreates a draft PR from current branch to default destination branch with title **pr-title**
- prq create  -d **branchname** -t **pr-title** : crecreates a draft PR from current branch to **branchname** with title **pr-title**
