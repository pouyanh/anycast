# Anycast Race Condition
Distributed race condition prevention

[![Build Status](https://travis-ci.com/pouyanh/anycast.svg?branch=master)](https://travis-ci.com/pouyanh/anycast)
[![Coverage Status](https://coveralls.io/repos/github/pouyanh/anycast/badge.svg?branch=master)](https://coveralls.io/github/pouyanh/anycast?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pouyanh/anycast/die-github-cache-die)](https://goreportcard.com/report/github.com/pouyanh/anycast)

Solutions to handle race conditions in ecosystems consisting of 3 types of actors: client, servant and platform

Clients will make requests and only one servant can accept the request to process it

## Contents
* [Problem](#problem)
* [Solution](#solution)
* [Run](#run)

## Problem
Multiple servants want to serve same client at the same time

## Solution

## Run
1. Run the environment
```bash
docker-compose up -d
```

2. Monitor interactions through browser
