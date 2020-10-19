# What's this tool for?
* You have a workstation you want to develop some stuffs on
* You want to synchronize your content (less than a few Megabytes) to a remote server that can only access the Internet in http/https.
* Then this tool is for you! :)

# Use it!
```bash
# Optionally, compile the Docker image on your side: docker build . -t sdenel/hstransfer 
docker run -v$PWD:/directory_to_upload/ -v$HOME/.ssh/id_rsa:/id_rsa -eHSTRANSFER_SSH_HOST="toto@something.com" -eHSTRANSFER_SSH_PATH="/var/www/apache/somedir/" -eHSTRANSFER_HTTP_PATH="https://something.com/somedir/" -ti sdenel/hstransfer 
# One easy way to use it despite this long command line is with an alias:
alias hstransfer-this-dir='docker run -v$PWD:/directory_to_upload/ -v$HOME/.ssh/id_rsa:/id_rsa -eHSTRANSFER_SSH_HOST="sde@..." -e HSTRANSFER_SSH_PATH="/var/www/..." -eHSTRANSFER_HTTP_PATH="https://download..." -ti sdenel/hstransfer'
```

# Requirements
* You should have Docker installed on your workstation
* You should have access to a public facing static web server that you can access in SSH from your workstation
* The remote server should be able to access this static server
* The remote computer is either on Linux or Windows, 

or from docker if you don't trust me enough:
TODO: compile downloader each time to ensure new hash?
Downloader: ensure new hash each time.

# Added value compared to a bunch of bash commands
* Ability to work on Linux and Windows, without any pre-requirement on the remote side
* Secured:
    * built-in encryption
    * binaries are automatically deleted from the static web server at the end of the session.
  
# See also
* https://hub.docker.com/repository/docker/sdenel/hstransfer

# Backlog...
Contributions welcomed!
* Ability to upload the content using:
    * AWS S3 cli (would allow to use AWS S3, https://www.vultr.com/, https://www.scaleway.com/fr/object-storage/ ...)
    * FTP
* Ability to perform the download even when a Basic Auth is in place
* Delete the binary if asked by the user (by catching some key)
* Hide the binary content as a jpg file (with the right file extension and headers). Potentially even hiding AES encryption with some trick.
* Provide a way to download the downloader in an encrypted way to keep it under the radar.
