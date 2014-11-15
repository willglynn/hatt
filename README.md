hatt
====

Hash all the things! Checksumming tool on crack.

<img src="http://i.imgur.com/n6dFbpv.png" alt="HASH ALL THE THINGS" width="400" height="300">

Example
-------

    cd /path/to/files
    hatt hash -m ../files.hatt
    cd ..
    hatt export -m files.hatt -o files.md5sum

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

`hatt hash` can be interrupted with SIGINT (`^C`), in which case it writes the in-progress
manifest and terminates.

`hatt hash` supports multithreading, i.e. `-t 4`/`--threads=8`. This should probably be 1-2 if
you're using rotational storage (i.e. a hard drive), since reading more files at once will just
cause your drive to spend more time seeking. On the other hand, if you're on a disk array or an
SSD, consider setting it to the number of CPUs in your machine for maximum performance.

Interoperability
----------------

`hatt` doesn't exist in a vacuum; there's lots of other tools that do similar things. `hatt`
manifests aren't really useful outside of `hatt`, but `hatt` can talk to other formats.

Check `hatt help export` or `hatt formats` for more.
