# prq
Interaction with pull requests from command line

## Config commands
- prq config : display config
- prq config org : set company name in azure config
- prq config pat : set PAT in azure config
- prq config project -a <name> : Adds a project with name <name> in azure config
- prq config project -d <name> : Removes a project with name <name> in azure config
- prq config repo -p <project-name> -a <name> : Adds a repo with name <name> to the project with name <project-name> in azure config
- prq config repo -p <project-name> -d <name> : Removes a repo with name <name> from the project with name <project-name> in azure config

## List PR commands 
- prq list : Lists all PR in status Active for azure projects and repos
- prq list --status active: Lists all PR in status Active for azure projects and repos
- prq list --status abandoned: Lists all PR in status Abandoned for azure projects and repos
- prq list --status cancelled: Lists all PR in status Cancelled for azure projects and repos
