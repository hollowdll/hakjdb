# kvdbserver

kvdbserver is the server process that listens to requests from kvdb clients.

# Configuration

Configurations are saved to a configuration file and can be changed there. This file is created with default configurations if it doesn't exist. The name of the configuration file is `.kvdbserver.json` and it is created to the data directory.

# Environment variables

It is also possible to change configurations with environment variables. Environment variables override configurations in the configuration file but the values are not saved. They are only read by the server process.

# Data directory

The server has a data directory `data/` that is created to the executable's parent directory if it doesn't exist. Server specific data is saved to this directory. Currently only configuration file gets saved here.
