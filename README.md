# Gocat

Yep. I just rewrote the GNU Cat program in Go.

## Usage

`List contents of file`: `gocat [filename]`

## Installation

Install using the following commands:

* Clone the project using `git clone https://github.com/wantyapps/gocat.git/`

* Change directory to the project using `cd gocat/`

* Install the binary/command using `make install`

* You are done! You can now remove the project's directory using `cd .. && rm -rf gocat/`

* You can now execute gocat by using the command `gocat`

Recommendation: Add the line `alias cat="gocat"` to your .bashrc/.zshrc to run
gocat instead of cat when using the command `cat`.

## Logging and Debugging

Please note that the program makes a log at the configured directory.

The log file's name is also configurable. See `Configuring` section below.

If you encountered an error or want to create an issue, please attach your log file.

_Do not run the program again if you want to file a bug report, because the program
will overwrite the existing log file._

## Configuring

First, create a config file in your home directory (`~`)
named `.gocat.yml` (notice the config language is Yaml).

Then enter the lines below to the file:

```yaml
LogfileDir:     ".gocat/"
LogfileName:    "gocat.log"
```

Notice that you have to manually create the log file's directory (`LogfileDir`).

If you want to place the log file inside the home (`~`) directory without any subdirectories, set the
`LogfileDir` value to `.`.

(The `LogfileDir` value is placed in the user's home directory aka `~`)

_This project was created only to sharpen my Go skills. You can still use it, just know that._
