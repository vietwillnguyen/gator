# Connect as postgres user (no password needed with sudo)
sudo -u postgres psql

# Connect to your gator database
sudo -u postgres psql -d gator

# Or from your Go application
psql -h localhost -U postgres -d gator
# When prompted, enter: postgres

# Go connection string
connStr := "host=localhost port=5432 user=postgres password=postgresddbname=gator sslmode=disable"