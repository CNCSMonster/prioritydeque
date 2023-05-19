# A double edge priorityQueue

it is developed based on "container/heap",which is a standard library

with this data structure,you can push data in it with 
O(n) time-complexity,and get the min or max value of all the values in it with O(logn) time-complexity.
you can also use PopMax or PopMin to remove the  max or min value in it.
for now,it has not been able to remove any value from it yet.

### some possible usage scenarios

1. some algorithms problem which need frequently find max or min value
2. to support a slide window which needs such characteristic ,for example ,it could be used for Linear scan register allocation.