# unpckr
unpckr is a cli tool that collects all files from sources into a single directory according to the specified rules


## Usage
unpckr has built-in usage:
```
usage: unpckr-cli [-h|--help] -s|--source "<value>" [-s|--source "<value>" ...]
                  [-d|--dest "<value>"] [-c|--conflict-rename (simpleRandom)]
                  [-r|--rename-all (hash|none)] [-p|--pattern "<value>"]
                  [-l|--log] [--log-level (INFO|WARN|none)]
                  [--remove-duplicates] [-z|--unzip]

                  Provides some unpacking functions

Arguments:

  -h  --help               Print help information
  -s  --source             Source folder, at least one
  -d  --dest               Destination folder. Default: temp/
  -c  --conflict-rename    Conflict rename strategy :[simpleRandom - add to
                           filename 10 random characters]. Default:
                           simpleRandom
  -r  --rename-all         Rename all files through selected strategy :[hash -
                           file get his hash string as name]. Default: none
  -p  --pattern            Rename all files trough specified pattern. Default:
                           none
  -l  --log                Show logging stream. Default: false
      --log-level          Logging level. Default: none
      --remove-duplicates  Remove duplicates basing on file`s hash. Default:
                           false
  -z  --unzip              Unzip all archives from the sources. Default: false
```


## Simple example
In this example, we are asking unpckr to collect files from two sources into the `D:\MyPhotos` folder. Rename all files using the -r key. Using `--remove-duplicates` to remove all duplicates (by hash). Also unzip all archives that will be inside the sources 
```
unpckr -s "C:\Users\Folder with Whitespace\photos" -s D:\files\photos\holidays -d D:\MyPhotos -r hash --remove-duplicates --unzip
```


## Order of operations/filters
***Will be created soon***


## Help
You can always use `unpckr -h` or `unpckr --help` to get actual list of keys and settings. Also, you can ask me everything trough [Issues](https://github.com/nekovalue/unpckr/issues)


## TODO
1. ~~Unit tests for *Argument parser*~~
3. ~~Copier closure handler~~
4. ~~Conflict rename function~~
5. ~~Rename all function~~
6. ~~Unzip function~~
7. ~~Duplicate remove function~~
8. ~~Integrate GitHub Actions~~
9. Advanced patterns support
10. Logging
11. Write operations order in README
12. Release 1.0
13. Comment all functions
14. Fatal log for copier functions
15. Move some packages to pkg folder
16. Unit tests for all filters
    