# API-scenario
## Scenario API testing from the command line.

![Work in progress](https://img.shields.io/badge/status-work%20in%20progress-red "Work in progress")
[![Build Status](https://travis-ci.com/thomaspoignant/api-scenario.svg?token=sVd5BLjwtrGWjxxeoYSx&branch=master)](https://travis-ci.com/thomaspoignant/api-scenario)
![Release version](https://img.shields.io/github/v/release/thomaspoignant/api-scenario "Release version")
![Go version](https://img.shields.io/github/go-mod/go-version/thomaspoignant/api-scenario?logo=go%20version "Go Version")
 
API-scenario is a simple command line tool that allow you to execute easily a scenario to test your APIs.  
It is perfect to make end to end tests, you could use in your CI workflow to validate your deployment or just to test 
locally your development.

---
# Why this project?
Our goal is to have a simple command line tool that can run scenario to test your apis directly from the command line.
You can create a scenario and start using it during development, but also use it in your
 [CI/CD](https://en.wikipedia.org/wiki/CI/CD) to validate your deployment. 

You can use variables in your scenario, so you are sure to run the same test suite on each environment and be sure of 
your release.
 
---
# Creating Your First Test
Creating a test is simple, you just have to write `json` file to describe you api calls, and describe assertions.

## The basic structure of the file is:
```json
{
  "name": "Simple API Test Example",
  "description": "A full description of your test scenario",
  "version": "1.0",
  "steps": [
    ...
  ]
}
```

This global fields allow to describe your scenario:
- **name**: The name of your scenario
- **description**: A complete description of what your scenario is doing
- **version**: The version of your scenario
- **steps**: Array of steps, it will describe all the steps of your scenario _(see [steps](#steps) for more details)_.

## Our first step
For our first step we will create a basic call who verify that an API answer with http code `200` when calling it.
```json
{
  "step_type": "request",
  "url": "{{baseUrl}}/api/users",
  "method": "GET",
  "headers": {
    "Content-Type": ["application/json"]
  },
  "assertions": [
    {
      "comparison": "equal_number",
      "value": "200",
      "source": "response_status"
    }
  ]
}
```
We manipulate different concepts here.
- **step_type**: The type of the step, here we are using request because we want to test a rest 
    API _(see [steps](#steps) to see the list of available step types)_.
- **url**: The URL of our request, `{{baseUrl}}` will be replaced before the call _([see Using Variables in Requests for details](#using-variables-in-requests))_
- **method**: The HTTP verb of our request.
- **headers**: The list of headers we sent with the request.
- **assertions**: This is the list of checks we are doing when we have received the response.
    - **comparison**: The type of check we are doing, here we are testing that number are equals.
    - **value**: The expected value
    - **source**: On what part of the response we are looking, here we are checking the response status.


Now the first scenario is build, we can run it _([see complete scenario](examples/first-test.json))_.
```shell script
$ api-scenario run --scenario="examples/first-test.json" --variable="baseUrl:https://reqres.in"
```
What we are doing is here is running our scenario file, and we ask to replace every occurrence 
of `{{baseUrl}}` by `https://reqres.in`.

---
# Steps

---
# Request Chaining
## Using Variables to Pass Data Between Steps
Request steps can define variables that extract data from HTTP responses returned when running the test.
To create a variable, add a `variables` block to your step and specify the location of the data you'd like to extract 
from the response, and the **name** of this variable.

```json
"variables": [
    {
      "source": "response_json",
      "property": "point",
      "name": "active"
    }
  ]
```
**A variable is:**

|                   |  |
|---                |---
|**Source**         |The location of the data to extract. Data can be extracted from<br><ul><li>HTTP header values - `response_header`</li><li>Response bodies - `response_json`</li><li>Response status code - `response_status`</li></ul>
|**Property**       |The property of the source data to retrieve.<br>For HTTP headers this is the name of the header.<br>For JSON content, see below.<br>Unused status code.
|**Variable Name**  |The name of the variable to assign the extracted value to.<br>In subsequent requests you can retrieve the value of the variable by this name.<br>[See Using Variables in Requests](#using-variables-in-requests).

## Extracting Data from JSON Body Content
Data from a JSON response body can be extracted by specifying the path of the target data using standard JavaScript 
notation. [View sample JSON](./examples). 

## Global Variables
Some variables could be set up at launch, for that you can add options to the `run` command to pass it.
Common values (base URLs, API tokens, etc.) that are shared across requests within a test, or tests within a bucket, 
should be stored in an Initial Variable.
Once defined, the variable is available to all requests within the test.

To add a variable just use the option `--variable` or `-V` and specify the `key:value` of this variable.
```shell script
$ ./api-scenario run -F your_file.json --variable="baseUrl:http://www.google.com/" -V "token:token1234"
```

__Note that if you create a variable in a step with the same name of a global variable it will override it.__

## Add / Override headers
Overriding headers works the same as [global variables](#global-variables).  
You can add a header for all your requests by using the option `--header` or `-H`, it will add or override the header 
for all requests.

```shell script
$ ./api-scenario run -F your_file.json --header="Content-Type:application/json" -H "Authorization: Bearer Token123"
```

## Built-in Variables and Functions
|**Variable/Function**       |**Description**       |**Example Output**
|---	                 |---	            |---	
|**`{{timestamp}}`**                        |Integer Unix timestamp (seconds elapsed since January 1, 1970 00:00 UTC)       |`1384035195`
|**`{{utc_datetime}}`**                     |UTC datetime string in ISO 8601 format.                                        |`2013-11-07T19:24:41.418968`
|**`{{format_timestamp(value, format)}}`**  |Timestamp of the specified value in the specified format.<br>Any delimiters (e.g. -, /, ., *, etc.) can be used in the format with a combination of any of the following date/time format options. Also accepts variables. E.g. **`{{format_timestamp({{timestamp}}, YYYY-MM-DD)}}`**<br><ul><li>**YYYY** - 4 digit year (e.g. 2019)</li><li>**YYYY** - 4 digit year (e.g. 2016)</li><li>**YY** - 2 digit year (e.g. 16)</li><li>**MM** - month</li><li>**DD** - day</li><li>**HH** - 24 hour (e.g. 13 == 1pm)</li><li>**hh** - 12 hour (e.g. 01 == 1pm)</li><li>**mm** - minutes</li><li>**ss** - seconds</li></ul>|`2019-31-03`
|**`{{timestamp_offset(value)}}`**          |Integer Unix timestamp offset by the specified value in seconds (going back in time would be a negative offset value). Values should be passed without surrounding quotes.|`1383948795`
|**`{{random_int}}`**                       |Random integer between 0 and 18446744073709551615                              |`407370955`
|**`{{random_int(a,b)}}`**                  |Random integer value between a and b, inclusive.                               |`44674407370`
|**`{{random_string(length)}}`**            |Random alphanumeric string of the specified length (max 1000 characters).      |`ddo1qlQR81`
|**`{{uuid}}`**                             |Random universally unique identifier (UUID). 	                                |`99386c08-6da7-4833-bb31-e70ce747c921`
|**`{{encode_base64(value)}}`**             |Encodes value in Base64. Values should be passed without surrounding quotes. Also accepts variables e.g. `{{encode_base64({{username}}:{{password}})}}` |`dTpwDQo`=
|**`{{md5(value)}}`**                       |Generate an MD5 hash based on value. Values should be passed without surrounding quotes. Also accepts variables e.g. `{{md5({{timestamp}})}}` |`50b7fe4da64720232c25bc7c6d66f6c5`
|**`{{sha1(value)}}`**                      |Generate an SHA-1 hash based on value. Values should be passed without surrounding quotes. Also accepts variables e.g. `{{sha1({{timestamp}})}}` |`e0bd9304537cd8cb4e69ef5d73771fe218c484f5`
|**`{{sha256(value)}}`**                    |Generate an SHA-256 hash based on value. Values should be passed without surrounding quotes. Also accepts variables e.g. `{{sha1({{timestamp}})}}` |`e3376ffb4b1e2c04b0fe68b52e8654696814b4883b47a56ff5a7df883725d8c1`
|**`{{hmac_sha1(value,key)}}`**             |Generate an HMAC using the SHA-1 hashing algorithm based on value and key. Values should be passed without surrounding quotes. Also accepts variables e.g. `{{hmac_sha1({{timestamp}},key)}}` |`163a04cd86a82b948a7e85f0ed3cd3b5929a7d0c`
|**`{{hmac_sha256(value,key)}}`**           |Generate an HMAC using the SHA-256 hashing algorithm based on value and key. Values should be passed without surrounding quotes. Also accepts variables e.g. `{{hmac_sha1({{timestamp}},key)}}` |`eb0b5c5b2a04ac25ff52c886e115f2e60c0dd8d50bab076dc065e95f5fd37fb9`
|**`{{url_encode(value)}}`**                |Create a percent-encoded string suitable for URL querystrings. This is not required for URL or form parameters defined in the request editor which are automatically encoded. Only use this if you need to double encode a value in a URL or include a URL encoded string in a header value. 	|`This%20is%20100%25%20URL%20encoded.`

 	
## Using Variables in Requests
Once a variable has been defined, you can use it in any subsequent request.  
Variables can be used in any request data field including the method, URL, header values, parameter values and request bodies.

To include the value of a variable in a request, enter the name of the variable surrounded by double braces e.g. **`{{variable_name}}`**.  
If a variable is undefined when a request using that variable is executed, we will keep the variable un-replaced.
