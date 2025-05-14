# hookworm
Git commit hook integrator.

* Install hookworm as a strict hook: `hookworm install`
* Run all checks, but only log issues: `hokowork advise`
* Run all checks, fail on issues: `hookworm assert`

## Configuration

Hookworm configuration is a list of checks to be run.



```yml
tasks:
  - name: Sample check, returns true
    command: >-
      /bin/sh -c 'true'
  - name: Containerized true
    command: >-
      docker run ubuntu /bin/sh -c 'true'
```

## Todo

* be able to specify target folder when installing and running
* add mode `hookworm init` which adds a sample task book
* allow marker on task to not take return value into account
* `hookwork install` should place a sample `.hookworm.yml` with a dummy task 

