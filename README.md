# CalDAV2ICS

A super basic golang API server to let you easily share a CalDAV calendar (e.g. from Radicale) as an ICS url with a friend, no annoying cron jobs or anything. Simply set up an .env file (or environment variables) and run the binary and they'll be able to grab an ICS file from `/ics/<href>.ics?token=<token>`.