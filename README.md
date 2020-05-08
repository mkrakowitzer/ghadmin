# GitHub Admin CLI tool

ghadmin is a GitHub CLI tool for administrating GitHub organisations. I am building this tool because I am finding the GUI is incredibly cumbersome for managing a medium scale organisation. If you have stumbled across this team because you have a requirment to manage your github organisation, Terraform's github provider may be the better option.

I used the GitHub official CLI tool as a boilerplate for this project, Primarily because it looked like a good starting point to get up and running quickly while I learn golang.

This tool is in the very early stages of development; there are plenty of bugs and many missing features which I intend to add. If you spot bugs or have features that you'd like to see in `ghadmin` feel free to contribute.        

## Usage

- `ghadmin repo [list, edit]`
- `ghadmin repo list --all --org boringWorks`
- `ghadmin repo edit --disable-issues --org BoringWorks [reponame]`
- `ghadmin team list --org boringWorks members Platform`
- `ghadmin team create --name foo2 --org boringWorks`
- `ghadmin team delete --name foo2 --org boringWorks`
- `ghadmin team add --name foo1 --org boringWorks --repo test2 --permission pull`
- `ghadmin team list --org boringWorks`
- `ghadmin team list --org boringWorks members foo1`

## Installation

Install a prebuilt binary from the [releases page][]

