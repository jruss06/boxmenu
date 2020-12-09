# boxmenu
Openbox dynamic menu generator created in python3

This is a simple script to generate openbox menus while avoiding external dependencies

## Install
Simply download the script and place in your path

Check menu.xml for an example

## Usage
boxmenu [options]

options:
   
       -w, --write    writes menu.xml automatically 
       
## Config
path: $HOME/.config/boxmenu/config.json

- turn other menu on/off
- set favorite items
- set system items
- set categories (first part is name is menu and the second part is the tag it looks for)
