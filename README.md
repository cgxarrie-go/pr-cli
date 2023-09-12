# prq
Interaction with pull requests from command line

## Config commands
- prq config : display config
- prq config pat : set PAT in azure config

## List PR commands 
List will scan the current folder and subfolders searching for git repositories.
Then, PRs in the found repositories will be listd

- prq list : Lists all PR in status Active  (default)

### modifiers
-s : specify the status of the PRs to be listed

- prq list -s active: Lists all PR in status Active for azure projects and repos
- prq list -s abandoned: Lists all PR in status Abandoned for azure projects and repos
- prq list -s cancelled: Lists all PR in status Cancelled for azure projects and repos


## Create PR commands 
- prq create : creates a draft PR from current branch to default destination brnach with default title

default destination branch is **master**
deafult title is **PR from spurce-branch-name to destination-branch-name**

### modifiers
-d : specify destination brnach of the PR
-t : set the title of the PR

- prq create -d <branchname> : creates a draft PR from current branch to **branchname** with default title
- prq create -t <pr-title> : crecreates a draft PR from current branch to default destination branch with title **pr-title**
- prq create  -d <branchname> -t <pr-title> : crecreates a draft PR from current branch to **branchname** with title **pr-title**
