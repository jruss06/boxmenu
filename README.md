# boxmenu
Openbox dynamic menu generator created in python3.

This is a simple script to generate openbox menus while avoiding external dependencies.

## Go Version
There is now a golang branch that has most of the same features but has better performace and no python dependency (standalone executable)

### python
real	0m0.045s

user	0m0.038s

sys	0m0.007s
### go
real	0m0.003s

user	0m0.004s

sys	0m0.000s

## Install
Simply download the script and place in your path.

Check menu.xml for an example.

Starting from release v0.5.0 the config file is now mandatory.

## Usage
boxmenu [options]

options:
   
       -w, --write    writes menu.xml automatically 
       
## Config
path: $HOME/.config/boxmenu/config.json

- turn other menu on/off
- set favorite items
- set system items
- set categories (first part is the menu name and the second part is the tag from the desktop file)
