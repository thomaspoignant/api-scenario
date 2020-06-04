# API-scenario
## Scenario API testing from the command line.

[![Release version](https://img.shields.io/github/v/release/thomaspoignant/api-scenario "Release version")](https://github.com/thomaspoignant/api-scenario/releases)
[![Build Status](https://travis-ci.com/thomaspoignant/api-scenario.svg?token=sVd5BLjwtrGWjxxeoYSx&branch=master)](https://travis-ci.com/thomaspoignant/api-scenario)
[![Coverage Status](https://coveralls.io/repos/github/thomaspoignant/api-scenario/badge.svg)](https://coveralls.io/github/thomaspoignant/api-scenario)
[![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=thomaspoignant_api-scenario&metric=alert_status)](https://sonarcloud.io/dashboard?id=thomaspoignant_api-scenario)
![Go version](https://img.shields.io/github/go-mod/go-version/thomaspoignant/api-scenario?logo=go%20version "Go version")
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fthomaspoignant%2Fapi-scenario.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fthomaspoignant%2Fapi-scenario?ref=badge_shield)
 
API-scenario is a simple command line tool that allow you to execute easily a scenario to test your APIs.  

It is perfect to make end to end tests, you could use in your CI workflow to validate your deployment or just to test 
locally your development.


![api-scenario demo](api-scenario.gif "api-scenario demo")
---

- [Why this project?](#why-this-project)
- [Installation](#installation)
  - [Install using Homebrew (mac and linux)](#install-using-homebrew-mac-and-linux)
  - [Install using Scoop (windows)](#install-using-scoop-windows)
  - [Install using .deb file (linux)](#install-using-deb-file-linux)
  - [Install using .rpm file (linux)](#install-using-rpm-file-linux)
  - [Use docker image](#use-docker-image)
- [Command line usage](#command-line-usage)
  - [Version](#version)
  - [Execute your scenario](#execute-your-scenario)
    - [Save result into file](#save-result-into-file)
- [Creating Your First Test](#creating-your-first-test)
  - [The basic structure of the file](#the-basic-structure-of-the-file)
  - [Our first step](#our-first-step)
- [Steps](#steps)
  - [Pause](#pause)
  - [Request](#request)
    - [Headers](#headers)
    - [Assertions](#assertions)
      - [Assertion compostion](#assertion-compostion)
      - [Available comparison type](#available-comparison-type)
- [Request Chaining](#request-chaining)
  - [Using Variables to Pass Data Between Steps](#using-variables-to-pass-data-between-steps)
  - [Extracting Data from JSON Body Content](#extracting-data-from-json-body-content)
  - [Global Variables](#global-variables)
  - [Add / Override headers](#add--override-headers)
  - [Built-in Variables and Functions](#built-in-variables-and-functions)
  - [Using Variables in Requests](#using-variables-in-requests)

---
# Why this project?
Our goal is to have a simple command line tool that can run scenario to test your apis directly from the command line.
You can create a scenario and start using it during development, but also use it in your
 [CI/CD](https://en.wikipedia.org/wiki/CI/CD) to validate your deployment. 

You can use variables in your scenario, so you are sure to run the same test suite on each environment and be sure of 
your release.

---

# Installation
All the binaries are available in the [release](https://github.com/thomaspoignant/api-scenario/releases).

## Install using Homebrew (mac and linux)
If you are using Homebrew package manager, you can install api-scenario with Homebrew. 
 
```shell script
brew tap thomaspoignant/homebrew-tap
brew install api-scenario
```

## Install using Scoop (windows)

If you are on Windows and using scoop package manager, you can install api-scenario with [Scoop](https://scoop.sh/).
```shell script
scoop bucket add org https://github.com/thomaspoignant/scoop.git
scoop install api-scenario
```

## Install using .deb file (linux)
If you prefer to use directly the `.deb` file to install in your debian like linux distribution.  
_Don't forget to set the correct version number._
```shell script
wget https://github.com/thomaspoignant/api-scenario/releases/download/vX.X.X/api-scenario_X.X.X_Tux_64-bit.deb && dpkg -i api-scenario_*.deb
```

## Install using .rpm file (linux)
If you prefer to use directly the `.rpm` file to install in your centos/fedora like linux distribution.  
_Don't forget to set the correct version number._
```shell script
wget https://github.com/thomaspoignant/api-scenario/releases/download/vX.X.X/api-scenario_X.X.X_Tux_64-bit.rpm && rpm -i api-scenario_*.rpm
```

## Use docker image
To use the last version of our docker image you can pull `thomaspoignant/api-scenario:latest`.

```shell script
docker pull thomaspoignant/api-scenario:latest
```

You can also pull all the version directly but also directly a major version.  
Available images are:
- `docker pull thomaspoignant/api-scenario:latest` _(target the latest version)_
- `docker pull thomaspoignant/api-scenario:vX` _(target the major version)_
- `docker pull thomaspoignant/api-scenario:vX.X` _(target the minor version)_
- `docker pull thomaspoignant/api-scenario:vX.X.X` _(target a specific version)_

---
# Command line usage

## Version
```shell script
$ api-scenario version
# 0.1.0
```

## Execute your scenario
To execute your scenario please use the `run` options and specify your scenario file.

```shell script
$ api-scenario run --scenario="./scenario.json"
```

There are several options you can use with this command:

|Option                  |Short version  | Required |Description  |
|---                     |---            |---       |---
|`--scenario`            | `-s`          |✓         |Input file for the scenario.
|`--authorization-token` | `-t`          |          |Authorization token send in the Authorization headers.
|`--header`              | `-h`          |          |Header you want to override (format should be "**header_name:value**").<br>*You can have multiple values of this options*
|`--variable`            | `-h`          |          |Value for a variable used in your scenario (format should be "**variable_name:value**").<br>*You can have multiple values of this options*
|`--verbose`             | `-s`          |          |Run your scenario with debug information.
|`--quiet`               | `-s`          |          |Run your scenario in quiet mode.
|`--no-color`            |               |          |Do not display color on the output.
|`--output-file`         | `-f`          |          |Output file where to save the result _(use `--output-format` to specify if you want `JSON` or `YAML` output)_.
|`--output-format`       |               |          |Format of the output file, available values are `JSON` and `YAML` _(ignored if `--output-file` is not set, default value is `JSON`)_.

### Save result into file
To keep history of your scenario execution you can export the results into a file.
You just have to add the option `--output-file="<your file location>"` and it will save the result into a `JSON` file 
_(If you prefer `YAML` result add `--output-format=YAML`)_.

**Example:**
```shell script
$ api-scenario run --scenario="./scenario.json" --output-file="<your file location>" --output-format=YAML
```

---
# Creating Your First Test
Creating a test is simple, you just have to write `json` or `yaml` file to describe you api calls, and describe assertions.

## The basic structure of the file
<details open>
<summary><b>YAML</b></summary>
  
  ```yaml
  name: Simple API Test Example
 description: A full description ...
 version: '1.0'
 steps:
     - ...
  ```
</details>

<details>
 <summary><b>JSON</b></summary>
  
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
</details>

This global fields allow to describe your scenario:
- **name**: The name of your scenario
- **description**: A complete description of what your scenario is doing
- **version**: The version of your scenario
- **steps**: Array of steps, it will describe all the steps of your scenario _(see [steps](#steps) for more details)_.

## Our first step
For our first step we will create a basic call who verify that an API answer with http code `200` when calling it.

<details open>
<summary><b>YAML</b></summary>
  
  ```yaml
  - step_type: request
    url: {{baseUrl}}/api/users
    method: GET
    headers:
      Content-Type:
        - application/json
    assertions:
      - comparison: equal_number
        value: '200'
        source: response_status
  ```
</details>

<details>
 <summary><b>JSON</b></summary>
  
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
</details>

We manipulate different concepts here.
- **step_type**: The type of the step, here we are using request because we want to test a rest 
    API _(see [steps](#steps) to see the list of available step types)_.
- **url**: The URL of our request, `{{baseUrl}}` will be replaced before the call _([see Using Variables in Requests for details](#using-variables-to-pass-data-between-steps))_
- **method**: The HTTP verb of our request.
- **headers**: The list of headers we sent with the request.
- **assertions**: This is the list of checks we are doing when we have received the response.
    - **comparison**: The type of check we are doing, here we are testing that number are equals.
    - **value**: The expected value
    - **source**: On what part of the response we are looking, here we are checking the response status.


Now the first scenario is build, we can run it _(see complete scenario:  [YAML](examples/first-test.yml) / [JSON](examples/first-test.json))_.
```shell script
$ api-scenario run --scenario="examples/first-test.json" --variable="baseUrl:https://reqres.in"
```
What we are doing is here is running our scenario file, and we ask to replace every occurrence 
of `{{baseUrl}}` by `https://reqres.in`.

---
# Steps
There are different types of steps who allow performs different types of actions.  
To specify the type of a step we are using the property `step_type` is the step object. 

_If there is no step_type property we ignore the step._

## Pause
**`pause`** is simple, it is a step that wait X seconds.  
This is useful when you have asynchronous API and allows waiting before calling the next API in the scenario.

|Parameters      |Description  |
|---            |---
|**step_type**      | `pause`
|**duration**       | Number of seconds to wait.

**Example:** _Wait for 5 seconds_
<details open>
<summary><b>YAML</b></summary>
  
  ```yaml
  - step_type: pause
    duration: 5
  ```
</details>

<details>
 <summary><b>JSON</b></summary>
  
```json
{
  "step_type": "pause",
  "duration": 5
}
```
</details>

## Request
**`request`** is the step who can call a REST Api.  

|Parameters      |Description  |
|---            |---
|**step_type**      | `request`
|**url**            | URL of your endpoint
|**method**         | HTTP verb of your request _(GET, POST, PUT, DELETE, OPTIONS, PATCH)_
|**variables**      | Array of variables to extract from the response _([see Using Variables to Pass Data Between Steps for details](#))_
|**headers**        | Object who contains all the headers attach to the request _([see how to add headers](#headers))_
|**assertions**     | Array of assertions, this is the acceptance tests _([see how to create assertion tests](#assertions))_

### Headers
Headers are represented by an object containing all the headers to send.  
Each header is has the name of the header for key and an array of strings as value.

You can use variables in the headers, they will be replaced before sending the request _(see [Using Variables to Pass Data Between Steps](#using-variables-to-pass-data-between-steps) or [Global Variables](#global-variables))_.

**Example:**
<details open>
<summary><b>YAML</b></summary>
  
  ```yaml
headers:
      Accept-Charset:
        - utf-8
      Accept: 
        - application/scim+json
      Authorization:
        - {{auth}}
  ```
</details>

<details>
 <summary><b>JSON</b></summary>
  
```json
{
  "headers": {
      "Accept-Charset": [
        "utf-8"
      ],
      "Accept": [
        "application/scim+json"
      ],
      "Authorization": [
        "{{auth}}"
      ]
    }
}
```
</details>

### Assertions
Assertions are a big part of api-scenario, this is the acceptance tests of your request, it will allow you to simply write test to verify that you endpoint is doing what you want.

#### Assertion compostion

|Property        |Description  |
|---             |---
|**source**      | The location of the data to extract for comparison.<br>Authorized values are:<br><ul><li>response_status</li><li>response_time</li><li>response_json</li><li>response_header</li></ul>
|**comparison**  | The type of operation to perform when comparing the extracted data with the target value.  _([see Available comparison type](#available-comparison-type))_.
|**property**    | The property of the source data to retrieve.<br><ul><li>For **HTTP headers**, this is the name of the header.</li><li>Data from a **JSON** response body can be extracted by specifying the path of the data using standard JavaScript notation.</li><li>Unused for text content, status code, response time and response size.</li>
|**value**       | The expected value used to compare against the actual value. 


**Example:**
<details open>
<summary><b>YAML</b></summary>
  
  ```yaml
 - comparison: equals
   property: schemas
   value: User
   source: response_json
  ```
</details>

<details>
 <summary><b>JSON</b></summary>
  
```json
{
  "comparison": "equals",
  "property": "schemas",
  "value": "User",
  "source": "response_json"
}
```
</details>

#### Available comparison type

|Comparison         |Config name                 |Description  |
|---                |---                         |---
|**is empty** 	    |`empty`                     |The actual value exists and is an empty string or null.
|**is not empty** 	|`not_empty`                 |The actual value exists and is a value other than an empty string or null.
|**equals** 	    |`equal`                     |A string comparison of the actual and expected value. Non-string values are cast to a string before comparing. For comparing non-integer numbers, use equals (number).
|**does not equal** |`not_equal`                 |A string comparison of the actual and target value.
|**contains** 	    |`contains`                  |The actual value contains the target value as a substring.
|**does not contain** 	|`does_not_contains`     |The target value is not found within the actual value.
|**has key**        |`has_key`                   |Checks for the existence of the expected value within a dictionary's keys. The actual value must point to a dictionary (JSON only).
|**has value** 	    |`has_value`                 |Checks a list or dictionary for the existence of the expected value in any of the list or dictionary values. The actual value must point to a JSON list or dictionary (JSON only).
|**is null**        |`is_null`                   |Checks that a value for a given JSON key is null.
|**is a number**    |`is_a_number`               |Validates the actual value is (or can be cast to) a valid numeric value.
|**less than** 	    |`is_less_than`              |Validates the actual value is (or can be cast to) a number less than the target value.
|**less than or equal** |`is_less_than_or_equals`|Validates the actual value is (or can be cast to) a number less than or equal to the target value.
|**greater than** 	|`is_greater_than`           |Validates the actual value is (or can be cast to) a number greater than the target value.
|**greater than or equal** 	|`is_greater_than_or_equal`|Validates the actual value is (or can be cast to) a number greater than or equal to the target value.
|**equals (number)** 	|`equal_number`          |Validates the actual value is (or can be cast to) a number equal to the target value. This setting performs a numeric comparison: for example, "1.000" would be considered equal to "1".



---
# Request Chaining
## Using Variables to Pass Data Between Steps
Request steps can define variables that extract data from HTTP responses returned when running the test.
To create a variable, add a `variables` block to your step and specify the location of the data you'd like to extract 
from the response, and the **name** of this variable.

<details open>
<summary><b>YAML</b></summary>
  
  ```yaml
    variables:
      - source: response_json
        property: point
        name: active
  ```
</details>

<details>
 <summary><b>JSON</b></summary>
  
```json
{
  "variables": [
    {
      "source": "response_json",
      "property": "point",
      "name": "active"
    }
  ]
}
```
</details>

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
