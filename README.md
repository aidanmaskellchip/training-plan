# Training Plan Creator

## Overview
This project is in the early stages of its life. In time, it will become a creator of running training plans, based 
on your current abilities and goals.

## Local development

### Setup
- Run `make setup`

### Local Database
- Upon running `make setup`, the local postgres database is accessible through the port 8432.
- The credentials for accessing the db can be found in the `make setup-db` command in the `Makefile`

### Database Utils
Migrations can be invoked by running `make run-migrations`

### Testing
Run 'make go-tests'