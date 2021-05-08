<h1 align="center">AtomicGo | cursor</h1>

<p align="center">

<a href="https://github.com/atomicgo/cursor/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/cursor?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/cursor" target="_blank">
<img src="https://img.shields.io/github/workflow/status/atomicgo/cursor/Go?label=tests&style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/cursor" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/cursor?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/cursor">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-16-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://github.com/atomicgo/cursor/issues">
<img src="https://img.shields.io/github/issues/atomicgo/cursor.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

</p>

---

<p align="center">
<strong><a href="#install">Get The Module</a></strong>
|
<strong><a href="https://pkg.go.dev/github.com/atomicgo/cursor" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

## Description

Package cursor contains methods to move the cursor inside a terminal.

## Install

```console
# Execute this command inside your project
go get -u github.com/atomicgo/cursor
```

```go
// Add this to your imports
import "github.com/atomicgo/cursor"
```

## Usage

#### func  ClearLine

```go
func ClearLine()
```

#### func  ClearLines

```go
func ClearLines(n int)
```

#### func  ClearScreen

```go
func ClearScreen()
```

#### func  CloseAlternativeScreen

```go
func CloseAlternativeScreen()
```

#### func  Down

```go
func Down(n int)
```

Down moves the cursor n cells down. If the cursor is already at the edge of the screen, this has no effect.

#### func  Hide

```go
func Hide()
```

#### func  Left

```go
func Left(n int)
```

Left moves the cursor n cells left. If the cursor is already at the edge of the screen, this has no effect.

#### func  Move

```go
func Move(row int, column int)
```

Move moves the cursor to a specific row and column.

#### func  NextLine

```go
func NextLine(n int)
```

#### func  OpenAlternativeScreen

```go
func OpenAlternativeScreen()
```

#### func  PrevLine

```go
func PrevLine(n int)
```

#### func  RestorePosition

```go
func RestorePosition()
```

#### func  Right

```go
func Right(n int)
```

Right moves the cursor n cells right. If the cursor is already at the edge of the screen, this has no effect.

#### func  SavePosition

```go
func SavePosition()
```

#### func  Show

```go
func Show()
```

#### func  Up

```go
func Up(n int)
```

Up moves the cursor n cells up. If the cursor is already at the edge of the screen, this has no effect.

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
