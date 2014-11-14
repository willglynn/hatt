hatt
====

Hash all the things! Checksumming tool on crack.

<img src="http://i.imgur.com/n6dFbpv.png" alt="HASH ALL THE THINGS" width="400" height="300">

Example
-------

    cd /path/to/files
    hatt hash -m ../files.hatt

`hatt hash` is used to build a manifest, which includes information about all the files `hatt`
has processed.

Manifests list files along with various fingerprints – MD5, SHA1, SHA2-256, CRC32, etc. – for
the purpose of later comparison. By default, `hatt hash` computes *all* these hashes, on the
theory that CPU time is much cheaper than disk I/O. This way, if you ever need one specific
type of hash, odds are you already have it.

`hatt` is designed for use on large, semi-static datasets. Therefore, in addition to hashes
for each file, the manifest includes the file's modification time and size in bytes. By default,
`hatt hash` skips files with matching mtime and size under the assumption that they haven't
changed, greatly reducing overall I/O requirements. 

`hatt` should really have a command to export a manifest in various formats, like `.md5` and
`.sfv`, but it doesn't do that yet.
