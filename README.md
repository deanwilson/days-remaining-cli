# Days Remaining CLI

## Introduction

A small and simple command line tool to show the number of days remaining before tasks are due.

## Usage

Create a config file in your location of choice:

```
cat > days-remaining.cfg <<EOF
2025-02-20==Report due
2025-02-28==Payments due
2024-02-18==Expired!
EOF
```

Install the command line tool:

    go install -v github.com/deanwilson/days-remaining-cli

Assuming your `go` path is set correctly you can run `days-remaining-cli`.

```
days-remaining.go days-remaining.cfg

Expired! - Remaining: -366
Report due - Remaining: 2
Payments due - Remaining: 10
```

The days shown will vary based on the date you run the command.

### Author
Dean Wilson <dean.wilson+daysremaining@gmail.com>

### License
GPLv2