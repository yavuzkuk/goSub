# gokuk scan

This tool created with [Golang](https://go.dev/). With this tool, you can scan the website you provide as parameters.


- DNS Record
    - MX Record
    - NS Record
    - A Record (IPv4)
    - AAAA Record (IPv6)
- Check robots.txt file
- Subdomain Scan
- Directory Scan
- Web Server Infos

```
-d, --DNS Record Type string           A Record: IPv4 address
                                         AAAA Record:Ipv6 address
                                         MX Record:Mail record
                                         NS Record:Name server record (default "A-AAAA-MX-NX")
  -f, --Filter HTTP Status Code string   You can filter HTTP Statsus Code with -f parameter (default "200,404")
  -c, --count int                        Request count (default 10)
  -h, --help                             help for Cyrops
  -r, --robots.txt                       With default value the tool check the robots.txt file (default true)
  -u, --url string                       You need to specify URL
  -v, --version                          version for Cyrops
  -w, --wordlist string                  You can specify Wordlist (default "wordlist/seclistWebContent.txt")
```