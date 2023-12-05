# postit
Command Line tool used to run requests against a note.
A note holds all data needed for an API endpoint.

# Example note
We will call it catFact.json
```json
{
  "method": "GET",
  "url": "https://catfact.ninja/fact",
  "headers": {},
  "format": "json",
  "body": {},
  "credentials":{}
}
```

To run this on the CLI all you need to do is:
```shell
./postit -note catFact.json -log.enabled=true -log.directory=.
```

The -note flag is where you point ```postit``` to your note or collection of notes. 
You have the ability to turn on logging and save the request data with true or false on ```-log.enabled```.
The ```-log.directory``` flag points ```postit``` where to log the request data. 
The saved file will be ```.json``` format and have the following Named standard:
```Request-catFact-<Year-Month-Day-HH:MM:S>.json```
The conents look like below:
```json
[{
 "response_headers": {
  "Access-Control-Allow-Origin": [
   "*"
  ],
  "Cache-Control": [
   "no-cache, private"
  ],
  "Content-Type": [
   "application/json"
  ],
  "Date": [
   "Tue, 05 Dec 2023 20:03:39 GMT"
  ],
  "Server": [
   "nginx"
  ],
  "Set-Cookie": [
   "XSRF-TOKEN=eyJpdiI6IllZV2Y3V3hob2hRVGxlSXBNaDBUd3c9PSIsInZhbHVlIjoiZ2E1emVJVXg5L2l6TnRGdFh2eGV2WFJVTUI2ZDNtTWJVSGJ3OFpkQlM5UkMwRjlkbitYNFQvc0FuNFk0aFdZU2pIbUNVMkgvYVp1M0RjekhtQjQ2RlI2TnNVZFRVdkpGQWUzS3lLUjRLbWQvUXlCZUlHZUJzN21Ddkc0aGZsaW4iLCJtYWMiOiJkNTYwMjVjMWVlZjMwZDkzMjNjZTQ3ZjkxZTdmZmJjZTQ0MzA4MzNmYzBiYzUyNjg2ODhiMTY2Y2NjYmViOTU1IiwidGFnIjoiIn0%3D; expires=Tue, 05-Dec-2023 22:03:39 GMT; path=/; samesite=lax",
   "catfacts_session=eyJpdiI6IlNvczlWcXFSbGRBU0lvOUVaUkpNRUE9PSIsInZhbHVlIjoiQks2Uzl3SkRHY3dIM1NWNWJTVGUvUTNxdlZDUTY3QmxucTVFNXA2bk5FdmsyUWNoMC9vbTh1TUZXWjE1dEJDelB4OTZnZi81YzhHVDBPK3AwZlNqSzVNSjBvVDd2R2p2bFJ5bURRK0dFcnNLVStCNDhKd2ZXdXBPYmZJMVJrWGwiLCJtYWMiOiI1Y2FiNDNkZGYzMGYyNmE1NWQwNzlkNjkxMGM0MWIyZmFjZTJmZWMwYjkzNjlkZDhhNWY3NDc0MjhjNGFjZTRmIiwidGFnIjoiIn0%3D; expires=Tue, 05-Dec-2023 22:03:39 GMT; path=/; httponly; samesite=lax"
  ],
  "Vary": [
   "Accept-Encoding"
  ],
  "X-Content-Type-Options": [
   "nosniff"
  ],
  "X-Frame-Options": [
   "SAMEORIGIN"
  ],
  "X-Ratelimit-Limit": [
   "100"
  ],
  "X-Ratelimit-Remaining": [
   "99"
  ],
  "X-Xss-Protection": [
   "1; mode=block"
  ]
 }
},
{
	"response_body": {
		"fact": "A steady diet of dog food may cause blindness in your cat - it lacks taurine.",
		"length": 77
	}
}]
```
