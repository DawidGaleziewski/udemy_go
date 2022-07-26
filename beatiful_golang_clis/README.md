code along to video on youtube:
https://www.youtube.com/watch?v=SSRIn5DAmyw&t=722s&ab_channel=CloudNativeSkunkworks

# Intro
when creating a cli we want to have a entry point.
This will conect to something we can call "palettes" which are like painters palets.
Each of those will be a set of other libreries that can be used in some domain

i.e structure
```bash
- Toolbox
-- palette1 : [info, disk usage]
-- palette2 : [net, ping]

```

## Cobra
we will be using "Cobra" for creation of our CLI. Cobra is a library AND a cli

cli will allow us to setup a project:
```bash
$ ../goworkspace/bin/cobra-cli.exe init toolbox
```

after that we can add new commands
```bash
../goworkspace/bin/cobra-cli.exe add net
```

## nesting commands
we can nest commands by:
1. Creating a sperata package for a command
2. Add a subcommand to that commands cobra.Command struct instance in init function

```go
func init() {
	NetCmd.AddCommand(pingCmd) // we can nest this way sub commands in init
}
```