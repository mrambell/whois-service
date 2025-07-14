# whois-service
Ever needed a simple way to get information about an IP in JSON format?
This is a simple service to return Whois and ASN information about IPs passed.

## building blocks
Based on 
* whois - [github.com/likexian/whois](https://github.com/likexian/whois)
* BGP View https://bgpview.docs.apiary.io/

The program accepts a single IP as target, or an array of IPs 

examples with curl:
* curl https://localhost:8080/whois\?target=1.1.1.1

* curl -v  "https://api.rambs.de/whois/batch" \
  -H  "Content-Type: application/json" \
  -d '["8.8.4.4", "1.1.1.1", "8.8.8.8", "1.1.1.1"]'

ouput example:
```
{
  "target": "1.1.1.1",
  "result": {
    "% query time": "370 msec",
    "% when": "Mon Jul 14 09:15:14 UTC 2025",
    "% whois data copyright terms    http": "//www.apnic.net/db/dbcopyright.html",
    "abuse-c": "AA1412-AP",
    "abuse-mailbox": [
      "helpdesk@apnic.net",
      "helpdesk@apnic.net"
    ],
    "address": [
      "PO Box 3646",
      "South Brisbane, QLD 4101",
      "Australia",
      "6 Cordelia St",
      "PO Box 3646",
      "South Brisbane, QLD 4101",
      "Australia",
      "6 Cordelia St"
    ],
    "admin-c": [
      "AIC3-AP",
      "AR302-AP",
      "AR302-AP",
      "AIC3-AP"
    ],
    "auth": "# Filtered",
    "country": [
      "AU",
      "AU",
      "ZZ",
      "AU"
    ],
    "descr": [
      "APNIC and Cloudflare DNS Resolver project",
      "Routed globally by AS13335/Cloudflare",
      "Research prefix for APNIC Labs",
      "APNIC Research and Development"
    ],
    "e-mail": [
      "helpdesk@apnic.net",
      "helpdesk@apnic.net",
      "helpdesk@apnic.net",
      "research@apnic.net"
    ],
    "fax-no": "+61-7-38583199",
    "inetnum": "1.1.1.0 - 1.1.1.255",
    "irt": "IRT-APNICRANDNET-AU",
    "last-modified": [
      "2023-04-26T22:57:58Z",
      "2025-05-28T03:31:07Z",
      "2023-09-05T02:15:19Z",
      "2025-05-28T03:31:35Z",
      "2024-07-18T04:37:37Z",
      "2023-04-26T02:42:44Z"
    ],
    "mnt-by": [
      "APNIC-HM",
      "MAINT-APNICRANDNET",
      "APNIC-HM",
      "APNIC-ABUSE",
      "MAINT-APNICRANDNET",
      "MAINT-APNICRANDNET"
    ],
    "mnt-irt": "IRT-APNICRANDNET-AU",
    "mnt-lower": "MAINT-APNICRANDNET",
    "mnt-ref": "APNIC-HM",
    "mnt-routes": "MAINT-APNICRANDNET",
    "netname": "APNIC-LABS",
    "nic-hdl": [
      "AA1412-AP",
      "AIC3-AP"
    ],
    "org": "ORG-ARAD1-AP",
    "org-name": "APNIC Research and Development",
    "org-type": "LIR",
    "organisation": "ORG-ARAD1-AP",
    "origin": "AS13335",
    "phone": [
      "+61-7-38583100",
      "+000000000",
      "+61 7 3858 3100"
    ],
    "remarks": [
      "---------------",
      "All Cloudflare abuse reporting can be done via",
      "resolver-abuse@cloudflare.com",
      "---------------",
      "helpdesk@apnic.net was validated on 2021-02-09",
      "Generated from irt object IRT-APNICRANDNET-AU",
      "helpdesk@apnic.net was validated on 2021-02-09"
    ],
    "role": [
      "ABUSE APNICRANDNETAU",
      "APNICRANDNET Infrastructure Contact"
    ],
    "route": "1.1.1.0/24",
    "source": [
      "APNIC",
      "APNIC",
      "APNIC",
      "APNIC",
      "APNIC",
      "APNIC"
    ],
    "status": "ASSIGNED PORTABLE",
    "tech-c": [
      "AIC3-AP",
      "AR302-AP",
      "AR302-AP",
      "AIC3-AP"
    ]
  },
  "asn": 13335,
  "asn_name": "CLOUDFLARENET"
}
```