[![Build Status](https://travis-ci.org/wwcd/httpcli.svg?branch=master)](https://travis-ci.org/wwcd/httpcli)

# Overview

Simple httpcli for go.

- Default timeout 30 seconds
- Default skip SSL verify

# Usage

	rsp, err := httpcli.Get(context.Background(), "https://github.com")
