# Gator

A CLI-based RSS feed aggregator built with Go, PostgreSQL, Goose, and SQLC.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Goose](https://github.com/pressly/goose)
- [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html)

## Installation

### 1. Install PostgreSQL

**macOS (Homebrew):**
```sh
brew install postgresql
brew services start postgresql
```

**Ubuntu:**
```sh
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo service postgresql start
```

### 2. Install Goose

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### 3. Install SQLC

```sh
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Configuration

### Create a Database in PostgreSQL

1. Open a terminal and start the PostgreSQL interactive terminal:
    ```sh
    psql postgres
    ```
2. Create a new database (replace `dbname`, `user`, and `password` as needed):
    ```sql
    CREATE DATABASE dbname;
    ```
3. Exit `psql`:
    ```sh
    \q
    ```

On mac the user and password will be your local user. On linux you might need to change the password for the postgres user.

### Create a `.gatorconfig.json` file

Create a `.gatorconfig.json` file in your project root:

```json
{
    "database_url": "postgres://user:password@localhost:5432/dbname?sslmode=disable"
}
```

- Make sure not to forget the `?sslmode=disable`.
- Replace `user`, `password`, and `dbname` with your PostgreSQL credentials.

## Build

- Use Goose for database migrations:
    ```sh
    goose -dir sql/schema postgres "$DATABASE_URL" up
    ```
- Use SQLC to generate Go code from SQL:
    ```sh
    sqlc generate
    ```
- Compile the CLI binary:
    ```sh
    go build -o gator
    ```
  This will create an executable named `gator` in your project directory.


## Usage

After building and putting the binary in your `$PATH`, you can run the CLI with:

```sh
gator <command> [args]
```

### Commands

#### `register`

Add a new user to your DB.

```sh
gator register <username>
```

#### `login`

Login as the specified user.

```sh
gator login <username>
```

#### `users`

List all registered users.

```sh
gator users 
```

#### `addfeed`

Add a new RSS feed to your subscriptions.

```sh
gator addfeed <feed-url>
```
- `<feed-url>`: The URL of the RSS feed you want to subscribe to.

#### `agg`

Fetch and store new items from all subscribed feeds.

```sh
gator agg
```
Downloads the latest posts from all feeds and saves them to the database.

#### `feeds`

List all your subscribed-to RSS feeds.

```sh
gator feeds
```
Displays a table of all feeds you have added.

#### `follow`

Follow a specific RSS Feed that has alread been added

```sh
gator follow <feed-url>
```
- `<feed-url>`: The URL of the feed to remove (as shown in `feeds`).

#### `browse`

List posts from your subscribed-to feeds.

```sh
gator browse <limit>
```
- `--feed <limit>`: The number of posts to be shown (default is 2).

#### `following`

Lists the feeds that your user is subscribed to.

```sh
gator following <
```

#### `unfollow`

Unsubscribe from a particular RSS feed. 

```sh
gator unfollow <url>
```
- `--feed <url>`: The url of the feed you want to unsubscribe to.
