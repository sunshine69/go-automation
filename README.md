An example of using go for automation. Similar to ansible but in a go way

each project has this structure

`inventory` similar to ansible inventory, use the lib https://github.com/relex/aini to parse into hosts/group to use
Currently aini does not support anisble generator inventory style. Creating using ini is really cumber some, I might implement it in the future


`plays` - list of all go `playbook` file - Like deploy, build etc .
each play is a a directory for the playname and the playname is the go sub package name. Each play is one executable with extra vars. Can run alone or trigger by the main prog later.

`tasks` list all go tasks .go files used by plays (sharable)
Can be reuse by playbooks using go import. Not executable by itself.

`mods` - List all go mods files, that is built and copied to remote to exec using the utils ssh tool

`main.go` compile into one executable, run feed it with a host or a playbook exec file name it will run that. An option to list hosts and inventory. Inventory can be external file/dir or use embed to embed into the binary and extract on the fly to allow user can edit it later on.

