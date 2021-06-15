
connection pkg.go.dev

1 modify /etc/hosts

# pkg.go.dev
216.58.200.51 pkg.go.dev


2 flush dns

https://www.techrepublic.com/article/how-to-flush-the-dns-cache-on-linux/

How to flush your cache
The first thing we need to do is make sure that systemd-resolved is running. To do that, open a terminal window on your desktop or server and issue the command:

sudo systemctl is-active systemd-resolved
In the output of that command, you should only see:

active
If that's the case, you're okay to proceed. We'll then check some statistics for the DNS cache with the command:

sudo systemd-resolve --statistics
When you run that command, you should see a listing for Transactions, Cache, and DNSSEC Verdicts (Figure A).


The important bit of information is Current Cache Size. We're going to reset that to 0, by flushing  the cache with the command:

sudo systemd-resolve --flush-caches
Once again, issue the command:

sudo systemd-resolve --statistics
You should now see that the Current Cache Size is at 0 (Figure B).