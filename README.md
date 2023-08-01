# Field Meta Editor

A web application, it is designed for the previewing and editing of the primitive field information of the imported data, generating detailed field configurations for the exploitation within [@kanaries/graphic-walker](https://github.com/Kanaries/graphic-walker).

## Usage

### Prerequisites

#### 1. Golang

_Golang is the programming language used to develop the backend of this monorepo._

Install [Golang](https://golang.org/) (version >= 1.16.0).

On Mac OS, you can use [Homebrew](https://brew.sh/) to install Golang:

```bash
brew install go
```

You can also install the Golang IDE [GoLand](https://www.jetbrains.com/go/) (version >= 2020.3.3).

#### 2. Node.js

_Node.js is a JavaScript runtime built on Chrome's V8 JavaScript engine. It is used to run the main workflow of this monorepo._

Install [Node.js](https://nodejs.org/en/) (version >= 14.0.0).

On Mac OS, you can use [Homebrew](https://brew.sh/) to install Node.js:

```bash
brew install node
```

On Windows, you can visit the [official website](https://nodejs.org/en/download/) to download the installer.

#### 3. Yarn

_Yarn is a package manager for JavaScript projects. It is used to manage dependencies of the frontend and execute scripts in this monorepo._

Install [Yarn](https://yarnpkg.com/) (version >= 1.22.0).

Using npm:

```bash
npm install -g yarn
```

### Install dependencies

```bash
# execute the following command in the root directory of this monorepo, it will install dependencies for all packages
yarn install
```

### Start the backend

```bash
# start the server
go run .
```

### Start the web application

```bash
# serve with hot reload
yarn workspace web-app dev
```

### Configure the backend

Edit the file `packages/web-app/.env` to configure the backend.

```dotenv
VITE_SERVER_HOST = http://localhost:3000
```
