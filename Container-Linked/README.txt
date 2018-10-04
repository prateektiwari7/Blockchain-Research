PushBack:
This adds an element to the end of the list. It is like the append() method on a slice.

Slice
PushFront:
Inserts an element at the front of the list. This shows the advantage of using a container list—inserting at the front is fast.

Front, Next:
These funcs are used to iterate through the list. Iteration in a container list will be slower than in a slice or array.



Alternate to linked list 

Prateek-Tiwari:containerLinkedlist yudiz$ go run containerlist.go
Raju &{0xc42007a210 0xc42007a270 0xc42007a180 Rajasthan}
Rajasthan &{0xc42007a1b0 0xc42007a240 0xc42007a180 Raj}
Raj &{0xc42007a1e0 0xc42007a210 0xc42007a180 Prateek}
Prateek &{0xc42007a180 0xc42007a1b0 0xc42007a180 1}
1 <niladdress>