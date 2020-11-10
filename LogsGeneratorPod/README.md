## What is this?
```
kubectl run   log-generator --generator=run-pod/v1 \
 --image alpine --env "X=1000" --env "Y=1000" --env "REPEATS=10" --env "SLEEP=0.1s" \
 -- sh -c 'for x in $(seq $X); do for y in $(seq $Y); do printf "%05d-%05d " $x $y; printf "%0.sabracadabra" $(seq $REPEATS); echo; sleep $SLEEP; done; done'
```

- Two nested loops controlled via X and Y
- Repeats the word *abracadabra* REPEATS number of times in the same line. This is the length of the line
- Sleeps between each line SLEEP seconds
