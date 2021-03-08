Weather Forcast
---

# Services Description
* Cwb_Crawler : A scheduler service run in background, will get data from open-weather-data in periodicity.
* Web_Forecast : A Web API service, for user to login, and seek forecast data that they want.
* DB : A MySQL database for persistence data.

* Build images
```
build.sh build
```

* Lanuch Services
```
build.sh start
```

* Stop Services
```
build.sh stop
```

* List Services
```
build.sh
```

* Lookup logs of Services
```
build.sh logs
```


# API Example
* Login 

account : testaccount
password : abcd1234

Operation:
Success : will get JWT , and then copy it
```bash
curl -X POST \
  http://localhost:8080/login \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 8038614f-78a5-68f4-7db7-c53c1ffce127' \
  -d '{
    "name": "testaccount",
    "password": "abcd1234"
}'
```

Response example :
```json
{
    "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0ZXN0YWNjb3VudCIsImV4cCI6MTYxNTE4Mjc2NSwianRpIjoidGVzdGFjY291bnQxNjE1MTc5MTY1IiwiaWF0IjoxNjE1MTc5MTY1LCJpc3MiOiJmb3JlY2FzdCJ9.pBa7OCQGncLTxmQgBoj2jomROKBCHQ3Ib5zTOVvPCLU"
}
```

* Forecast Now 
Get the weather forecast data for now.

Operation:
* Paste JWT into authorization header behind Bearer

```bash
curl -X GET \
  http://localhost:8080/forecast/now \
  -H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0ZXN0YWNjb3VudCIsImV4cCI6MTYxNTE4MTEwMSwianRpIjoidGVzdGFjY291bnQxNjE1MTc3NTAxIiwiaWF0IjoxNjE1MTc3NTAxLCJpc3MiOiJmb3JlY2FzdCJ9.Izm7BRRfGAkqmMfIgTYjAdVJ8qIlZbHIclyTzvsgvDQ' \
  -H 'cache-control: no-cache' \
  -H 'postman-token: 93334391-b108-5fc1-1765-80c748118759'
```

Response Example:
```json
{
    "location": [
        {
            "locationName": "臺北市",
            "weatherElement": [
                {
                    "elementName": "Wx",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "陰天",
                                "parameterValue": 7
                            }
                        }
                    ]
                },
                {
                    "elementName": "PoP",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "10",
                                "parameterUnit": "百分比"
                            }
                        }
                    ]
                },
                {
                    "elementName": "CI",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "稍有寒意"
                            }
                        }
                    ]
                },
                {
                    "elementName": "MinT",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "16",
                                "parameterUnit": "C"
                            }
                        }
                    ]
                },
                {
                    "elementName": "MaxT",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "18",
                                "parameterUnit": "C"
                            }
                        }
                    ]
                }
            ]
        },
        {
            "locationName": "新北市",
            "weatherElement": [
                {
                    "elementName": "Wx",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "陰天",
                                "parameterValue": 7
                            }
                        }
                    ]
                },
                {
                    "elementName": "PoP",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "10",
                                "parameterUnit": "百分比"
                            }
                        }
                    ]
                },
                {
                    "elementName": "CI",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "稍有寒意"
                            }
                        }
                    ]
                },
                {
                    "elementName": "MinT",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "17",
                                "parameterUnit": "C"
                            }
                        }
                    ]
                },
                {
                    "elementName": "MaxT",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "19",
                                "parameterUnit": "C"
                            }
                        }
                    ]
                }
            ]
        },
        {
            "locationName": "桃園市",
            "weatherElement": [
                {
                    "elementName": "Wx",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "陰時多雲",
                                "parameterValue": 6
                            }
                        }
                    ]
                },
                {
                    "elementName": "PoP",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "10",
                                "parameterUnit": "百分比"
                            }
                        }
                    ]
                },
                {
                    "elementName": "CI",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "寒冷至稍有寒意"
                            }
                        }
                    ]
                },
                {
                    "elementName": "MinT",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "15",
                                "parameterUnit": "C"
                            }
                        }
                    ]
                },
                {
                    "elementName": "MaxT",
                    "time": [
                        {
                            "startTime": "2021-03-08 06:00:00",
                            "endTime": "2021-03-08 18:00:00",
                            "parameter": {
                                "parameterName": "18",
                                "parameterUnit": "C"
                            }
                        }
                    ]
                }
            ]
        }
    ]
}
```


# Table Introduce
* User : For store user profile
* Location : for store location 
* Element : Store element name and long decription
* Record : Store forecast data for locations
