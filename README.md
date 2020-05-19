# API-scenario
## Scenario API testing from the command line.

![Work in progress](https://img.shields.io/badge/status-work%20in%20progress-red "Work in progress")
[![Build Status](https://travis-ci.com/thomaspoignant/api-scenario.svg?token=sVd5BLjwtrGWjxxeoYSx&branch=master)](https://travis-ci.com/thomaspoignant/api-scenario)
![Release version](https://img.shields.io/github/v/release/thomaspoignant/api-scenario "Release version")
![Go version](https://img.shields.io/github/go-mod/go-version/thomaspoignant/api-scenario?logo=go%20version "Go Version")
 
API-scenario is a simple command line tool that allow you to execute easily a scenario to test your APIs.  
It is perfect to make end to end tests, you could use in your CI workflow to validate your deployment or just to test locally your development.

---
# Why this project?

---


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

 	
 	 	
