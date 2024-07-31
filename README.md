# gokuk scan

Bu araç [Golang](https://go.dev/) ile oluşturulmuştur. Bu araçla, parametre olarak verdiğiniz web sitesini tarayabilirsiniz.

------

This tool created with [Golang](https://go.dev/). With this tool, you can scan the website you provide as parameters.


- DNS Record
    - MX Record
    - NS Record
    - A Record (IPv4)
    - AAAA Record (IPv6)
    - TXT Record
- Check robots.txt file
- Subdomain Scan
- Directory Scan
- Web Server Infos

 ### Parameters

```
 -d, --DNS Record Type string           A Record: IPv4 address
                                         AAAA Record:Ipv6 address
                                         MX Record:Mail record
                                         NS Record:Name server record
                                         TXT Record:Domain info text (default "A-AAAA-NS-MX-TXT")
  -f, --Filter HTTP Status Code string   You can filter HTTP Statsus Code with -f parameter (default "200,404")
  -c, --count int                        Request count (default 10)
  -h, --help                             help for Cyrops
      --no-robots                        Disable the robots.txt feature
  -r, --robots                           With default value the tool check the robots.txt file (default true)
  -s, --subdomain-wordlist string        You can specify Subdomain Wordlist (default "wordlist/seclistSubdomains5000.txt")
  -u, --url string                       You need to specify URL
  -v, --version                          version for Cyrops
  -w, --wordlist string                  You can specify Directory Wordlist (default "wordlist/seclistWebContent.txt")
```