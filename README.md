# Initial Avatars
This package creates an avatar image based off the initials given and colors that are passed in. It will output something similar to googles default avatar for a user.

## Commands
There are two commands that use this package. `Avatar` is a command line tool that create the svg file (Not implemented yet). `AvatarServer` is an API endpoint that can be used by a web service to get an svg avatar on the fly.

## Using the server
The only endpoint that is open is `/avatar/{initials}`. `initials` is used as the text inside the circle. You can tack on the `color` query parameter like so `/avatar/wd?color=ff0000` and give it a 6 digit hex value for the color you want to use. If a `color` is not supplied it will generate a random color for you.

## Development
If a `.env` file exist in the root directory of where the server is running, it will update all the environment variables you set in it. Currently it is only looking for port.
```
AVATAR_PORT = 8080
```
This will bind the server to listen on port `8080`. If not variable is found then port `80` is used by default.

## Systemd Daemon
**NOTE: Must compile before trying to run it.**

I have included a systemd service file that you can use to set this server to run on system boot. If you are using this file it is recommended that it is put in `/usr/lib/systemd/system` and any changes you want to make are put in `/etc/systemd/system`. In the later I would change the `ExecStart` to point to where you installed the server (default is `/usr/bin/AvatarServer`). Also add an `Environment` variable here for `AVATAR_PORT` if you would like for the server to listen on something other than port `80`.

So for example:
1. Create a folder `/etc/systemd/system/avatarServer.service.d/`
2. Create file in this folder named `avatarServer.conf` (only requirement is it is suffixed with `.conf`)
3. Add settings that you would like to include or override with the default file (the one located in `/usr/lib/systemd/system`)
4. Such a file may look like this:
  ```
[Service]
Environment="AVATAR_PORT=8080"
ExecStart=/home/username/bin/AvatarServer
Requires=apache2.service
  ```
This file would have AvatarServer listening on port `8080`, changes where the server is found, and also requries that apache2 be running.

[Read more here](https://www.digitalocean.com/community/tutorials/systemd-essentials-working-with-services-units-and-the-journal) to learn how to use systemd.
