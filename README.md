#Go Jobs
[![Go Report Card](https://goreportcard.com/badge/github.com/djap96/gojobs)](https://goreportcard.com/report/github.com/djap96/gojobs)

##Installation

Get binary
```
go install gj-getbadge
go install gj-background
go install gj-maxsound
```
or quick use
```
go run gj-getbadge
go run gj-background
go run gj-maxsound
```

##Usage


###GJ GetBadge
Get a custom badge from [shields.io api](http://shields.io) in your workspace quickly.  
gj-getbadge support web color alias.
To get more information about the command please run `gj-getbadge`

```
gj-getbadge Stable v2.3.4
```
[![Stable](http://djap96.github.io/gojobs/assets/stable.svg)](#)

```
gj-getbadge Beta v3.1.1 c=tomato
```
[![Beta](http://djap96.github.io/gojobs/assets/beta.svg)](#)


```
gj-getbadge "Client Mode" "Not Supported" c=black
```
[![Client](http://djap96.github.io/gojobs/assets/client.svg)](#)

###GJ Background
Change the background for your Gnome Desktop Enviroment, you need to change the directories in gj-background.go before install!

```
gj-background
gj-background chicks
```
"chicks" exists as alias in gj-background.go

###GJ MaxSound
Change the sound capacity to 150% or whatever percentage, use at your own risk,,,
```
gj-maxsound
gj-maxsound 200
```
