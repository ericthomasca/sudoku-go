# Sudoku Solver

[![Go Report Card](https://goreportcard.com/badge/github.com/ericthomasca/sudoku-go)](https://goreportcard.com/report/github.com/ericthomasca/sudoku-go)

A simple Sudoku solver written in Go.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Sudoku Solver is a Go package that provides functionality to solve Sudoku puzzles. It offers a simple HTTP API to send unsolved Sudoku puzzles and receive the solved puzzles in JSON format.

## Features

- Solve Sudoku puzzles of varying difficulty levels.
- Accept unsolved puzzles via HTTP POST requests.
- Return solved puzzles in JSON format.
- Lightweight and easy-to-use.

## Installation

To install the Sudoku solver package, you can use `go get`:

```bash
go get github.com/ericthomasca/sudoku-go
