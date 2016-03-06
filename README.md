#Go Tasks
[![Go Report Card](https://goreportcard.com/badge/github.com/djap96/gotasks)](https://goreportcard.com/report/github.com/djap96/gotasks)

##Installation

Get binary
```
go install gt-getbadge
go install gt-background
go install gt-maxsound
```
or quick use
```
go run gt-getbadge
go run gt-background
go run gt-maxsound
```

##Usage


###GT GetBadge
Get a custom badge from [shields.io api](http://shields.io) in your workspace quickly.  
gt-getbadge support web color alias.
To get more information about the command please run `gt-getbadge`

```
gt-getbadge Stable v2.3.4
```
[![Stable](http://djap96.github.io/gotasks/assets/stable.svg)](#)

```
gt-getbadge Beta v3.1.1 c=tomato
```
[![Beta](http://djap96.github.io/gotasks/assets/beta.svg)](#)


```
gt-getbadge "Client Mode" "Not Supported" c=black
```
[![Client](http://djap96.github.io/gotasks/assets/client.svg)](#)

###GT Background
Change the background for your Gnome Desktop Enviroment, you need to change the directories in gt-background.go before install!

```
gt-background
gt-background chicks
```
"chicks" exists as alias in gt-background.go

###GT MaxSound
Change the sound capacity to 150% or whatever percentage, use at your own risk,,,
```
gt-maxsound
gt-maxsound 200
```
