# go-kornflake

[![CircleCI](https://circleci.com/gh/Zenithar/go-kornflake.svg?style=svg)](https://circleci.com/gh/Zenithar/go-kornflake)

Distributed snowflake identifier microservice

## Example

```
± echo "{}" | bin/kornflake cli bigflakeAPI get -s localhost:5555
{
  "identifier": "28724559237309938838027238375425"
}%                                                                                                                                                                                                                                                                            
± echo "{}" | bin/kornflake cli snowflakeAPI get -s localhost:5555
{
  "identifier": "972178778925039616"
}%  
```