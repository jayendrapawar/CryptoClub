# Go Project Setup

This guide will walk you through setting up a Go project from scratch. It includes instructions for setting up your environment, initializing your project, and managing dependencies.

## Prerequisites

- Go installed (version 1.16 or higher)
- Git installed

## Steps

### 1. Install Go

Download and install Go from the [official Go website](https://golang.org/dl/).

### 2. Set Up Go Environment

Ensure your Go environment variables are set correctly. Add the following lines to your shell configuration file (`.bashrc`, `.zshrc`, etc.):

### 3. Set Up Local Environment


API_URL=https://api.livecoinwatch.com/coins/list // https://livecoinwatch.github.io/lcw-api-docs/?go#
API_KEY=3d1d04fe-ae2a-4404-92dc-b4cf4cb1c61b
MONGO_URI=mongodb+srv://<username>:<password>@renegado.fr5ai7g.mongodb.net/?retryWrites=true&w=majority&appName=Renegado
DB_NAME=crypto
COLLECTION_NAME=prices

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
