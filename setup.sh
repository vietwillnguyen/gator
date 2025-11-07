# install Go
curl -sS https://webi.sh/golang | sh

# For Linux/WSL
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.bashrc
# next, reload your shell configuration
source ~/.bashrc

# Complete Bootdev
go install github.com/bootdotdev/bootdev@latest
bootdev login

sudo apt update
sudo apt install postgresql postgresql-contrib

cp .gatorconfig.json ~/.gatorconfig.json

# Goose is a database migration tool written in Go. 
# It runs migrations from a set of SQL files, 
# making it a perfect fit for this project (we wanna stay close to the raw SQL).
go install github.com/pressly/goose/v3/cmd/goose@latest

# SQLC is an amazing Go program that generates Go code from SQL queries. 
# It's not exactly an ORM, but rather a tool that makes working with raw SQL easy and type-safe.
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest


# postgres://postgres:postgres@localhost:5432/gator
