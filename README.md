# gator

A CLI RSS feed aggregator built in Go.

## Dependencies

In order to run gator, you will need to have Postgres and Go installed.

## Installation
### Install via Go

To install gator, run the following command:

`go install github.com/jzaager/gator@latest`

### Build from source

```sh
git clone https://github.com/jzaager/gator.git
cd gator
go build -o gator
```

## Configuration

Create a .gatorconfig.json file in your home directory. The config file should store 2 values in the following format:

```json
{
  "db_url": "postgres://username:password@host:5432/database?sslmode=disable",
}
```
You will need to configure and replace the db_url with your own local postgres instance. [Postgres documentation](https://www.postgresql.org/docs/)

## Usage

`gator [command] <args>`

The following commands are supported:

| Command                            | Description                                                                    |
|------------------------------------|--------------------------------------------------------------------------------|
| register <name>                    | Add a new user                                                                 |
| login <name>                       | Log in an existing user                                                        |
| reset                              | Clear all users, feeds, and follows from database                              |
| users                              | Displays all registered users and current user                                 |
| addfeed <feed_name> <feed_url>     | Add an RSS feed (automatically follows for current user)                       |
| feeds                              | Displays a list of all available feeds                                         |
| follow <feed_url>                  | Follows the feed for the logged in user                                        |
| unfollow <feed_url>                | Unfollows the feed for the logged in user                                      |
| following                          | Displays titles of all feeds the current user is following                     |
| browse <limit [optional]>          | Displays most recent posts from feeds followed by the current user (default=2) |
| agg <time_between_requests>        | Fetches posts from all RSS feeds available                                     |
