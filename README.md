# MiraiSafe

Mirai can take your whole network down and you probably don't want to have it happened.
Use this tool as a part of The Porte Solutions Open Source Packet

## Structure

Since Mirai spread first by first entering a quick scanning stage where it proliferates by haphazardly
sending TCP SYN probes to pseudo-random IPv4 addresses, on Telnet TCP ports 23 and 2323, then we
created a check for those ports.

Once Mirai discovers open Telnet ports, it tries to infect the devices by brute forcing the login
credentials. Mirai tries to login using a list of ten username and password combinations. These ten
combinations are chosen randomly from a pre-configured list 62 credentials which are frequently used as
the default for IoT devices.

After successfully logging in, Mirai sends the victim IP and related credentials to a reporting server.
Initially, Mirai tries to assess and identify the environment in which it is running. This information
is then used to download second stage payloads and device specific malware. For instance, the payload 
for a ARM based device will be different than a MIPS one.

After successfully infecting a device, Mirai covers its tracks by deleting the downloaded binary and 
using  a pseudo-random alphanumeric string as its process name. As a result, Mirai infections do not 
persist after system reboots. So as to strengthen itself, the malware also terminates different services 
which are bound to TCP/22 or TCP/23, including other Mirai variations. At this point, the bot waits for 
commands from it’s command and control server (C2) while at the same time looking out for other 
vulnerable devices.

