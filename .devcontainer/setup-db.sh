#!/bin/bash
set -e

# Copy config
cp .gatorconfig.json ~/.gatorconfig.json 2>/dev/null || true

# Start PostgreSQL
sudo service postgresql start
sleep 3

# Setup PostgreSQL
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'postgres';"
sudo -u postgres psql -c "ALTER USER postgres WITH SUPERUSER;"
sudo -u postgres createdb -O postgres gator 2>/dev/null || echo "Database gator already exists"

echo "PostgreSQL setup complete!"