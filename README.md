# go-check-tcp-status

## Overview

It is a tool for alerting based on the state of a specific TCP. When TIME_WAIT exceeds 1000 or when SYN_RCV becomes a high number, etc.

## Usage

```
Usage:
  check-tcp-status [OPTIONS]

Application Options:
  -s, --status=   tcp status
                  TCP_ESTABLISHED|TCP_SYN_SENT|TCP_SYN_RECV|TCP_FIN_WAIT1|TCP_FIN_WAIT2
                  TCP_TIME_WAIT|TCP_CLOSE|TCP_CLOSE_WAIT|TCP_LAST_ACK|TCP_LISTEN|TCP_CLOSING|TCP_NEW_SYN_RECV
  -w, --warning=  Warning threshold (num)
  -c, --critical= Critical threshold (num)
  -d, --debug

Help Options:
  -h, --help      Show this help message
```
