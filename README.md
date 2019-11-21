# HellSpawner
The toolset for creating games for OpenDiablo2

## Building
In the root folder, run `go get -d` to pull down all dependencies.

To run the project, run `go run .` from the root folder.

## Build Dependencies
Reboot after doing this to ensure environment variables are configured properly.
### Linux
TBD
### OSX
TBD
### Windows
GCC is required and can be installed with [Chocolatey](https://chocolatey.org/) by executing the following command:
```shell script
choco install mingw
```

## General Concepts
The toolset will be used to develop games for OpenDiablo 2 (including the Diablo2 project, which will also be
built in this).

Projects will simply consist of a collectin of MPQs in a folder. On startup, the toolset will have a tree of
files in a project navigator window that will allow you to select files inside of a specific MPQ. Optionally you can
view all MPQs in a 'combined' mode which will list the entire structure of all MPQs combined.

When double-clicking on a file, a corresponding editor will pop up for that file, allowing you to view/modify
that data (on supported formats). You will also be able to create new MPQs as well.

In addition to the formats already in Diablo2, several new formats will be created to allow the creation
of screens, quests, gameplay elements, NPC interactions, etc. This can be used to create completely new
games, and will also be used to port Diablo2/LOD to the engine.

## New File Formats
The following is a list of file formats that have been created for this engine.

### Script (.ods)
These files are OpenDiablo2 scripts that can be interpreted by the engine. The syntax is currently not defined,
but will be some form of basic interpreter syntax, like the following:
```basic
FUNC OnMenuExit
    SETSCENE /Path/To/NewScene
ENDFUNC

FUNC OnItemPickup Source:ITEM
    GIVEITEM Source, 1
    PLAYSOUND /Path/To/Sound.wav NOREPEAT
ENDFUNC
```

### Scene Format
This file will define scenes, including buttons, menus, actions, etc.

File specification TBD.
