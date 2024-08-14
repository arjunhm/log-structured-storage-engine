## log structured storage engine


### Key-Value Pair

- Key
- Value
- Size (key size + value size)
- Deleted (bool)

### MemTable

in-memory append-only log.  
flushes to disk once capacity exceeds.  

- Size (current size)
- Entries (KeyValue array)

**operations**
- get (reads KV)
- put (adds KV)
- del (adds KV but sets deleted as true)
- isFull (checks if memtable is exceeding capacity)
- flush (writes contents to SSTable)
- clear (clears memtable)

### SSTable
- Size
- 
