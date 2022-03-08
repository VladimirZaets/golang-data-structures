# Repository contains the list of data structures implemented by Golang

## Data structure complexity

### Static and Dynamic array

||Static Array| Dynamic Array|
| ----------- | ----------- | ----------- |
| Access | O(1) | O(1) |
| Search | O(n) | O(n) |
| Insertion | N/A | O(n) |
| Appending | N/A | O(1) |
| Deletion | N/A | O(n) |

----

### Single and Doubly Linked List

||Singly Linked List| Doubly Linked List|
| ----------- | ----------- | ----------- |
| Search | O(n) | O(n) |
| Insert at head | O(1) | O(1) |
| Insert at tail | O(1) | O(1) |
| Remove at head | O(1) | O(1) |
| Remove at tail | O(n) | O(1) |
| Remove in middle | O(n) | O(n) |

----

### Stack

|||
| ----------- | ----------- |
| Pushing | O(1) |
| Popping | O(1) |
| Peeking | O(1) |
| Searching | O(n) |
| Size | O(1) |

----

### Queue

|||
| ----------- | ----------- |
| Enqueue | O(1) |
| Dequeue | O(1) |
| Peeking | O(1) |
| Contains | O(n) |
| Removal | O(n) |
| Is Empty | O(1) |

----

### Priority Queue with binary heap

|||
| ----------- | ----------- |
| Binary Heap Construction  | O(n) |
| Polling | O(log(n)) |
| Peeking | O(1) |
| Adding | O(log(n)) |
| Naive Removing | O(n) |
| Naive Contains | O(n) |
| Removing with help of hash table | O(log(n)) |
| Containts check with help of hash table | O(1) |

----

### Union Find
|||
| ----------- | ----------- |
| Construction  | O(n) |
| Union | O(1)+ |
| Find | O(1)+ |
| Get Component size | O(1)+ |
| Check if connected | O(1)+ |
| Count | O(1) |
