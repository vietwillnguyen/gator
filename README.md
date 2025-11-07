
----
Running On Codespace?
1. 

Install postresql
sudo apt update
sudo apt install postgresql postgresql-contrib

Ensure the installation worked. The psql command-line utility is the default client for Postgres. Use it to make sure you're on version 15+ of Postgres:
psql --version

(Linux only) Update postgres password:
sudo passwd postgres

2. Check sudo Configuration:

It's possible that the codespace user does not have permission to use sudo without entering a password, or that there’s a misconfiguration in the sudoers file. You can check the sudoers file to ensure that the codespace user has the appropriate permissions.

You can edit the sudoers file (if you have root access) with:

sudo visudo


Make sure there’s a line like this allowing codespace to execute sudo:


codespace ALL=(ALL) NOPASSWD: ALL