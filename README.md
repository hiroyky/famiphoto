# famiphoto

## deploy procedure (current temporary)
```
// boot docker compose
$ make init

// build & debian packaging
$ make dc_pkg

// deploy
$ scp dst/famiphoto_api_xxxxx.deb 192.168.11.30:/tmp
```

at Server:

```
$ cd /tmp
$ sudo dpkg -i ./famiphoto_api_xxxxxx.deb

// Edit env file
$ sudo nano /etc/famiphoto/env

$ sudo systemctl restart famiphoto_api
$ sudo ssytemctl status famiphoto_api -l 

// if error
$ journalctl _PID=xxxxxx
$ tail -f /var/log/famiphoto/error.log
```