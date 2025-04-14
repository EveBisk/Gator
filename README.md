# Gator

Gator is a command-line blog aggregator that allows users to manage RSS feeds, follow blogs, and browse posts. It provides functionality for user management, feed management, and post aggregation.

## Prerequisites

To run Gator, you need the following installed on your system:

1. **PostgreSQL**: Gator uses PostgreSQL as its database. Make sure you have it installed and running.
2. **Go**: Gator is written in Go. Install Go (version 1.23.1 or higher) from [golang.org](https://golang.org).

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/EveBisk/gator.git
   cd gator
   ```

2. Run `go install .` from inside the project3.

## Configuration

1. Create a configuration file: Gator requires a configuration file to connect to the database. Create an env file in the root of the project with the following content:
```
DB_URL="postgres://<username>:<password>@localhost:5432/gator"
```
Replace <username> and <password> with your PostgreSQL credentials.

2. Set up the database:
- Create a PostgreSQL database named gator.
- Run the database migrations using a tool like goose or manually apply the SQL scripts in the schema directory.

## Examples

Here are a few commands you can run with Gator:

- Register a new user: `gator register alice`
- Login as a user: `gator login alice`.
- Add a new RSS feed: `gator addfeed "Tech Blog" "https://example.com/rss"`.
- Follow a feed: `gator follow "https://example.com/rss"`

### Available Commands

| Command             | Description                                                                 |
|---------------------|-----------------------------------------------------------------------------|
| `register <name>`   | Register a new user with the given name.                                   |
| `login <name>`      | Login as an existing user.                                                 |
| `reset`             | Reset the database by deleting all users.                                  |
| `users`             | List all users in the database.                                            |
| `addfeed <name> <url>` | Add a new RSS feed with the given name and URL (requires login).         |
| `feeds`             | List all available feeds.                                                  |
| `follow <url>`      | Follow a feed by its URL (requires login).                                 |
| `following`         | List all feeds the current user is following (requires login).             |
| `unfollow <url>`    | Unfollow a feed by its URL (requires login).                               |
| `browse [limit]`    | Browse posts from followed feeds, optionally limiting the number of posts. |
| `agg <duration>`    | Periodically scrape feeds every `<duration>` (e.g., `10s`, `1m`).          |