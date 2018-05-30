# barchar

barchar is a simple visualization command for result of `uniq -q`.

## Usage

```shell
$ for i in {1..1000}; do echo label_$(((RANDOM%15+1))); done | sort | uniq -c | sort -nr | barchar -bar-width 10 -label-width 10
label_15   ########## 84
label_10   ######### 82
label_14   ######### 77
label_11   ######## 70
label_3    ######## 68
label_2    ######## 68
label_4    ####### 67
label_12   ####### 65
label_6    ####### 64
label_5    ####### 64
label_1    ####### 64
label_7    ####### 61
label_13   ###### 57
label_9    ###### 56
label_8    ###### 53
```

# License

MIT